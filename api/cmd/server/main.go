package main

import (
	"flag"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/http"
	"log"

	_ "github.com/lib/pq"
)

// Entry point
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

	r := http.SetupRoutes(c)
	r.Run(":8080")
}
