package services

import (
	"context"
	"errors"
	"log"

	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/storage"
)

// ErrURLNotFound indicates that the short key was not found.
var ErrURLNotFound = errors.New("URL not found")

// RedirectServiceIface defines the behavior of the redirect service.
type RedirectServiceIface interface {
	GetOriginalURL(ctx context.Context, shortKey string) (string, error)
}

// Ensure RedirectService explicitly implements RedirectServiceIface.
// This is a compile-time check to ensure the contract is fulfilled.
var _ RedirectServiceIface = (*RedirectService)(nil)

// RedirectService encapsulates the core logic for URL redirection.
// This struct now implicitly implements the RedirectServiceIface.
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
		switch {
		case errors.Is(err, storage.ErrNotFound):
			return "", ErrURLNotFound
		default:
			log.Printf("unexpected server error: %v", err)
			return "", err
		}
	}
	return longURL, nil
}
