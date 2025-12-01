package container

import (
	"context"
	"database/sql"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/config"
	"frogsmash/internal/infrastructure/email"
	"frogsmash/internal/infrastructure/messages"
	"frogsmash/internal/infrastructure/redis"
	"frogsmash/internal/infrastructure/storage"
)

type InfraServices struct {
	DB             shared.DBWithTxStarter
	UploadService  storage.UploadService
	EmailService   email.EmailService
	Dispatcher     messages.Dispatcher
	MessageService messages.MessageService
	RedisClient    redis.RedisClient
}

func NewInfraServices(cfg *config.Config, ctx context.Context) (*InfraServices, error) {
	db, err := sql.Open("postgres", cfg.DatabaseConfig.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	dbWithTxStarter := shared.NewPostgresDB(db)

	storageClient, err := storage.NewStorageClient(
		ctx,
		cfg.StorageConfig.StorageAccountID,
		cfg.StorageConfig.StorageAccessKey,
		cfg.StorageConfig.StorageSecretKey,
		cfg.StorageConfig.StorageBucket,
	)

	if err != nil {
		return nil, err
	}

	err = storageClient.Ping(ctx)
	if err != nil {
		return nil, err
	}

	uploadService := storage.NewUploadService(storageClient, cfg.AppConfig.MaxFileSize)

	emailClient := email.NewMailjetClient(cfg.MailConfig.MailjetAPIKey, cfg.MailConfig.MailjetSecretKey, cfg.MailConfig.SenderEmail)
	templateRenderer, err := email.NewTemplateRenderer(cfg.MailConfig.TemplateGlobPattern)
	if err != nil {
		return nil, err
	}

	emailService := email.NewEmailService(emailClient, templateRenderer, cfg.AppConfig.AppURL)

	dispatcher := messages.NewDispatcher()

	redisClient := redis.NewRedisClient(cfg.MessageConfig.RedisAddress, "mystream", "mygroup", "consumer1")

	messageService, err := messages.NewMessageService(ctx, redisClient, dispatcher, dbWithTxStarter)
	if err != nil {
		return nil, err
	}

	return &InfraServices{
		UploadService:  uploadService,
		EmailService:   emailService,
		DB:             dbWithTxStarter,
		Dispatcher:     dispatcher,
		MessageService: messageService,
		RedisClient:    redisClient,
	}, nil
}
