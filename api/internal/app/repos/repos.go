package repos

import (
	"database/sql"
	"frogsmash/internal/app/models"

	"github.com/lib/pq"
)

type EventsRepo struct {
	DB *sql.DB
}

type ItemsRepo struct {
	DB *sql.DB
}

func NewEventsRepo(db *sql.DB) *EventsRepo {
	return &EventsRepo{DB: db}
}

func (r *EventsRepo) LogEvent(winnerId, loserId string) error {
	_, err := r.DB.Exec(
		"INSERT INTO events (winner_id, loser_id) VALUES ($1, $2)",
		winnerId, loserId,
	)
	return err
}

func NewItemsRepo(db *sql.DB) *ItemsRepo {
	return &ItemsRepo{DB: db}
}

func (r *ItemsRepo) GetRandomItems(numberOfItems int) ([]models.Item, error) {
	var items []models.Item
	query := "SELECT id, name, image_url, score FROM items ORDER BY RANDOM() LIMIT $1"
	rows, err := r.DB.Query(query, numberOfItems)
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

func (r *ItemsRepo) GetItemsByIds(ids []string) ([]*models.Item, error) {
	var items []*models.Item
	query := "SELECT id, name, image_url, score FROM items WHERE id = ANY($1)"
	rows, err := r.DB.Query(query, pq.Array(ids))
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
