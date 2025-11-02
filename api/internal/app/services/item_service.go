package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"math"
)

type ItemsRepo interface {
	GetRandomItems(numberOfItems int, ctx context.Context, db repos.DBTX) ([]models.Item, error)
	GetItemsByIds(ids []string, ctx context.Context, db repos.DBTX) ([]*models.Item, error)
}

// TODO: read kfactor from config and only expose getter
type ItemService struct {
	Repo          ItemsRepo
	EventsService *EventsService
	kFactor       float64
}

func NewItemService(repo ItemsRepo, eventsService *EventsService) *ItemService {
	return &ItemService{Repo: repo, EventsService: eventsService, kFactor: 32.0}
}

func (s *ItemService) GetComparisonItems(ctx context.Context, db repos.DBTX) (*models.Item, *models.Item, error) {
	items, err := s.Repo.GetRandomItems(2, ctx, db)
	if err != nil {
		return nil, nil, err
	}
	if len(items) < 2 {
		return nil, nil, fmt.Errorf("not enough items available for comparison")
	}
	return &items[0], &items[1], nil
}

func (s *ItemService) CompareItems(winnerId, loserId string, ctx context.Context, db repos.DBTX) error {
	if winnerId == loserId {
		return fmt.Errorf("winner and loser cannot be the same")
	}
	// Validate items exist
	items, err := s.Repo.GetItemsByIds([]string{winnerId, loserId}, ctx, db)
	if err != nil {
		return err
	}
	if len(items) != 2 {
		return fmt.Errorf("one or both items not found")
	}
	// Log event to be picked up by worker later
	return s.EventsService.LogEvent(winnerId, loserId, ctx, db)
}

func (s *ItemService) UpdateEloScores(winner, loser *models.Item) {
	expectedWinner := 1 / (1 + math.Pow(10, (loser.Score-winner.Score)/400))
	expectedLoser := 1 / (1 + math.Pow(10, (winner.Score-loser.Score)/400))
	winner.Score += s.kFactor * (1 - expectedWinner)
	loser.Score += s.kFactor * (0 - expectedLoser)
}
