package mock

import (
	"context"
	"errors"
)

// ErrNotFound is a sentinel error for when a key is not found,
// matching the expected error from a real storage implementation.
var ErrNotFound = errors.New("key not found")

// MockStorage is a mock implementation of the Storage interface.
type MockStorage struct {
	data          map[string]string
	simulateError bool
}

// NewMockStorage creates a new MockStorage instance.
func NewMockStorage(initialData map[string]string) *MockStorage {
	if initialData == nil {
		initialData = make(map[string]string)
	}
	return &MockStorage{
		data: initialData,
	}
}

// Get simulates retrieving a value from the "database".
func (m *MockStorage) Get(ctx context.Context, shortKey string) (string, error) {
	if m.simulateError {
		return "", errors.New("mock storage connection error")
	}

	longURL, ok := m.data[shortKey]
	if !ok {
		// Return the specific error your service expects.
		return "", ErrNotFound
	}
	return longURL, nil
}
