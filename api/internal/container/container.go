package container

import (
	"context"
	"frogsmash/internal/config"
)

type Container struct {
	Auth          *Auth
	Comparison    *Comparison
	Config        *config.Config
	InfraServices *InfraServices
}

func NewContainer(cfg *config.Config, ctx context.Context) (*Container, error) {
	infraServices, err := NewInfraServices(cfg, ctx)
	if err != nil {
		return nil, err
	}

	auth := NewAuth(cfg, infraServices.EmailService)
	comparison := NewComparison(cfg, infraServices.DB)

	return &Container{
		Auth:          auth,
		Comparison:    comparison,
		InfraServices: infraServices,
		Config:        cfg,
	}, nil
}
