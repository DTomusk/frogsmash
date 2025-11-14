package repos

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/models"
)

type EventsRepo struct{}

func NewEventsRepo() *EventsRepo {
	return &EventsRepo{}
}

func (r *EventsRepo) LogEvent(winnerId, loserId, userId string, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO events (winner_id, loser_id, user_id) VALUES ($1, $2, $3)",
		winnerId, loserId, userId,
	)
	return err
}

func (r *EventsRepo) GetNextUnprocessedEvent(ctx context.Context, db DBTX) (*models.Event, error) {
	query := "SELECT id, winner_id, loser_id FROM events WHERE processed_at IS NULL AND failed_to_process = FALSE ORDER BY created_at ASC LIMIT 1"
	row := db.QueryRowContext(ctx, query)
	var event models.Event
	if err := row.Scan(&event.ID, &event.WinnerID, &event.LoserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &event, nil
}

func (r *EventsRepo) SetEventProcessed(eventID string, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx,
		"UPDATE events SET processed_at = NOW() WHERE id = $1", eventID,
	)
	return err
}

func (r *EventsRepo) SetEventFailed(eventID string, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx,
		"UPDATE events SET failed_to_process = TRUE WHERE id = $1", eventID,
	)
	return err
}
