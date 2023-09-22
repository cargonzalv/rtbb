package bidding

import (
	"github.com/adgear/rtb-bidder/internal/bidder"
	openrtb2 "github.com/prebid/openrtb/v20/openrtb2"
	"github.com/valyala/fasthttp"
)

//go:generate mockgen  -destination=mocks.gen.go -source=$GOFILE -package=$GOPACKAGE

// Service Interface to define the bidding service behaviour.
type Bidder interface {
	Bid(req *openrtb2.BidRequest) (resp *openrtb2.BidResponse, err error)
}

// Handler Interface for the bidding service web requests handler.
type Handler interface {
	GetTvTile(ctx *fasthttp.RequestCtx)
	PostBidJson(ctx *fasthttp.RequestCtx)
}

// Router is proxy type of bidder.Router required by wire resolver.
type Router bidder.Router
