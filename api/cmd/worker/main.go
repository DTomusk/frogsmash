package main

import (
	"context"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// TODO: consider lighter version of container for workers
	c, err := container.NewBaseContainer(cfg, ctx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	workerContainer, err := container.NewWorkerContainer(c, ctx)
	if err != nil {
		log.Fatalf("Failed to create worker container: %v", err)
	}

	messageConsumer := workerContainer.MessageConsumer
	messageConsumer.SetUpAndRunWorker(ctx)
}
