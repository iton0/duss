package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iton0/duss/key-gen-service/internal/core/services"
)

// KeygenHandler handles API requests related to key generation.
type KeygenHandler struct {
	keygenService services.KeygenServiceIface
}

// NewKeygenHandler creates a new instance of KeygenHandler.
func NewKeygenHandler(ks services.KeygenServiceIface) *KeygenHandler {
	return &KeygenHandler{keygenService: ks}
}

// HandleKeygen handles the GET /api/v1/generate-key endpoint.
// It calls the key generation service to create a new key and returns it.
func (h *KeygenHandler) HandleKeygen(c *gin.Context) {
	// Call the keygen service to generate a new short key.
	// The keygenService interface should have a method like GenerateKey()
	// that handles the actual key generation logic.
	shortKey, err := h.keygenService.GenerateKey()
	if err != nil {
		// If an error occurs during key generation, return an internal server error.
		// A common reason for this could be a failure to connect to a database or a random number generator.
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// If successful, return the newly generated short key to the client.
	// We use c.JSON to return a JSON response with the key.
	c.JSON(http.StatusOK, gin.H{"short_key": shortKey})
}
