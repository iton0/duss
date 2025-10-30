package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iton0/duss/api-gateway-service/internal/api"
	"github.com/iton0/duss/api-gateway-service/internal/core/services"
	"github.com/iton0/duss/api-gateway-service/internal/infrastructure/clients"
	"github.com/iton0/duss/api-gateway-service/internal/infrastructure/web"
	"github.com/joho/godotenv"
)

// TODO: need to add api authentication for the gateway
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	hostPort := os.Getenv("PUBLIC_GATEWAY_PORT")

	// 1. Get configuration from environment variables.
	shortenerServiceURL := os.Getenv("SHORTENER_SERVICE_URL")
	if shortenerServiceURL == "" {
		log.Fatal("SHORTENER_SERVICE_URL environment variable is not set")
	}

	redirectServiceURL := os.Getenv("REDIRECT_SERVICE_URL")
	if redirectServiceURL == "" {
		log.Fatal("REDIRECT_SERVICE_URL environment variable is not set")
	}

	// 2. Initialize the clients for the other services.
	// These clients know how to communicate over the network.
	shortenerClient := clients.NewHTTPShortenerClient(shortenerServiceURL)
	redirectClient := clients.NewHTTPRedirectClient(redirectServiceURL)

	// 3. Initialize the core gateway service.
	// This service contains the business logic for routing and delegates tasks to the clients.
	gatewayService := services.NewGatewayService(shortenerClient, redirectClient)

	// 4. Initialize the API handler.
	// The handler receives requests and uses the gateway service to fulfill them.
	gatewayHandler := api.NewGatewayHandler(gatewayService)

	// 5. Initialize the router.
	// The router maps public-facing URL paths to the API handlers.
	router := web.NewRouter(gatewayHandler)

	// 6. Start the server.
	log.Printf("Starting API Gateway on port %s...\n", hostPort)
	if err := router.Run(fmt.Sprintf(":%s", hostPort)); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
