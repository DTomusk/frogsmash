package services

import (
	"context"
	"fmt"
	"frogsmash/internal/app/comparison/models"
	"frogsmash/internal/app/shared"
)

type ItemsRepo interface {
	GetRandomItems(numberOfItems int, tenantID string, ctx context.Context, db shared.DBTX) ([]models.Item, error)
	GetItemsByIds(ids []string, ctx context.Context, db shared.DBTX) ([]*models.Item, error)
	GetItemById(id string, ctx context.Context, db shared.DBTX) (*models.Item, error)
	UpdateItemScore(itemID string, newScore float64, ctx context.Context, db shared.DBTX) error
	GetLeaderboardItems(limit int, offset int, tenantID string, ctx context.Context, db shared.DBTX) ([]*models.LeaderboardItem, error)
	GetTotalItemCount(tenantID string, ctx context.Context, db shared.DBTX) (int, error)
}

type ComparisonService interface {
	GetComparisonItems(ctx context.Context, db shared.DBTX, tenantID string) (*models.Item, *models.Item, error)
	CompareItems(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error
	GetLeaderboardPage(limit int, offset int, tenantID string, ctx context.Context, db shared.DBTX) ([]*models.LeaderboardItem, int, error)
}

type comparisonService struct {
	Repo          ItemsRepo
	EventsService EventsService
}

func NewComparisonService(repo ItemsRepo, eventsService EventsService) ComparisonService {
	return &comparisonService{Repo: repo, EventsService: eventsService}
}

func (s *comparisonService) GetComparisonItems(ctx context.Context, db shared.DBTX, tenantID string) (*models.Item, *models.Item, error) {
	items, err := s.Repo.GetRandomItems(2, tenantID, ctx, db)
	if err != nil {
		return nil, nil, err
	}
	if len(items) < 2 {
		return nil, nil, fmt.Errorf("not enough items available for comparison")
	}
	return &items[0], &items[1], nil
}

func (s *comparisonService) CompareItems(winnerId, loserId, userId string, ctx context.Context, db shared.DBTX) error {
	if winnerId == loserId {
		return fmt.Errorf("winner and loser cannot be the same")
	}
	// Validate items exist
	items, err := s.Repo.GetItemsByIds([]string{winnerId, loserId}, ctx, db)
	if err != nil {
		return err
	}
	if len(items) != 2 {
		return fmt.Errorf("one or both items not found")
	}
	// Log event to be picked up by worker later
	return s.EventsService.LogEvent(winnerId, loserId, userId, ctx, db)
}

func (s *comparisonService) GetLeaderboardPage(limit int, offset int, tenantID string, ctx context.Context, db shared.DBTX) ([]*models.LeaderboardItem, int, error) {
	// Placeholder implementation, replace with repo call
	var items []*models.LeaderboardItem
	items, err := s.Repo.GetLeaderboardItems(limit, offset, tenantID, ctx, db)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.Repo.GetTotalItemCount(tenantID, ctx, db)
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}
