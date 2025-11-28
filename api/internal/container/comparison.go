package container

import (
	"frogsmash/internal/app/comparison/repos"
	"frogsmash/internal/app/comparison/services"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/config"
	"time"
)

type Comparison struct {
	ComparisonService services.ComparisonService
	SubmissionService services.SubmissionService
	ScoreUpdater      services.ScoreUpdater
}

func NewComparison(cfg *config.Config, db shared.DBWithTxStarter, uploadService services.UploadService, verificationService services.VerificationService) *Comparison {
	eventsRepo := repos.NewEventsRepo()
	eventsService := services.NewEventsService(eventsRepo)

	itemsRepo := repos.NewItemsRepo()
	comparisonService := services.NewComparisonService(itemsRepo, eventsService)
	updateInterval := time.Duration(cfg.AppConfig.ScoreUpdateInterval) * time.Second

	submissionRepo := repos.NewSubmissionRepo()
	submissionService := services.NewSubmissionService(uploadService, submissionRepo, verificationService)

	scoreUpdater := services.NewScoreUpdater(db, eventsRepo, itemsRepo, cfg.AppConfig.KFactor, updateInterval)
	return &Comparison{
		ComparisonService: comparisonService,
		SubmissionService: submissionService,
		ScoreUpdater:      scoreUpdater,
	}
}
