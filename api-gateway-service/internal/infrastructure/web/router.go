package web

import (
	"github.com/gin-gonic/gin"
	"github.com/iton0/duss/api-gateway-service/internal/api"
)

// NewRouter creates a new Gin router and registers all key generator routes.
func NewRouter(gatewayHandler *api.GatewayHandler) *gin.Engine {
	router := gin.Default()

	// TODO: what are the endpoitns that i register
	// properly just the other services but pulbic facing

	return router
}
