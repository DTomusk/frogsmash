package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	kFactor, err := strconv.ParseFloat(os.Getenv("KFACTOR"), 64)
	if err != nil {
		return nil, fmt.Errorf("could not load KFACTOR from environment: %v", err)
	}

	scoreUpdateInterval, err := strconv.Atoi(os.Getenv("SCORE_UPDATE_INTERVAL_SECONDS"))
	if err != nil {
		return nil, fmt.Errorf("could not load SCORE_UPDATE_INTERVAL_SECONDS from environment: %v", err)
	}

	cfg := &Config{
		DatabaseURL:         os.Getenv("DATABASE_URL"),
		Port:                os.Getenv("PORT"),
		AllowedOrigin:       os.Getenv("ALLOWED_ORIGIN"),
		KFactor:             kFactor,
		ScoreUpdateInterval: scoreUpdateInterval,
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
	DatabaseURL         string
	Port                string
	KFactor             float64
	ScoreUpdateInterval int
	AllowedOrigin       string
}
