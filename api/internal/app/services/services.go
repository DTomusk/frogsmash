package services

import "frogsmash/internal/app/models"

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

type ItemService struct {
	// Add dependencies if needed
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (s *ItemService) GetComparisonItems() (*models.Item, *models.Item, error) {
	// Placeholder implementation
	return &models.Item{ID: "1"}, &models.Item{ID: "2"}, nil
}
