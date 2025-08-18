package web

import (
	"github.com/gin-gonic/gin"
	"github.com/iton0/duss/key-gen-service/internal/api"
)

// NewRouter creates a new Gin router and registers all key generator routes.
func NewRouter(keygenHandler *api.KeygenHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/api/v1/generate-key", keygenHandler.HandleKeygen)

	return router
}
