package services

// TODO: what error(s)

// GatewayServiceIface defines the behavior of the gateway service.
type GatewayServiceIface interface {
	// TODO: what function(s) should go here
}

// Ensure GatewayService explicitly implements GatewayServiceIface.
// This is a compile-time check to ensure the contract is fulfilled.
var _ GatewayServiceIface = (*GatewayService)(nil)

// GatewayService encapsulates the core logic for URL gatewayion.
// This struct now implicitly implements the GatewayServiceIface.
type GatewayService struct {
	// TODO: what will this strcutr have if anything
}

// NewGatewayService creates a new GatewayService instance.
func NewGatewayService() *GatewayService {
	return &GatewayService{}
}
