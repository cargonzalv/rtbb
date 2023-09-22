package health

import (
	"context"

	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/valyala/fasthttp"
)

//go:generate mockgen  -destination=mocks.gen.go -source=interface.go -package=health

// Service Interface is to define the health service behaviour.
type Observer interface {
	IsLive() bool
	IsReady() bool
	StartMonitors(context.Context)
	StopMonitors()
}

// Handler Interface is for the health web requests handler.
type Handler interface {
	GetLivenessHandler(ctx *fasthttp.RequestCtx)
	GetReadinessHandler(ctx *fasthttp.RequestCtx)
}

// Router is proxy type of bidder.Router required by wire resolver.
type Router bidder.Router
