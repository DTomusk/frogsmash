package container

import (
	"frogsmash/internal/app/comparison/services"
	"frogsmash/internal/config"
	"time"
)

type ScoreUpdaterContainer struct {
	*BaseContainer
	ScoreUpdater services.ScoreUpdater
}

func NewScoreUpdaterContainer(c *BaseContainer, cfg *config.Config) *ScoreUpdaterContainer {
	updateInterval := time.Duration(cfg.AppConfig.ScoreUpdateInterval) * time.Second
	scoreUpdater := services.NewScoreUpdater(
		c.InfraServices.DB,
		c.Comparison.EventsRepo,
		c.Comparison.ItemsRepo,
		cfg.AppConfig.KFactor,
		updateInterval)

	return &ScoreUpdaterContainer{
		BaseContainer: c,
		ScoreUpdater:  scoreUpdater,
	}
}
