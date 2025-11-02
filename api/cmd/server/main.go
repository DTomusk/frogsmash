package main

import (
	"context"
	"flag"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	"frogsmash/internal/delivery/http"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "frogsmash/docs"

	_ "github.com/lib/pq"
)

// @title Frog Smash API
// @version 1.0
// @description The API for comparing frogs and other things
func main() {
	// TODO: use elsewhere
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go c.ScoreUpdater.Run(ctx)

	// Listen for termination signals to gracefully shut down background services
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	r := http.SetupRoutes(c)
	r.Run(":8080")
}
