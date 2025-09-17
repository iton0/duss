package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/api-gateway-service/internal/core/services"
)

// ShortenRequest represents the request body for shortening a URL.
type ShortenRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// GatewayHandler holds the necessary dependencies for the handler.
type GatewayHandler struct {
	gatewayService services.GatewayServiceIface
}

// NewGatewayHandler creates a new GatewayHandler instance.
func NewGatewayHandler(gs services.GatewayServiceIface) *GatewayHandler {
	return &GatewayHandler{gatewayService: gs}
}

// HandleShorten handles the POST /shorten request.
func (h *GatewayHandler) HandleShorten(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body or URL format"})
		return
	}

	shortURL, err := h.gatewayService.ShortenURL(c.Request.Context(), req.URL)
	if err != nil {
		// Log the error internally and return a generic server error to the client.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// HandleRedirect handles the GET /:shortKey request.
func (h *GatewayHandler) HandleRedirect(c *gin.Context) {
	shortKey := c.Param("shortKey")
	if shortKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short key is required"})
		return
	}

	originalURL, err := h.gatewayService.RedirectURL(c.Request.Context(), shortKey)
	if err != nil {
		// Differentiate between a "not found" error and a generic server error.
		// For example, if the error is due to a non-existent key.
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Perform the HTTP redirect to the original URL.
	c.Redirect(http.StatusMovedPermanently, originalURL)
}
