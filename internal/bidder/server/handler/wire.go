package handler

import (
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/rtb-bidder/config"
)

// ProvideRouter provides `bidder.Router` implementation.
func ProvideRouter(logger log.Service, metrics metric.Service, cfg *config.Config) Router {
	return &router{
		webServerName: cfg.Http.Name,
		serviceName:   cfg.App.Name,
		samplingRate:  cfg.Http.NotFoundSamplingRate,
		logger:        logger,
		metrics:       metrics,
	}
}
