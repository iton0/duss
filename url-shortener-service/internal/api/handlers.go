package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/url-shortener-service/internal/core/services"
)

// RequestBody defines the structure for the JSON request body.
type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// ResponseBody defines the structure for the JSON response body.
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

type ShortenerHandler struct {
	shortenerService services.ShortenerServiceIface
}

// NewShortenerHandler creates a new ShortenerHandler instance.
func NewShortenerHandler(ss services.ShortenerServiceIface) *ShortenerHandler {
	return &ShortenerHandler{shortenerService: ss}
}

// HandleShortener handles the POST /api/v1/shorten request.
func (h *ShortenerHandler) HandleShortener(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: URL is required and must be a valid format"})
		return
	}

	shortenedURL, err := h.shortenerService.Shorten(c.Request.Context(), req.URL)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrInvalidURL):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		case errors.Is(err, services.ErrBlacklistedURL):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		case errors.Is(err, services.ErrDuplicatedKey):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	// Assuming a base URL for the shortened link, e.g., "http://localhost:8080/"
	// The base URL would be a configuration value in a real application.
	// For this example, let's assume it's hardcoded.
	baseURL := "http://localhost:8081/"
	fullShortURL := baseURL + shortenedURL.ShortKey

	c.JSON(http.StatusCreated, ShortenResponse{ShortURL: fullShortURL})
}
