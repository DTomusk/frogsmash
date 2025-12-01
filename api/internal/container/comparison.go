package container

import (
	"frogsmash/internal/app/comparison/repos"
	"frogsmash/internal/app/comparison/services"
	"frogsmash/internal/app/shared"
	"frogsmash/internal/config"
)

type Comparison struct {
	// TODO: we shouldn't be exposing repos here directly; refactor to services only
	EventsRepo        repos.EventsRepo
	ItemsRepo         repos.ItemsRepo
	ComparisonService services.ComparisonService
	SubmissionService services.SubmissionService
}

func NewComparison(cfg *config.Config, db shared.DBWithTxStarter, uploadService services.UploadService, verificationService services.VerificationService) *Comparison {
	eventsRepo := repos.NewEventsRepo()
	eventsService := services.NewEventsService(eventsRepo)

	itemsRepo := repos.NewItemsRepo()
	comparisonService := services.NewComparisonService(itemsRepo, eventsService)

	submissionRepo := repos.NewSubmissionRepo()
	submissionService := services.NewSubmissionService(uploadService, submissionRepo, verificationService, cfg.AppConfig.TotalDataLimit)

	return &Comparison{
		EventsRepo:        eventsRepo,
		ItemsRepo:         itemsRepo,
		ComparisonService: comparisonService,
		SubmissionService: submissionService,
	}
}
