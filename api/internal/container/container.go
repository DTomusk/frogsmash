package container

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/app/services"
	"frogsmash/internal/config"
	"time"
)

// TODO: do we want the allowed origin here?
type Container struct {
	DB            *sql.DB
	ItemsService  *services.ItemService
	ScoreUpdater  *services.ScoreUpdater
	UploadService *services.UploadService
	AllowedOrigin string
}

func NewContainer(cfg *config.Config) (*Container, error) {
	ctx := context.Background()
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

	updateInterval := time.Duration(cfg.ScoreUpdateInterval) * time.Second

	scoreUpdater := services.NewScoreUpdater(db, eventsRepo, itemsRepo, cfg.KFactor, updateInterval)

	storageClient, err := repos.NewStorageClient(ctx, cfg.StorageAccountID, cfg.StorageAccessKey, cfg.StorageSecretKey, cfg.StorageBucket)
	if err != nil {
		return nil, err
	}
	uploadService := services.NewUploadService(storageClient)

	return &Container{
		DB:            db,
		ItemsService:  itemsService,
		ScoreUpdater:  scoreUpdater,
		UploadService: uploadService,
		AllowedOrigin: cfg.AllowedOrigin,
	}, nil
}
