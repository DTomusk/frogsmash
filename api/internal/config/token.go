package config

type TokenConfig struct {
	JWTSecret                       string
	TokenLifetimeMinutes            int
	RefreshTokenLifetimeDays        int
	VerificationCodeLength          int
	VerificationCodeLifetimeMinutes int
}

func NewTokenConfig() *TokenConfig {
	return &TokenConfig{
		JWTSecret:                       getEnv("JWT_SECRET"),
		TokenLifetimeMinutes:            getInt("JWT_TOKEN_LIFETIME_MINUTES"),
		RefreshTokenLifetimeDays:        getInt("REFRESH_TOKEN_LIFETIME_DAYS"),
		VerificationCodeLength:          getInt("VERIFICATION_CODE_LENGTH"),
		VerificationCodeLifetimeMinutes: getInt("VERIFICATION_CODE_LIFETIME_MINUTES"),
	}
}
