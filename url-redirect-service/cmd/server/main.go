package main

import (
	"context"
	"log"
	"time"

	"github.com/iton0/duss/url-redirect-service/internal/api"
	"github.com/iton0/duss/url-redirect-service/internal/core/services"
	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/storage"
	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/web"
)

func main() {
	// A new context is created for the application's lifecycle.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Initialize the concrete Redis client.
	// TODO: update this to the actual redis implementation
	redisClient, err := storage.NewRedisClient(ctx, "localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}

	// 2. Initialize the core service.
	redirectService := services.NewRedirectService(redisClient)

	// 3. Initialize the API handler.
	redirectHandler := api.NewRedirectHandler(redirectService)

	// 4. Initialize the router and register the handler.
	router := web.NewRouter(redirectHandler)

	log.Println("Starting redirect service on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
