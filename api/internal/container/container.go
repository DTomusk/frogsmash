package container

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/app/services"
	"frogsmash/internal/app/utils"
	"frogsmash/internal/config"
	"time"
)

// TODO: do we want the allowed origin here?
// TODO: consider injecting max file size into upload service instead
type Container struct {
	DB             *sql.DB
	ItemsService   *services.ItemService
	ScoreUpdater   *services.ScoreUpdater
	UploadService  *services.UploadService
	AuthService    *services.AuthService
	AllowedOrigin  string
	MaxRequestSize int64
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

	err = storageClient.Ping(ctx)
	if err != nil {
		return nil, err
	}

	uploadService := services.NewUploadService(storageClient, cfg.MaxFileSize)

	userRepo := repos.NewUserRepo()
	hasher := utils.NewHasher()
	authService := services.NewAuthService(userRepo, hasher)

	return &Container{
		DB:             db,
		ItemsService:   itemsService,
		ScoreUpdater:   scoreUpdater,
		UploadService:  uploadService,
		AuthService:    authService,
		AllowedOrigin:  cfg.AllowedOrigin,
		MaxRequestSize: cfg.MaxFileSize,
	}, nil
}
