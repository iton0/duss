package services

import (
	"context"
)

// GatewayServiceIface defines the behavior of the gateway service.
type GatewayServiceIface interface {
	ShortenURL(ctx context.Context, originalURL string) (string, error)
	RedirectURL(ctx context.Context, shortURL string) (string, error)
}

// Ensure GatewayService explicitly implements GatewayServiceIface.
// This is a compile-time check to ensure the contract is fulfilled.
var _ GatewayServiceIface = (*GatewayService)(nil)

// GatewayService encapsulates the core logic for URL gatewaying.
type GatewayService struct {
	shortenerClient ShortenerServiceClient
	redirectClient  RedirectServiceClient
}

// These are the client interfaces that the gateway depends on.
type ShortenerServiceClient interface {
	Shorten(ctx context.Context, originalURL string) (string, error)
}

type RedirectServiceClient interface {
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
}

// ShortenURL implements the GatewayServiceIface.
func (s *GatewayService) ShortenURL(ctx context.Context, originalURL string) (string, error) {
	return s.shortenerClient.Shorten(ctx, originalURL)
}

// RedirectURL implements the GatewayServiceIface.
func (s *GatewayService) RedirectURL(ctx context.Context, shortURL string) (string, error) {
	return s.redirectClient.GetOriginalURL(ctx, shortURL)
}

// NewGatewayService creates a new GatewayService instance with its dependencies.
func NewGatewayService(shortenerClient ShortenerServiceClient, redirectClient RedirectServiceClient) *GatewayService {
	return &GatewayService{
		shortenerClient: shortenerClient,
		redirectClient:  redirectClient,
	}
}
