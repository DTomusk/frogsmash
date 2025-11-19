package repos

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/comparison/models"
	"frogsmash/internal/app/shared"

	"github.com/lib/pq"
)

type ItemsRepo struct{}

func NewItemsRepo() *ItemsRepo {
	return &ItemsRepo{}
}

func (r *ItemsRepo) GetRandomItems(numberOfItems int, ctx context.Context, db shared.DBTX) ([]models.Item, error) {
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

func (r *ItemsRepo) GetItemById(id string, ctx context.Context, db shared.DBTX) (*models.Item, error) {
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

func (r *ItemsRepo) GetItemsByIds(ids []string, ctx context.Context, db shared.DBTX) ([]*models.Item, error) {
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

func (r *ItemsRepo) UpdateItemScore(itemID string, newScore float64, ctx context.Context, db shared.DBTX) error {
	_, err := db.ExecContext(ctx,
		"UPDATE items SET score = $1 WHERE id = $2", newScore, itemID,
	)
	return err
}

func (r *ItemsRepo) GetLeaderboardItems(limit int, offset int, ctx context.Context, db shared.DBTX) ([]*models.LeaderboardItem, error) {
	var items []*models.LeaderboardItem
	query := "SELECT id, name, image_url, score, RANK() OVER (ORDER BY score DESC) as rank, created_at, license FROM items ORDER BY score DESC LIMIT $1 OFFSET $2"
	rows, err := db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.LeaderboardItem
		if err := rows.Scan(&item.ID, &item.Name, &item.ImageURL, &item.Score, &item.Rank, &item.CreatedAt, &item.License); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *ItemsRepo) GetTotalItemCount(ctx context.Context, db shared.DBTX) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM items"
	row := db.QueryRowContext(ctx, query)
	if err := row.Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}
