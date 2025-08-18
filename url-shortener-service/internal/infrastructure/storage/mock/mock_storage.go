package mock

import (
	"context"
	"errors"

	"github.com/iton0/duss/shared/domain"
)

// ErrDuplicatedKey is a sentinel error for when a key already exists,
// matching the expected error from a real storage implementation.
var ErrDuplicatedKey = errors.New("URL already taken")

// MockPostgresStorage is a mock implementation of the Storage interface for the url-shortener-service.
type MockPostgresStorage struct {
	// The map stores long URLs, with short keys as keys.
	data          map[string]*domain.URL
	simulateError bool
}

// NewMockPostgresStorage creates a new MockPostgresStorage instance.
func NewMockPostgresStorage() *MockPostgresStorage {
	return &MockPostgresStorage{
		data: make(map[string]*domain.URL),
	}
}

// Save simulates saving a URL to the "database".
func (m *MockPostgresStorage) Save(ctx context.Context, url *domain.URL) error {
	if m.simulateError {
		return errors.New("mock storage save error")
	}

	// Check for a duplicate key before saving.
	if _, ok := m.data[url.ShortKey]; ok {
		return ErrDuplicatedKey
	}

	m.data[url.ShortKey] = url
	return nil
}

// Get is provided for testing convenience, though it's not part of the primary Storage interface for this service.
func (m *MockPostgresStorage) Get(ctx context.Context, shortKey string) (*domain.URL, error) {
	url, ok := m.data[shortKey]
	if !ok {
		return nil, errors.New("key not found")
	}
	return url, nil
}
