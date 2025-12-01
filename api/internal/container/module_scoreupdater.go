package container

import (
	"frogsmash/internal/app/comparison/services"
	"frogsmash/internal/config"
	"time"
)

type ScoreUpdaterContainer struct {
	*Container
	ScoreUpdater services.ScoreUpdater
}

func NewScoreUpdaterContainer(c *Container, cfg *config.Config) *ScoreUpdaterContainer {
	updateInterval := time.Duration(cfg.AppConfig.ScoreUpdateInterval) * time.Second
	scoreUpdater := services.NewScoreUpdater(
		c.InfraServices.DB,
		c.Comparison.EventsRepo,
		c.Comparison.ItemsRepo,
		cfg.AppConfig.KFactor,
		updateInterval)

	return &ScoreUpdaterContainer{
		Container:    c,
		ScoreUpdater: scoreUpdater,
	}
}
