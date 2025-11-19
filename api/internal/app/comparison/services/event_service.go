package services

import (
	"context"
	"frogsmash/internal/app/comparison/models"
	"frogsmash/internal/app/shared"
)

type EventsRepo interface {
	LogEvent(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error
	SetEventProcessed(eventID string, ctx context.Context, db shared.DBTX) error
	SetEventFailed(eventID string, ctx context.Context, db shared.DBTX) error
	GetNextUnprocessedEvent(ctx context.Context, db shared.DBTX) (*models.Event, error)
}

type EventsService struct {
	Repo EventsRepo
}

func NewEventsService(repo EventsRepo) *EventsService {
	return &EventsService{Repo: repo}
}

func (s *EventsService) LogEvent(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error {
	return s.Repo.LogEvent(winnerId, loserId, userId, ctx, db)
}

func (s *EventsService) SetEventProcessed(eventID string, ctx context.Context, db shared.DBTX) error {
	return s.Repo.SetEventProcessed(eventID, ctx, db)
}
