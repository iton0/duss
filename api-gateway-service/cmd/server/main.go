package main

import (
	"log"
	"os"

	"github.com/iton0/duss/api-gateway-service/internal/api"
	"github.com/iton0/duss/api-gateway-service/internal/core/services"
	"github.com/iton0/duss/api-gateway-service/internal/infrastructure/clients"
	"github.com/iton0/duss/api-gateway-service/internal/infrastructure/web"
)

func main() {
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
	log.Println("Starting API Gateway on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
