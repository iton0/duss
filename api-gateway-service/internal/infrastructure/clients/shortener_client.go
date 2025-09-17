package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iton0/duss/api-gateway-service/internal/core/services"
)

// HTTPShortenerClient is a concrete implementation of the ShortenerServiceClient interface.
type HTTPShortenerClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPShortenerClient creates a new HTTP client for the shortening service.
func NewHTTPShortenerClient(baseURL string) services.ShortenerServiceClient {
	return &HTTPShortenerClient{
		client:  &http.Client{Timeout: 5 * time.Second},
		baseURL: baseURL,
	}
}

// shortenRequest mirrors the expected JSON structure of the shortener service.
type shortenRequest struct {
	URL string `json:"url"`
}

// shortenResponse mirrors the expected JSON structure of the shortener service.
type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

// Shorten sends an HTTP POST request to the shortening service.
func (c *HTTPShortenerClient) Shorten(ctx context.Context, originalURL string) (string, error) {
	requestBody, err := json.Marshal(shortenRequest{URL: originalURL})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/api/v1/shorten", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to shortener service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("shortener service returned non-200 status: %d", resp.StatusCode)
	}

	var responseBody shortenResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to decode shortener service response: %w", err)
	}

	return responseBody.ShortURL, nil
}
