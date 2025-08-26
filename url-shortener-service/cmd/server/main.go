package main

import (
	"context"
	"time"

	"github.com/iton0/duss/url-shortener-service/internal/api"
	"github.com/iton0/duss/url-shortener-service/internal/core/services"
	"github.com/iton0/duss/url-shortener-service/internal/infrastructure/storage"
	"github.com/iton0/duss/url-shortener-service/internal/infrastructure/web"
)

func main() {
	// A new context is created for the application's lifecycle.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// INitilaize sotrage client

	// intiialize core serivce
	// 2. Initialize the core service.

	// initalie ape handler
	// 3. Initialize the API handler.

	// intialize router  and register handler
	// 4. Initialize the router and register the handler.
}
