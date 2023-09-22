package bidding

import (
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
)

// Enforce `Router` interface implementation.
var _ bidder.Router = (*router)(nil)

type router struct {
	// Health monitoring controller.
	handler Handler
}

// AddRoutes add `bidding package` routes.
func (r *router) AddRoutes(router *fasthttprouter.Router) {
	router.GET("/impressions/tile", r.handler.GetTvTile)
	router.POST("/bidj", r.handler.PostBidJson)
}
