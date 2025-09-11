package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/btcsuite/btcd/btcutil/base58"
)

var ErrInvalidURL = errors.New("url cannot be empty")

type KeygenServiceIface interface {
	GenerateKey(url string) (string, error)
}

// This is a compile-time check to ensure the contract is fulfilled.
var _ KeygenServiceIface = (*KeygenService)(nil)

type KeygenService struct{}

// NewKeygenService creates a new KeygenService instance.
func NewKeygenService() *KeygenService {
	return &KeygenService{}
}

// GenerateKey creates a cryptographic short key.
func (s *KeygenService) GenerateKey(url string) (string, error) {
	if url == "" {
		return "", ErrInvalidURL
	}

	salt := fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(10000))
	data := url + salt

	h := sha256.New()

	if _, err := io.WriteString(h, data); err != nil {
		return "", ErrInvalidURL
	}

	hashBytes := h.Sum(nil)

	// NOTE: adjust the length to control the key's size and collision rate
	truncatedHash := hashBytes[:8]

	// Base58 encode the truncated hash to make it more compact and URL-safe
	encodedKey := base58.Encode(truncatedHash)

	return encodedKey, nil
}
