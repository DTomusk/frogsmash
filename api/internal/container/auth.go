package container

import (
	"frogsmash/internal/app/auth/factories"
	"frogsmash/internal/app/auth/repos"
	"frogsmash/internal/app/auth/services"
	"frogsmash/internal/config"
)

type Auth struct {
	AuthService         *services.AuthService
	JwtService          *services.JwtService
	UserService         *services.UserService
	VerificationService *services.VerificationService
}

func NewAuth(cfg *config.Config, emailService services.EmailService) *Auth {
	userRepo := repos.NewUserRepo()
	refreshTokenRepo := repos.NewRefreshTokenRepo()
	hasher := services.NewBCryptHasher()
	jwtService := services.NewJwtService([]byte(cfg.TokenConfig.JWTSecret), cfg.TokenConfig.TokenLifetimeMinutes)
	verificationRepo := repos.NewVerificationRepo()

	authService := services.NewAuthService(
		userRepo,
		refreshTokenRepo,
		hasher,
		jwtService,
		cfg.TokenConfig.RefreshTokenLifetimeDays,
	)

	verificationService := services.NewVerificationService(
		userRepo,
		verificationRepo,
		emailService,
		cfg.TokenConfig.VerificationCodeLength,
		cfg.TokenConfig.VerificationCodeLifetimeMinutes,
	)
	userFactory := factories.NewUserFactory(hasher)

	userService := services.NewUserService(userFactory, userRepo, verificationService)

	return &Auth{
		AuthService:         authService,
		JwtService:          jwtService,
		UserService:         userService,
		VerificationService: verificationService,
	}
}
