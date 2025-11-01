package services

import (
	"fmt"
	"frogsmash/internal/app/models"
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

type ItemService struct {
	Repo ItemsRepo
}

func NewItemService(repo ItemsRepo) *ItemService {
	return &ItemService{Repo: repo}
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
