package api

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iton0/duss/shared/domain" // Import the shared domain model
	"github.com/iton0/duss/url-shortener-service/internal/core/services"
)

// CreateURLRequest is the API request body.
type CreateURLRequest struct {
	LongURL string `json:"long_url"`
}

// ShortenHandler encapsulates the API logic.
type ShortenHandler struct {
	shortenerService *services.ShortenerService
}

func (h *ShortenHandler) CreateShortenedURL(w http.ResponseWriter, r *http.Request) {
	var req CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the core service to handle the business logic, passing the data as the LongURL field
	shortenedURL, err := h.shortenerService.Shorten(r.Context(), req.LongURL)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	// Respond to the client with the created domain.URL struct
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(shortenedURL)
}
