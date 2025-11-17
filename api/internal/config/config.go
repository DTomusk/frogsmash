package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig      *AppConfig
	DatabaseConfig *DatabaseConfig
	MailConfig     *MailConfig
	StorageConfig  *StorageConfig
	TokenConfig    *TokenConfig
}

func NewConfig() (*Config, error) {
	if os.Getenv("ENV") != "production" {
		// Load .env file in non-production environments (not running in docker)
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	}

	cfg := &Config{
		AppConfig:      NewAppConfig(),
		DatabaseConfig: NewDatabaseConfig(),
		MailConfig:     NewMailConfig(),
		StorageConfig:  NewStorageConfig(),
		TokenConfig:    NewTokenConfig(),
	}

	return cfg, nil
}
