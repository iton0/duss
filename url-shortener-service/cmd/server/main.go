package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iton0/duss/url-shortener-service/internal/api"
	"github.com/iton0/duss/url-shortener-service/internal/core/services"
	"github.com/iton0/duss/url-shortener-service/internal/infrastructure/storage"
	"github.com/iton0/duss/url-shortener-service/internal/infrastructure/web"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// A new context is created for the application's lifecycle.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Initialize storage client.
	// In a real app, DSN would come from a config file or environment variable.
	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		postgresDSN = "postgres://user:password@localhost:5432/duss?sslmode=disable"
	}

	pgStore, err := storage.NewPostgresClient(ctx, postgresDSN)
	if err != nil {
		log.Fatalf("failed to initialize PostgreSQL client: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL")

	// 2. Initialize the core service.
	// Key-gen-service URL would also come from a config.
	keyGenServiceURL := os.Getenv("KEY_GEN_SERVICE_URL")
	if keyGenServiceURL == "" {
		keyGenServiceURL = "http://localhost:8082" // Default URL for the key-gen-service
	}

	shortenerService := services.NewShortenerService(pgStore, keyGenServiceURL)

	// 3. Initialize the API handler.
	shortenerHandler := api.NewShortenerHandler(shortenerService)

	// 4. Initialize the router and register the handler.
	router := web.NewRouter(shortenerHandler)

	// 5. Start the server.
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: router,
	}

	// Graceful shutdown logic.
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
