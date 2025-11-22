package container

import (
	"frogsmash/internal/app/verification/repos"
	"frogsmash/internal/app/verification/services"
	"frogsmash/internal/config"
)

type Verification struct {
	VerificationService services.VerificationService
}

func NewVerification(cfg *config.Config, userService services.UserService, emailService services.EmailService) *Verification {
	verificationRepo := repos.NewVerificationRepo()
	verificationService := services.NewVerificationService(
		userService,
		verificationRepo,
		emailService,
		cfg.TokenConfig.VerificationCodeLength,
		cfg.TokenConfig.VerificationCodeLifetimeMinutes,
	)

	return &Verification{
		VerificationService: verificationService,
	}
}
