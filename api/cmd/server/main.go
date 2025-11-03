package main

import (
	"context"
	"flag"
	"frogsmash/internal/config"
	"frogsmash/internal/container"
	appHttp "frogsmash/internal/delivery/http"
	"log"
	"net/http"
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

	r := appHttp.SetupRoutes(c)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// Server runs in a goroutine
	go func() {
		log.Printf("Starting server on port %s\n", cfg.Port)
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
