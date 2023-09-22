package health

import (
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
)

// Enforce `bidder.Router` interface implementation.
var _ bidder.Router = (*router)(nil)

type router struct {
	// Health monitoring controller.
	handler Handler
}

// AddRoutes add `bidding package` routes.
func (r *router) AddRoutes(router *fasthttprouter.Router) {
	// k8s health probes handlers. This probes used to control health of
	// the application.
	// * Failing liveness probe will restart the container.
	// * Failing readiness probe will stop the application from serving traffic.
	router.GET("/health/liveness", r.handler.GetLivenessHandler)
	router.GET("/health/readiness", r.handler.GetReadinessHandler)
}
