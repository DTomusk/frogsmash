package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
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
	DatabaseURL string
	Port        string
}
