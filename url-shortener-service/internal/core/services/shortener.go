package services

import (
	"context"
	"errors"
	"log"
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
	// TODO: needs to call from the keygen service
	shortKey := generateUniqueKey()

	// create the domain.URL entity
	newURL := &domain.URL{
		ShortKey:  shortKey,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		Redirects: 0,
	}

	// Pass the domain.URL entity to the storage layer to be persisted.
	err := s.storage.Save(ctx, newURL)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidURL):
		case errors.Is(err, ErrBlacklistedURL):
		case errors.Is(err, ErrDuplicatedKey):
		default:
			log.Printf("unexpected server error: %v", err)
			return nil, err
		}
	}

	return newURL, nil
}
