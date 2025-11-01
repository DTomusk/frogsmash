package services

import (
	"fmt"
	"frogsmash/internal/app/models"
	"math"
)

type EventsRepo interface {
	LogEvent(winnerId, loserId string) error
}

type EventsService struct {
	Repo EventsRepo
}

func NewEventsService(repo EventsRepo) *EventsService {
	return &EventsService{Repo: repo}
}

func (s *EventsService) LogEvent(winnerId, loserId string) error {
	return s.Repo.LogEvent(winnerId, loserId)
}

type ItemsRepo interface {
	GetRandomItems(numberOfItems int) ([]models.Item, error)
}

// TODO: read kfactor from config and only expose getter
type ItemService struct {
	Repo    ItemsRepo
	kFactor float64
}

func NewItemService(repo ItemsRepo) *ItemService {
	return &ItemService{Repo: repo, kFactor: 32.0}
}

func (s *ItemService) GetComparisonItems() (*models.Item, *models.Item, error) {
	items, err := s.Repo.GetRandomItems(2)
	if err != nil {
		return nil, nil, err
	}
	if len(items) < 2 {
		return nil, nil, fmt.Errorf("not enough items available for comparison")
	}
	return &items[0], &items[1], nil
}

func (s *ItemService) UpdateEloScores(winner, loser *models.Item) {
	expectedWinner := 1 / (1 + math.Pow(10, (loser.Score-winner.Score)/400))
	expectedLoser := 1 / (1 + math.Pow(10, (winner.Score-loser.Score)/400))
	winner.Score += s.kFactor * (1 - expectedWinner)
	loser.Score += s.kFactor * (0 - expectedLoser)
}
