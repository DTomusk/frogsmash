package container

import (
	"frogsmash/internal/app/auth/repos"
	"frogsmash/internal/app/auth/services"
	verification "frogsmash/internal/app/verification/services"
	"frogsmash/internal/config"
)

type Auth struct {
	AuthService services.AuthService
	JwtService  services.TokenService
}

func NewAuth(cfg *config.Config, userService services.UserService, verificationService verification.VerificationService) *Auth {
	refreshTokenRepo := repos.NewRefreshTokenRepo()
	hasher := services.NewBCryptHasher()
	jwtService := services.NewJwtService([]byte(cfg.TokenConfig.JWTSecret), cfg.TokenConfig.TokenLifetimeMinutes)

	authService := services.NewAuthService(
		refreshTokenRepo,
		hasher,
		jwtService,
		userService,
		verificationService,
		cfg.TokenConfig.RefreshTokenLifetimeDays,
	)

	return &Auth{
		AuthService: authService,
		JwtService:  jwtService,
	}
}
