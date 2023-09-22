package bidding

import (
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/rtb-bidder/internal/serializer"
)

// ProvideBidder provides `Bidder` implementation.
func ProvideBidder(logger log.Service, metrics metric.Service) Bidder {
	return &service{logger, metrics}
}

// ProvideHandler provides `Handler` implementation.
func ProvideHandler(logger log.Service, bidder Bidder, serializer serializer.JsonSerializer) Handler {
	return &handler{
		logger,
		bidder,
		serializer,
	}
}

// ProvideRouter provides the `bidder.Router` implementation.
func ProvideRouter(handler Handler) Router {
	return &router{
		handler: handler,
	}
}
