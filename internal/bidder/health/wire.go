package health

import (
	"time"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/rtb-bidder/config"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
)

// ProvideService provides health `Observer` implementation.
func ProvideService(logger log.Service, probes []probe.Probe, cfg *config.Config) Observer {
	loopInterval := time.Duration(cfg.HealthCheck.LoopIntervalSeconds) * time.Second
	return &service{
		logger:       logger,
		loopInterval: loopInterval,
		probes:       probes,
		status:       initStatus(probes),
	}
}

// ProvideHandler provides `Handler` implementation.
func ProvideHandler(logger log.Service, service Observer) Handler {
	return &controller{service, logger}
}

// ProvideRouter provide `bidder.Router` implementation.
func ProvideRouter(handler Handler) Router {
	return &router{
		handler: handler,
	}
}
