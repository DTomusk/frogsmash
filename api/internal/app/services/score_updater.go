package services

import (
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

func (su *ScoreUpdater) Run() {
	// Implementation for processing events and updating scores would go here
	log.Println("ScoreUpdater is running...")
	for {
		log.Println("Checking for unprocessed events...")
		time.Sleep(10 * time.Second)
	}
}
