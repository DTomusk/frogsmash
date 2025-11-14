package services

import (
	"context"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
)

type EventsRepo interface {
	LogEvent(winnerId, loserId, userId string, ctx context.Context, db repos.DBTX) error
	SetEventProcessed(eventID string, ctx context.Context, db repos.DBTX) error
	SetEventFailed(eventID string, ctx context.Context, db repos.DBTX) error
	GetNextUnprocessedEvent(ctx context.Context, db repos.DBTX) (*models.Event, error)
}

type EventsService struct {
	Repo EventsRepo
}

func NewEventsService(repo EventsRepo) *EventsService {
	return &EventsService{Repo: repo}
}

func (s *EventsService) LogEvent(winnerId, loserId, userId string, ctx context.Context, db repos.DBTX) error {
	return s.Repo.LogEvent(winnerId, loserId, userId, ctx, db)
}

func (s *EventsService) SetEventProcessed(eventID string, ctx context.Context, db repos.DBTX) error {
	return s.Repo.SetEventProcessed(eventID, ctx, db)
}
