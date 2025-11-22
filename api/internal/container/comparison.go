package container

import (
	"frogsmash/internal/app/comparison/repos"
	"frogsmash/internal/app/comparison/services"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/config"
	"time"
)

type Comparison struct {
	ItemsService services.ItemService
	ScoreUpdater services.ScoreUpdater
}

func NewComparison(cfg *config.Config, db shared.DBWithTxStarter) *Comparison {
	eventsRepo := repos.NewEventsRepo()
	eventsService := services.NewEventsService(eventsRepo)

	itemsRepo := repos.NewItemsRepo()
	itemsService := services.NewItemService(itemsRepo, eventsService)
	updateInterval := time.Duration(cfg.AppConfig.ScoreUpdateInterval) * time.Second

	scoreUpdater := services.NewScoreUpdater(db, eventsRepo, itemsRepo, cfg.AppConfig.KFactor, updateInterval)
	return &Comparison{
		ItemsService: itemsService,
		ScoreUpdater: scoreUpdater,
	}
}
