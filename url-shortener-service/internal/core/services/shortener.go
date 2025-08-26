package services

import (
	"context"
	"errors"
	"time"

	"github.com/iton0/duss/shared/domain"
	"github.com/iton0/duss/url-shortener-service/internal/infrastructure/storage"
)

var (
	ErrInvalidURL     = errors.New("URL not valid")
	ErrBlacklistedURL = errors.New("URL rejected")
	ErrDuplicatedKey  = errors.New("URL already taken")
)

// ShortenerService encapsulates the business logic.
type ShortenerServiceIface interface {
	Shorten(ctx context.Context, longURL string) (*domain.URL, error)
}

var _ ShortenerServiceIface = (*ShortenerService)(nil)

type ShortenerService struct {
	storage storage.Storage
}

func NewShortenerService(s storage.Storage) *ShortenerService {
	return &ShortenerService{storage: s}
}

func (s *ShortenerService) Shorten(ctx context.Context, longURL string) (*domain.URL, error) {
	// Generate a unique short key here (e.g., call your key-gen-service)
	// TODO: needs to call from the keygen service
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
