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

type EventsService interface {
	LogEvent(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error
	SetEventProcessed(eventID string, ctx context.Context, db shared.DBTX) error
}

type eventsService struct {
	Repo EventsRepo
}

func NewEventsService(repo EventsRepo) EventsService {
	return &eventsService{Repo: repo}
}

func (s *eventsService) LogEvent(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error {
	return s.Repo.LogEvent(winnerId, loserId, userId, ctx, db)
}

func (s *eventsService) SetEventProcessed(eventID string, ctx context.Context, db shared.DBTX) error {
	return s.Repo.SetEventProcessed(eventID, ctx, db)
}
