package observability

import (
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

var _ bidder.Router = (*server)(nil)

type server struct{}

// ProvideRouter provide `bidder.Router` implementation.
func (*server) AddRoutes(router *fasthttprouter.Router) {
	// Prometheus metrics handler
	router.GET("/metrics", fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler()))
}
