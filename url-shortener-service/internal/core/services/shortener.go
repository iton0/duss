package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/iton0/duss/shared/domain"
	"github.com/iton0/duss/url-shortener-service/internal/infrastructure/storage"
)

var (
	ErrInvalidURL     = errors.New("URL not valid")
	ErrBlacklistedURL = errors.New("URL rejected")
	ErrDuplicatedKey  = errors.New("URL already taken")
	ErrServiceError   = errors.New("service error")
)

// ShortenerService encapsulates the business logic.
type ShortenerServiceIface interface {
	Shorten(ctx context.Context, longURL string) (*domain.URL, error)
}

var _ ShortenerServiceIface = (*ShortenerService)(nil)

type ShortenerService struct {
	storage       storage.Storage
	keyGenService string // The URL of the key-gen-service
}

func NewShortenerService(s storage.Storage, keyGenServiceURL string) *ShortenerService {
	return &ShortenerService{
		storage:       s,
		keyGenService: keyGenServiceURL,
	}
}

// KeyGenResponse is the structure for the response from the key-gen-service.
type KeyGenResponse struct {
	ShortKey string `json:"short_key"`
}

func (s *ShortenerService) Shorten(ctx context.Context, longURL string) (*domain.URL, error) {
	// Call the key-gen-service to get a unique key.
	shortKey, err := s.getUniqueKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get unique key: %w", err)
	}

	// create the domain.URL entity
	newURL := &domain.URL{
		ShortKey:  shortKey,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		Redirects: 0,
	}

	// Pass the domain.URL entity to the storage layer to be persisted.
	err = s.storage.Save(ctx, newURL)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidURL):
			return nil, ErrInvalidURL
		case errors.Is(err, ErrBlacklistedURL):
			return nil, ErrBlacklistedURL
		case errors.Is(err, ErrDuplicatedKey):
			return nil, ErrDuplicatedKey
		default:
			log.Printf("unexpected server error: %v", err)
			return nil, ErrServiceError // Return a generic service error to the caller
		}
	}

	return newURL, nil
}

// getUniqueKey makes an HTTP request to the key-gen-service to obtain a unique key.
func (s *ShortenerService) getUniqueKey(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.keyGenService+"/api/v1/generate-key", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call key-gen-service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("key-gen-service returned non-OK status: %d", resp.StatusCode)
	}

	var keyResp KeyGenResponse
	if err := json.NewDecoder(resp.Body).Decode(&keyResp); err != nil {
		return "", fmt.Errorf("failed to decode key-gen-service response: %w", err)
	}

	if keyResp.ShortKey == "" {
		return "", errors.New("key-gen-service returned an empty key")
	}

	return keyResp.ShortKey, nil
}
