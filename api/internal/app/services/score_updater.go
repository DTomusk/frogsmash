package services

import (
	"context"
	"database/sql"
	"fmt"
	"frogsmash/internal/app/models"
	"frogsmash/internal/app/repos"
	"log"
	"math"
	"time"
)

type EventRepo interface {
	GetNextUnprocessedEvent(ctx context.Context, db repos.DBTX) (*models.Event, error)
	SetEventProcessed(eventID string, ctx context.Context, db repos.DBTX) error
}

// TODO: read kfactor from config and only expose getter
type ScoreUpdater struct {
	db        *sql.DB
	EventRepo EventRepo
	ItemsRepo ItemsRepo
	kFactor   float64
}

func NewScoreUpdater(db *sql.DB, er EventRepo, ir ItemsRepo) *ScoreUpdater {
	return &ScoreUpdater{db: db, EventRepo: er, ItemsRepo: ir, kFactor: 32.0}
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

	winner, loser, err := su.GetWinnerAndLoser(event.WinnerID, event.LoserID, ctx)
	if err != nil {
		log.Printf("Error getting winner and loser: %v", err)
		return
	}

	// Do a transaction to update scores and mark event as processed
	tx, err := su.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("Error beginning transaction: %v", err)
		return
	}
	defer tx.Rollback()

	su.updateEloScores(winner, loser)

	log.Printf("Updated scores - Winner: %s (%.2f), Loser: %s (%.2f)", winner.ID, winner.Score, loser.ID, loser.Score)
	if err := su.ItemsRepo.UpdateItemScore(winner.ID, winner.Score, ctx, tx); err != nil {
		log.Printf("Error updating winner score: %v", err)
		return
	}
	if err := su.ItemsRepo.UpdateItemScore(loser.ID, loser.Score, ctx, tx); err != nil {
		log.Printf("Error updating loser score: %v", err)
		return
	}

	if err := su.EventRepo.SetEventProcessed(event.ID, ctx, tx); err != nil {
		log.Printf("Error processing event: %v", err)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error committing transaction: %v", err)
		return
	}

	log.Printf("Event ID %s processed successfully.", event.ID)
}

func (su *ScoreUpdater) GetWinnerAndLoser(winnerId, loserId string, ctx context.Context) (*models.Item, *models.Item, error) {
	winner, err := su.ItemsRepo.GetItemById(winnerId, ctx, su.db)
	if err != nil {
		log.Printf("Error fetching winner item: %v", err)
		return nil, nil, err
	}
	if winner == nil {
		log.Println("Winner item not found.")
		return nil, nil, fmt.Errorf("winner item not found")
	}

	loser, err := su.ItemsRepo.GetItemById(loserId, ctx, su.db)
	if err != nil {
		log.Printf("Error fetching loser item: %v", err)
		return nil, nil, err
	}
	if loser == nil {
		log.Println("Loser item not found.")
		return nil, nil, fmt.Errorf("loser item not found")
	}
	return winner, loser, nil
}

func (su *ScoreUpdater) updateEloScores(winner, loser *models.Item) {
	expectedWinner := 1 / (1 + math.Pow(10, (loser.Score-winner.Score)/400))
	expectedLoser := 1 / (1 + math.Pow(10, (winner.Score-loser.Score)/400))
	winner.Score += su.kFactor * (1 - expectedWinner)
	loser.Score += su.kFactor * (0 - expectedLoser)
}
