package services

import (
	"context"
	"frogsmash/internal/app/models"
	"log"
	"time"
)

type EventReader interface {
	GetNextUnprocessedEvent() (*models.Event, error)
}

type ScoreUpdater struct {
	EventReader EventReader
}

func NewScoreUpdater(r EventReader) *ScoreUpdater {
	return &ScoreUpdater{EventReader: r}
}

func (su *ScoreUpdater) Run(ctx context.Context) {
	// Implementation for processing events and updating scores would go here
	log.Println("ScoreUpdater is running...")
	for {
		select {
		case <-ctx.Done():
			log.Println("ScoreUpdater is stopping...")
			return
		default:
			log.Println("Checking for unprocessed events...")
			time.Sleep(10 * time.Second)
		}
	}
}
