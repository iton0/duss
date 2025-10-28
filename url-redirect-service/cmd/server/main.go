package main

import (
	"context"
	"log"
	"time"

	"github.com/iton0/duss/url-redirect-service/internal/api"
	"github.com/iton0/duss/url-redirect-service/internal/core/services"
	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/storage"
	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/web"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// TODO: update this to the actual redis implementation
	redisClient, err := storage.NewRedisClient(ctx, "localhost:6379", "", 0)
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}

	redirectService := services.NewRedirectService(redisClient)
	redirectHandler := api.NewRedirectHandler(redirectService)
	router := web.NewRouter(redirectHandler)

	log.Println("Starting redirect service on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
