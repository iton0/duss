package web

import (
	"github.com/gin-gonic/gin"
	"github.com/iton0/duss/url-redirect-service/internal/api"
)

// NewRouter creates a new Gin router and registers all routes.
func NewRouter(redirectHandler *api.RedirectHandler) *gin.Engine {
	// gin.Default() provides middleware for logging and recovery from panics
	router := gin.Default()

	// Register the GET /:shortKey endpoint to the appropriate handler
	router.GET("/:shortKey", redirectHandler.HandleRedirect)

	return router
}
