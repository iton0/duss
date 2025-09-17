package web

import (
	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/api-gateway-service/internal/api"
)

// NewRouter creates a new Gin router and registers all gateway routes.
func NewRouter(gatewayHandler *api.GatewayHandler) *gin.Engine {
	router := gin.Default()

	// Public API endpoints
	router.POST("/shorten", gatewayHandler.HandleShorten)
	router.GET("/:shortKey", gatewayHandler.HandleRedirect)

	return router
}
