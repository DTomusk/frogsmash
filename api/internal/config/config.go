package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	if os.Getenv("ENV") != "production" {
		// Load .env file in non-production environments (not running in docker)
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	kFactor, err := strconv.ParseFloat(os.Getenv("KFACTOR"), 64)
	if err != nil {
		return nil, fmt.Errorf("could not load KFACTOR from environment: %v", err)
	}

	scoreUpdateInterval, err := strconv.Atoi(os.Getenv("SCORE_UPDATE_INTERVAL_SECONDS"))
	if err != nil {
		return nil, fmt.Errorf("could not load SCORE_UPDATE_INTERVAL_SECONDS from environment: %v", err)
	}

	maxFileSzie, err := strconv.ParseInt(os.Getenv("MAX_FILE_SIZE_MB"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("could not load MAX_FILE_SIZE_MB from environment: %v", err)
	}

	jwtLifetimeMinutes, err := strconv.Atoi(os.Getenv("JWT_TOKEN_LIFETIME_MINUTES"))
	if err != nil {
		return nil, fmt.Errorf("could not load JWT_TOKEN_LIFETIME_MINUTES from environment: %v", err)
	}

	refreshTokenLifetimeDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_LIFETIME_DAYS"))
	if err != nil {
		return nil, fmt.Errorf("could not load REFRESH_TOKEN_LIFETIME_DAYS from environment: %v", err)
	}

	verificationCodeLength, err := strconv.Atoi(os.Getenv("VERIFICATION_CODE_LENGTH"))
	if err != nil {
		return nil, fmt.Errorf("could not load VERIFICATION_CODE_LENGTH from environment: %v", err)
	}

	verificationCodeLifetimeMinutes, err := strconv.Atoi(os.Getenv("VERIFICATION_CODE_LIFETIME_MINUTES"))
	if err != nil {
		return nil, fmt.Errorf("could not load VERIFICATION_CODE_LIFETIME_MINUTES from environment: %v", err)
	}

	cfg := &Config{
		DatabaseURL:                     os.Getenv("DATABASE_URL"),
		Port:                            os.Getenv("PORT"),
		AllowedOrigin:                   os.Getenv("ALLOWED_ORIGIN"),
		KFactor:                         kFactor,
		MaxFileSize:                     maxFileSzie * 1024 * 1024, // Convert MB to bytes
		ScoreUpdateInterval:             scoreUpdateInterval,
		StorageAccountID:                os.Getenv("STORAGE_ACCOUNT_ID"),
		StorageAccessKey:                os.Getenv("STORAGE_ACCESS_KEY"),
		StorageSecretKey:                os.Getenv("STORAGE_SECRET_KEY"),
		StorageBucket:                   os.Getenv("STORAGE_BUCKET"),
		JWTSecret:                       os.Getenv("JWT_SECRET"),
		TokenLifetimeMinutes:            jwtLifetimeMinutes,
		RefreshTokenLifetimeDays:        refreshTokenLifetimeDays,
		VerificationCodeLength:          verificationCodeLength,
		VerificationCodeLifetimeMinutes: verificationCodeLifetimeMinutes,
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("could not load DATABASE_URL from environment")
	}
	if cfg.Port == "" {
		return nil, fmt.Errorf("could not load PORT from environment")
	}

	return cfg, nil
}

type Config struct {
	DatabaseURL                     string
	Port                            string
	KFactor                         float64
	ScoreUpdateInterval             int
	MaxFileSize                     int64
	AllowedOrigin                   string
	StorageAccountID                string
	StorageAccessKey                string
	StorageSecretKey                string
	StorageBucket                   string
	JWTSecret                       string
	TokenLifetimeMinutes            int
	RefreshTokenLifetimeDays        int
	VerificationCodeLength          int
	VerificationCodeLifetimeMinutes int
}
