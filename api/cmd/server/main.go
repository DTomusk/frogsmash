package main

import (
	"context"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	appHttp "frogsmash/internal/delivery/shared/http"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "frogsmash/docs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// @title Frog Smash API
// @version 1.0
// @description The API for comparing frogs and other things
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c, err := container.NewBaseContainer(cfg, appCtx)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	apiContainer := container.NewAPIContainer(c)

	r := appHttp.SetupRoutes(apiContainer)

	srv := &http.Server{
		Addr:    ":" + cfg.AppConfig.Port,
		Handler: r,
	}

	// Server runs in a goroutine
	go func() {
		log.Printf("Starting server on port %s\n", cfg.AppConfig.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	// This blocks until an interrupt or terminate signal is received
	<-quitCh
	log.Println("Shutting down server...")

	cancel()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
