package repos

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/models"

	"github.com/lib/pq"
)

type EventsRepo struct{}

type ItemsRepo struct{}

func NewEventsRepo() *EventsRepo {
	return &EventsRepo{}
}

func (r *EventsRepo) LogEvent(winnerId, loserId string, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx,
		"INSERT INTO events (winner_id, loser_id) VALUES ($1, $2)",
		winnerId, loserId,
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

func NewItemsRepo() *ItemsRepo {
	return &ItemsRepo{}
}

func (r *ItemsRepo) GetRandomItems(numberOfItems int, ctx context.Context, db DBTX) ([]models.Item, error) {
	var items []models.Item
	query := "SELECT id, name, image_url, score FROM items ORDER BY RANDOM() LIMIT $1"
	rows, err := db.QueryContext(ctx, query, numberOfItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.ImageURL, &item.Score); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (r *ItemsRepo) GetItemById(id string, ctx context.Context, db DBTX) (*models.Item, error) {
	query := "SELECT id, name, image_url, score FROM items WHERE id = $1"
	row := db.QueryRowContext(ctx, query, id)
	var item models.Item
	if err := row.Scan(&item.ID, &item.Name, &item.ImageURL, &item.Score); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &item, nil
}

func (r *ItemsRepo) GetItemsByIds(ids []string, ctx context.Context, db DBTX) ([]*models.Item, error) {
	var items []*models.Item
	query := "SELECT id, name, image_url, score FROM items WHERE id = ANY($1)"
	rows, err := db.QueryContext(ctx, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.ImageURL, &item.Score); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *ItemsRepo) UpdateItemScore(itemID string, newScore float64, ctx context.Context, db DBTX) error {
	_, err := db.ExecContext(ctx,
		"UPDATE items SET score = $1 WHERE id = $2", newScore, itemID,
	)
	return err
}
