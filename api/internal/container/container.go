package container

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/clients"
	"frogsmash/internal/app/factories"
	"frogsmash/internal/app/repos"
	"frogsmash/internal/app/services"
	"frogsmash/internal/config"
	"frogsmash/internal/email"
	"time"
)

// TODO: do we want the allowed origin here?
// TODO: consider injecting max file size into upload service instead
type Container struct {
	DB                  *sql.DB
	ItemsService        *services.ItemService
	ScoreUpdater        *services.ScoreUpdater
	UploadService       *services.UploadService
	AuthService         *services.AuthService
	JwtService          *services.JwtService
	AllowedOrigin       string
	MaxRequestSize      int64
	EmailService        *email.EmailService
	UserService         *services.UserService
	VerificationService *services.VerificationService
}

func NewContainer(cfg *config.Config) (*Container, error) {
	ctx := context.Background()
	db, err := sql.Open("postgres", cfg.DatabaseConfig.DatabaseURL)
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

	updateInterval := time.Duration(cfg.AppConfig.ScoreUpdateInterval) * time.Second

	scoreUpdater := services.NewScoreUpdater(db, eventsRepo, itemsRepo, cfg.AppConfig.KFactor, updateInterval)

	storageClient, err := clients.NewStorageClient(ctx, cfg.StorageConfig.StorageAccountID, cfg.StorageConfig.StorageAccessKey, cfg.StorageConfig.StorageSecretKey, cfg.StorageConfig.StorageBucket)
	if err != nil {
		return nil, err
	}

	err = storageClient.Ping(ctx)
	if err != nil {
		return nil, err
	}

	uploadService := services.NewUploadService(storageClient, cfg.AppConfig.MaxFileSize)

	emailClient := email.NewMailjetClient(cfg.MailConfig.MailjetAPIKey, cfg.MailConfig.MailjetSecretKey, cfg.MailConfig.SenderEmail)
	templateRenderer, err := email.NewTemplateRenderer(cfg.MailConfig.TemplateGlobPattern)
	if err != nil {
		return nil, err
	}
	emailService := email.NewEmailService(emailClient, templateRenderer, cfg.AppConfig.AppURL)

	userRepo := repos.NewUserRepo()
	refreshTokenRepo := repos.NewRefreshTokenRepo()
	verificationRepo := repos.NewVerificationRepo()
	hasher := services.NewBCryptHasher()
	tokenService := services.NewJwtService([]byte(cfg.TokenConfig.JWTSecret), cfg.TokenConfig.TokenLifetimeMinutes)
	userFactory := factories.NewUserFactory(hasher)
	verificationService := services.NewVerificationService(
		userRepo,
		verificationRepo,
		emailService,
		cfg.TokenConfig.VerificationCodeLength,
		cfg.TokenConfig.VerificationCodeLifetimeMinutes,
	)
	userService := services.NewUserService(userFactory, userRepo, verificationService)
	authService := services.NewAuthService(
		userRepo,
		refreshTokenRepo,
		hasher,
		tokenService,
		cfg.TokenConfig.RefreshTokenLifetimeDays,
	)

	return &Container{
		DB:                  db,
		ItemsService:        itemsService,
		ScoreUpdater:        scoreUpdater,
		UploadService:       uploadService,
		AuthService:         authService,
		JwtService:          tokenService,
		AllowedOrigin:       cfg.AppConfig.AllowedOrigin,
		MaxRequestSize:      cfg.AppConfig.MaxFileSize,
		EmailService:        emailService,
		UserService:         userService,
		VerificationService: verificationService,
	}, nil
}
