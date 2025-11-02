package container

import (
	"database/sql"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/app/services"
	"frogsmash/internal/config"
)

type Container struct {
	DB           *sql.DB
	ItemsService *services.ItemService
	ScoreUpdater *services.ScoreUpdater
}

func NewContainer(cfg *config.Config) (*Container, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	eventsRepo := repos.NewEventsRepo()
	eventsService := services.NewEventsService(eventsRepo)

	itemsRepo := repos.NewItemsRepo()
	itemsService := services.NewItemService(itemsRepo, eventsService)

	scoreUpdater := services.NewScoreUpdater(db, eventsRepo, itemsRepo)

	return &Container{
		DB:           db,
		ItemsService: itemsService,
		ScoreUpdater: scoreUpdater,
	}, nil
}
