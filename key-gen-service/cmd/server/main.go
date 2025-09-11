package main

import (
	"log"

	"github.com/iton0/duss/key-gen-service/internal/api"
	"github.com/iton0/duss/key-gen-service/internal/core/services"
	"github.com/iton0/duss/key-gen-service/internal/infrastructure/web"
)

func main() {
	keygenService := services.NewKeygenService()
	keygenHandler := api.NewKeygenHandler(keygenService)
	router := web.NewRouter(keygenHandler)

	log.Println("Starting key generator service on :8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
