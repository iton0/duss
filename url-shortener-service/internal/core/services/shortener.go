package services

import (
	"context"
	"time"

	"github.com/iton0/duss/shared/domain"
)

// ShortenerService encapsulates the business logic.
type ShortenerService struct {
	storage ports.URLStorage
}

func (s *ShortenerService) Shorten(ctx context.Context, longURL string) (*domain.URL, error) {
	// Generate a unique short key here (e.g., call your key-gen-service)
	shortKey := generateUniqueKey()

	// Create the domain.URL entity. This is the heart of the business logic.
	newURL := &domain.URL{
		ShortKey:  shortKey,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		Redirects: 0,
	}

	// Pass the domain.URL entity to the storage layer to be persisted.
	err := s.storage.Save(ctx, newURL)
	if err != nil {
		return nil, err
	}

	return newURL, nil
}
