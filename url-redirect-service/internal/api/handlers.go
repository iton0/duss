package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/url-redirect-service/internal/core/services"
)

// RedirectHandler holds the necessary dependencies for the handler.
type RedirectHandler struct {
	// Now depends on the RedirectServiceIface interface
	redirectService services.RedirectServiceIface
}

// NewRedirectHandler creates a new RedirectHandler instance.
// The constructor now accepts the new interface type.
func NewRedirectHandler(rs services.RedirectServiceIface) *RedirectHandler {
	return &RedirectHandler{redirectService: rs}
}

// HandleRedirect handles the GET /:shortKey request using Gin's context.
func (h *RedirectHandler) HandleRedirect(c *gin.Context) {
	shortKey := c.Param("shortKey")

	if shortKey == "" {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	longURL, err := h.redirectService.GetOriginalURL(c.Request.Context(), shortKey)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrURLNotFound):
			c.String(http.StatusNotFound, "Not Found")
			return
		default:
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
