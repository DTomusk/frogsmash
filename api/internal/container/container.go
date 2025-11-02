package container

import (
	"database/sql"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/app/services"
	"frogsmash/internal/config"
)

type Container struct {
	DB            *sql.DB
	EventsService *services.EventsService
	ItemsService  *services.ItemService
}

func NewContainer(cfg *config.Config) (*Container, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	eventsRepo := repos.NewEventsRepo(db)
	eventsService := services.NewEventsService(eventsRepo)

	itemsRepo := repos.NewItemsRepo(db)
	itemsService := services.NewItemService(itemsRepo, eventsService)

	return &Container{
		DB:           db,
		ItemsService: itemsService,
	}, nil
}
