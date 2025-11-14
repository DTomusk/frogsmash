package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
)

type ItemsRepo interface {
	GetRandomItems(numberOfItems int, ctx context.Context, db repos.DBTX) ([]models.Item, error)
	GetItemsByIds(ids []string, ctx context.Context, db repos.DBTX) ([]*models.Item, error)
	GetItemById(id string, ctx context.Context, db repos.DBTX) (*models.Item, error)
	UpdateItemScore(itemID string, newScore float64, ctx context.Context, db repos.DBTX) error
	GetLeaderboardItems(limit int, offset int, ctx context.Context, db repos.DBTX) ([]*models.LeaderboardItem, error)
	GetTotalItemCount(ctx context.Context, db repos.DBTX) (int, error)
}

type ItemService struct {
	Repo          ItemsRepo
	EventsService *EventsService
}

func NewItemService(repo ItemsRepo, eventsService *EventsService) *ItemService {
	return &ItemService{Repo: repo, EventsService: eventsService}
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

func (s *ItemService) CompareItems(winnerId, loserId, userId string, ctx context.Context, db repos.DBTX) error {
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
	return s.EventsService.LogEvent(winnerId, loserId, userId, ctx, db)
}

func (s *ItemService) GetLeaderboardPage(limit int, offset int, ctx context.Context, db repos.DBTX) ([]*models.LeaderboardItem, int, error) {
	// Placeholder implementation, replace with repo call
	var items []*models.LeaderboardItem
	items, err := s.Repo.GetLeaderboardItems(limit, offset, ctx, db)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.Repo.GetTotalItemCount(ctx, db)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
