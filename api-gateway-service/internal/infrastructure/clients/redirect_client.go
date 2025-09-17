package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/iton0/duss/api-gateway-service/internal/core/services"
)

// HTTPRedirectClient is a concrete implementation of the RedirectServiceClient interface.
type HTTPRedirectClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPRedirectClient creates a new HTTP client for the redirect service.
func NewHTTPRedirectClient(baseURL string) services.RedirectServiceClient {
	return &HTTPRedirectClient{
		client:  &http.Client{Timeout: 5 * time.Second},
		baseURL: baseURL,
	}
}

// redirectResponse mirrors the expected JSON structure of the redirect service.
type redirectResponse struct {
	OriginalURL string `json:"original_url"`
}

// GetOriginalURL sends an HTTP GET request to the redirect service.
func (c *HTTPRedirectClient) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/api/v1/redirect?key=%s", c.baseURL, shortURL), nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request to redirect service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("redirect service returned non-200 status: %d", resp.StatusCode)
	}

	var responseBody redirectResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to decode redirect service response: %w", err)
	}

	return responseBody.OriginalURL, nil
}
