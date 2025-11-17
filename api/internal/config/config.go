package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppURL                          string
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
	MailjetAPIKey                   string
	MailjetSecretKey                string
	SenderEmail                     string
	TemplateGlobPattern             string
}

func NewConfig() (*Config, error) {
	if os.Getenv("ENV") != "production" {
		// Load .env file in non-production environments (not running in docker)
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	cfg := &Config{
		AppURL:                          getEnv("APP_URL"),
		DatabaseURL:                     getEnv("DATABASE_URL"),
		Port:                            getEnv("PORT"),
		AllowedOrigin:                   getEnv("ALLOWED_ORIGIN"),
		KFactor:                         getFloat("KFACTOR"),
		MaxFileSize:                     getInt64("MAX_FILE_SIZE_MB") * 1024 * 1024, // Convert MB to bytes
		ScoreUpdateInterval:             getInt("SCORE_UPDATE_INTERVAL_SECONDS"),
		StorageAccountID:                getEnv("STORAGE_ACCOUNT_ID"),
		StorageAccessKey:                getEnv("STORAGE_ACCESS_KEY"),
		StorageSecretKey:                getEnv("STORAGE_SECRET_KEY"),
		StorageBucket:                   getEnv("STORAGE_BUCKET"),
		JWTSecret:                       getEnv("JWT_SECRET"),
		TokenLifetimeMinutes:            getInt("JWT_TOKEN_LIFETIME_MINUTES"),
		RefreshTokenLifetimeDays:        getInt("REFRESH_TOKEN_LIFETIME_DAYS"),
		VerificationCodeLength:          getInt("VERIFICATION_CODE_LENGTH"),
		VerificationCodeLifetimeMinutes: getInt("VERIFICATION_CODE_LIFETIME_MINUTES"),
		MailjetAPIKey:                   getEnv("MAILJET_API_KEY"),
		MailjetSecretKey:                getEnv("MAILJET_SECRET_KEY"),
		SenderEmail:                     getEnv("SENDER_EMAIL"),
		TemplateGlobPattern:             getEnv("TEMPLATE_GLOB_PATTERN"),
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("could not load DATABASE_URL from environment")
	}
	if cfg.Port == "" {
		return nil, fmt.Errorf("could not load PORT from environment")
	}

	return cfg, nil
}
