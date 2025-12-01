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
	auth := NewAuth(cfg, user.UserService, infraServices.MessageClient)
	comparison := NewComparison(cfg, infraServices.DB, infraServices.UploadService, verification.VerificationService)

	infraServices.Dispatcher.RegisterHandler("send_verification_email", verification.VerificationService)

	return &Container{
		Auth:          auth,
		Comparison:    comparison,
		InfraServices: infraServices,
		Verification:  verification,
		User:          user,
		Config:        cfg,
	}, nil
}
