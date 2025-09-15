package api

import (
	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/api-gateway-service/internal/core/services"
)

// GatewayHandler holds the necessary dependencies for the handler.
type GatewayHandler struct {
	// Now depends on the GatewayServiceIface interface
	gatewayService services.GatewayServiceIface
}

// NewGatewayHandler creates a new GatewayHandler instance.
// The constructor now accepts the new interface type.
func NewGatewayHandler(rs services.GatewayServiceIface) *GatewayHandler {
	return &GatewayHandler{gatewayService: rs}
}

// HandleGateway handles the GET /:shortKey request using Gin's context.
func (h *GatewayHandler) HandleGateway(c *gin.Context) {
	// TODO: implement handle gateway logic
}
