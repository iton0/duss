package services

import (
	"errors"
)

// TODO: what type of error should i create here
var _ = errors.New("")

type KeygenServiceIface interface {
	GenerateKey() (string, error)
}

// This is a compile-time check to ensure the contract is fulfilled.
var _ KeygenServiceIface = (*KeygenService)(nil)

type KeygenService struct{}

// NewKeygenService creates a new KeygenService instance.
func NewKeygenService() *KeygenService {
	return &KeygenService{}
}

// GenerateKey creates a cryptographic short key.
func (s *KeygenService) GenerateKey() (string, error) {
	// TODO: implement gnereate key logic
}
