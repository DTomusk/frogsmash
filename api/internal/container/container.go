package container

import (
	"database/sql"
	"frogsmash/internal/config"
)

type Container struct {
	DB *sql.DB
}

func NewContainer(cfg *config.Config) (*Container, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Container{
		DB: db,
	}, nil
}
