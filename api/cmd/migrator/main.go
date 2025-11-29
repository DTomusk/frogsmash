package main

import (
	"flag"
	"log"

	"frogsmash/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	m, err := migrate.New(
		"file://db/migrations",
		cfg.DatabaseConfig.DatabaseURL,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if *verbose {
		log.Println("Running migrations...")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration error: %v", err)
	}

	log.Println("Migrations applied successfully.")
}
