package main

import (
	"flag"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/http"
	"log"

	_ "frogsmash/docs"

	_ "github.com/lib/pq"
)

// @title Frog Smash API
// @version 1.0
// @description The API for comparing frogs and other things
func main() {
	verbose := flag.Bool("v", false, "enable verbose output")
	flag.Parse()
	if *verbose {
		log.Println("Verbose mode enabled")
	}
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	c, err := container.NewContainer(cfg)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	// Start background score updater
	go c.ScoreUpdater.Run()

	r := http.SetupRoutes(c)
	r.Run(":8080")
}
