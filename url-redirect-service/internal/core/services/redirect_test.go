package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/iton0/duss/url-redirect-service/internal/infrastructure/storage"
)

// MockStorage is a mock implementation of the Storage interface.
type MockStorage struct {
	ReturnURL string
	ReturnErr error
}

func (m *MockStorage) Get(ctx context.Context, key string) (string, error) {
	return m.ReturnURL, m.ReturnErr
}

func (m *MockStorage) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return m.ReturnErr
}

func (m *MockStorage) Close() error {
	return nil
}

func TestGetOriginalURL(t *testing.T) {
	ctx := context.Background()

	t.Run("Success - Found", func(t *testing.T) {
		mockStorage := &MockStorage{ReturnURL: "http://example.com/long-url", ReturnErr: nil}
		redirectService := NewRedirectService(mockStorage)

		longURL, err := redirectService.GetOriginalURL(ctx, "short-key")
		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}
		if longURL != "http://example.com/long-url" {
			t.Errorf("expected 'http://example.com/long-url', but got %s", longURL)
		}
	})

	t.Run("Not Found - Key Not Found", func(t *testing.T) {
		mockStorage := &MockStorage{ReturnURL: "", ReturnErr: storage.ErrNotFound}
		redirectService := NewRedirectService(mockStorage)

		_, err := redirectService.GetOriginalURL(ctx, "nonexistent-key")
		if !errors.Is(err, ErrURLNotFound) {
			t.Errorf("expected ErrURLNotFound, but got %v", err)
		}
	})

	t.Run("Error - Generic Storage Error", func(t *testing.T) {
		expectedErr := errors.New("connection failed")
		mockStorage := &MockStorage{ReturnURL: "", ReturnErr: expectedErr}
		redirectService := NewRedirectService(mockStorage)

		_, err := redirectService.GetOriginalURL(ctx, "any-key")
		if err == nil {
			t.Errorf("expected an error, but got nil")
		}
		if !errors.Is(err, expectedErr) {
			t.Errorf("expected a wrapped error containing %v, but got %v", expectedErr, err)
		}
	})
}
