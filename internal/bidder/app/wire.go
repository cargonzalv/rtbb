//go:build wireinject
// +build wireinject

package app

import (
	"github.com/adgear/rtb-bidder/config"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/adgear/rtb-bidder/internal/bidder/bidding"
	"github.com/adgear/rtb-bidder/internal/bidder/health"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
	"github.com/adgear/rtb-bidder/internal/bidder/server"
	"github.com/adgear/rtb-bidder/internal/bidder/server/handler"
	"github.com/adgear/rtb-bidder/internal/lds"
	"github.com/adgear/rtb-bidder/internal/observability"
	"github.com/adgear/rtb-bidder/internal/serializer"
	"github.com/google/wire"
)

// ProvideRouters is the function providing the list of routers from all packages, which have `bidder.Router` interface implemented.
func ProvideRouters(health health.Router, bidding bidding.Router, handler handler.Router, observability observability.Router, ldsRouter lds.LdsRouter) *[]bidder.Router {
	return &[]bidder.Router{
		health,
		bidding,
		handler,
		observability,
		&ldsRouter,
	}
}

// Using by the wire to generate gependency tree.
// This file is not include into release.
func Wire() *App {
	panic(wire.Build(
		// observability package providers
		observability.ProvideLogger,
		observability.ProvideMetrics,
		observability.ProvideRouter,

		// application Config providers
		config.ProvideConfig,

		// server package providers
		server.ProvideRestServer,
		server.ProvideFastHttpRouter,

		// bidding package providers
		bidding.ProvideBidder,
		bidding.ProvideHandler,
		bidding.ProvideRouter,

		// health package providers
		health.ProvideService,
		health.ProvideHandler,
		health.ProvideRouter,

		handler.ProvideRouter,

		// probe package providers
		probe.ProvideProbes,

		// serializer package providers
		serializer.ProvideJsonSerializer,

		//lds package providers
		lds.ProvideLocalDataStore,
		lds.ProvideLdsRouter,

		ProvideRouters,
		wire.Struct(new(App), "*"),
	))
}
