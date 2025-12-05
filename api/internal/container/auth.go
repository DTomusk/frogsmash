package container

import (
	"frogsmash/internal/app/auth/repos"
	"frogsmash/internal/app/auth/services"
	"frogsmash/internal/config"
)

type Auth struct {
	AuthService services.AuthService
	JwtService  services.TokenService
}

func NewAuth(cfg *config.Config, userService services.UserService, messageService services.MessageProducer) *Auth {
	refreshTokenRepo := repos.NewRefreshTokenRepo()
	hasher := services.NewBCryptHasher()
	jwtService := services.NewJwtService([]byte(cfg.TokenConfig.JWTSecret), cfg.TokenConfig.TokenLifetimeMinutes)
	googleService := services.NewGoogleService(cfg.AppConfig.GoogleClientID)

	authService := services.NewAuthService(
		refreshTokenRepo,
		hasher,
		jwtService,
		userService,
		messageService,
		googleService,
		cfg.TokenConfig.RefreshTokenLifetimeDays,
	)

	return &Auth{
		AuthService: authService,
		JwtService:  jwtService,
	}
}
