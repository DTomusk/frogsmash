package container

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/clients"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/app/services"
	"frogsmash/internal/config"
	"frogsmash/internal/email"
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
	JwtService     *services.JwtService
	AllowedOrigin  string
	MaxRequestSize int64
	EmailService   *email.EmailService
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

	storageClient, err := clients.NewStorageClient(ctx, cfg.StorageAccountID, cfg.StorageAccessKey, cfg.StorageSecretKey, cfg.StorageBucket)
	if err != nil {
		return nil, err
	}

	err = storageClient.Ping(ctx)
	if err != nil {
		return nil, err
	}

	uploadService := services.NewUploadService(storageClient, cfg.MaxFileSize)

	userRepo := repos.NewUserRepo()
	refreshTokenRepo := repos.NewRefreshTokenRepo()
	verificationRepo := repos.NewVerificationRepo()
	hasher := services.NewBCryptHasher()
	tokenService := services.NewJwtService([]byte(cfg.JWTSecret), cfg.TokenLifetimeMinutes)
	authService := services.NewAuthService(userRepo, refreshTokenRepo, hasher, tokenService, verificationRepo, cfg.RefreshTokenLifetimeDays, cfg.VerificationCodeLength, cfg.VerificationCodeLifetimeMinutes)

	emailClient := email.NewMailjetClient(cfg.MailjetAPIKey, cfg.MailjetSecretKey, cfg.SenderEmail)
	templateRenderer, err := email.NewTemplateRenderer(cfg.TemplateGlobPattern)
	if err != nil {
		return nil, err
	}
	emailService := email.NewEmailService(emailClient, templateRenderer, cfg.AppURL)

	return &Container{
		DB:             db,
		ItemsService:   itemsService,
		ScoreUpdater:   scoreUpdater,
		UploadService:  uploadService,
		AuthService:    authService,
		JwtService:     tokenService,
		AllowedOrigin:  cfg.AllowedOrigin,
		MaxRequestSize: cfg.MaxFileSize,
		EmailService:   emailService,
	}, nil
}
