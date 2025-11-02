package services

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"log"
	"time"
)

type EventRepo interface {
	GetNextUnprocessedEvent(ctx context.Context, db repos.DBTX) (*models.Event, error)
	SetEventProcessed(eventID string, ctx context.Context, db repos.DBTX) error
}

type ScoreUpdater struct {
	db        *sql.DB
	EventRepo EventRepo
}

func NewScoreUpdater(db *sql.DB, er EventRepo) *ScoreUpdater {
	return &ScoreUpdater{db: db, EventRepo: er}
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
			su.handleEvent(ctx)
			time.Sleep(10 * time.Second)
		}
	}
}

func (su *ScoreUpdater) handleEvent(ctx context.Context) {
	event, err := su.EventRepo.GetNextUnprocessedEvent(ctx, su.db)
	if err != nil {
		log.Printf("Error fetching event: %v", err)
		return
	}
	if event == nil {
		log.Println("No unprocessed events found.")
		return
	}
	log.Printf("Processing event ID: %s", event.ID)
	// Do a transaction to update scores and mark event as processed
	tx, err := su.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)
		return
	}
	defer tx.Rollback()

	// Update scores and mark event as processed
	if err := su.EventRepo.SetEventProcessed(event.ID, ctx, tx); err != nil {
		log.Printf("Error processing event: %v", err)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return
	}
}
