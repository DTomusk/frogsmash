package container

import (
	"context"
	"frogsmash/internal/config"
)

type BaseContainer struct {
	Comparison    *Comparison
	Config        *config.Config
	Verification  *Verification
	User          *User
	InfraServices *InfraServices
}

func NewBaseContainer(cfg *config.Config, ctx context.Context) (*BaseContainer, error) {
	infraServices, err := NewInfraServices(cfg, ctx)
	if err != nil {
		return nil, err
	}

	user := NewUser(cfg)
	verification := NewVerification(cfg, user.UserService, infraServices.EmailService)
	comparison := NewComparison(cfg, infraServices.DB, infraServices.UploadService, verification.VerificationService)

	return &BaseContainer{
		Comparison:    comparison,
		InfraServices: infraServices,
		Verification:  verification,
		User:          user,
		Config:        cfg,
	}, nil
}
