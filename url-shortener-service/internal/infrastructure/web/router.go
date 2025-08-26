package web

import (
	"github.com/gin-gonic/gin"
	"github.com/iton0/duss/url-shortener-service/internal/api"
)

// NewRouter creates a new Gin router and registers all routes.
func NewRouter(shortenerHandler *api.ShortenerHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/api/v1/shorten", shortenerHandler.HandleShortener)

	return router
}
