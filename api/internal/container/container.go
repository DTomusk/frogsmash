package container

import (
	"context"
	"frogsmash/internal/config"
)

type Container struct {
	Auth          *Auth
	Comparison    *Comparison
	Config        *config.Config
	Verification  *Verification
	User          *User
	InfraServices *InfraServices
}

func NewContainer(cfg *config.Config, ctx context.Context) (*Container, error) {
	infraServices, err := NewInfraServices(cfg, ctx)
	if err != nil {
		return nil, err
	}

	user := NewUser(cfg)
	verification := NewVerification(cfg, user.UserService, infraServices.EmailService)
	// TODO: pass in infra queue instead so we can queue verification emails instead of sending them directly
	auth := NewAuth(cfg, user.UserService, verification.VerificationService)
	comparison := NewComparison(cfg, infraServices.DB)

	return &Container{
		Auth:          auth,
		Comparison:    comparison,
		InfraServices: infraServices,
		Verification:  verification,
		User:          user,
		Config:        cfg,
	}, nil
}
