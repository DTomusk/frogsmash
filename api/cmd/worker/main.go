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
	c, err := container.NewContainer(cfg, ctx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	workerContainer := container.NewWorkerContainer(c)

	messageService := workerContainer.InfraServices.MessageService
	messageService.SetUpAndRunWorker(ctx)
}
