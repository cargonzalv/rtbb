package handler

import (
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
)

type router struct {
	webServerName string
	serviceName   string
	samplingRate  float32
	logger        log.Service
	metrics       metric.Service
}

var _ bidder.Router = (*router)(nil)

// ProvideRouter provide `bidder.Router` implementation.
func (r *router) AddRoutes(router *fasthttprouter.Router) {
	infoHandler := infoHandler(r.serviceName)
	router.GET("/info", infoHandler)
	router.NotFound = notFoundHandler(r.logger, r.metrics, r.webServerName, r.serviceName, r.samplingRate)
}
