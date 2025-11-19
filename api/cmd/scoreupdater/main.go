package main

import (
	"context"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := container.NewContainer(cfg, ctx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	// Handle graceful shutdown
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the score updater
	go c.Comparison.ScoreUpdater.Run(ctx)
	log.Println("ScoreUpdater running...")

	<-quitCh
	log.Println("Shutting down ScoreUpdater...")
	cancel()
}
