package container

import (
	"context"
	"database/sql"
	"frogsmash/internal/config"
	"frogsmash/internal/infrastructure/email"
	"frogsmash/internal/infrastructure/storage"
)

type InfraServices struct {
	DB            *sql.DB
	UploadService *storage.UploadService
	EmailService  *email.EmailService
}

func NewInfraServices(cfg *config.Config, ctx context.Context) (*InfraServices, error) {
	db, err := sql.Open("postgres", cfg.DatabaseConfig.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

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

	return &InfraServices{
		UploadService: uploadService,
		EmailService:  emailService,
		DB:            db,
	}, nil
}
