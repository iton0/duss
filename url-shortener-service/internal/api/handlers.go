package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/url-shortener-service/internal/core/services"
)

type ShortenerHandler struct {
	shortenerService services.ShortenerServiceIface
}

// NewRedirectHandler creates a new RedirectHandler instance.
// The constructor now accepts the new interface type.
func NewShortenerHandler(ss services.ShortenerServiceIface) *ShortenerHandler {
	return &ShortenerHandler{shortenerService: ss}
}

// HandleRedirect handles the GET /:shortKey request using Gin's context.
func (h *ShortenerHandler) HandleShortener(c *gin.Context) {
	// TODO: implement the shorterner version of this function
	shortKey := c.Param("shortKey")

	if shortKey == "" {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	longURL, err := h.shortenerService.Shorten()(c.Request.Context(), shortKey)
	if err != nil {
		if errors.Is(err, services.ErrURLNotFound) {
			c.String(http.StatusNotFound, "Not Found")
			return
		}
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
