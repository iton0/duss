package services

import (
	"context"
	"errors"
	"log"

	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/storage"
)

// ErrURLNotFound indicates that the short key was not found.
var ErrURLNotFound = errors.New("URL not found")

// RedirectService encapsulates the core logic for URL redirection.
type RedirectService struct {
	storage storage.Storage
}

// NewRedirectService creates a new RedirectService instance.
func NewRedirectService(s storage.Storage) *RedirectService {
	return &RedirectService{storage: s}
}

// GetOriginalURL retrieves the long URL for a given short key.
func (s *RedirectService) GetOriginalURL(ctx context.Context, shortKey string) (string, error) {
	longURL, err := s.storage.Get(ctx, shortKey)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return "", ErrURLNotFound
		}
		// Log the error here before returning it.
		log.Printf("storage error: %v", err)
		return "", err
	}
	return longURL, nil
}
