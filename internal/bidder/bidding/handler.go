package bidding

import (
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/rtb-bidder/internal/serializer"
	openrtb2 "github.com/prebid/openrtb/v20/openrtb2"
	"github.com/valyala/fasthttp"
)

type handler struct {
	logger     log.Service
	bidder     Bidder
	serializer serializer.JsonSerializer
}

func (c *handler) GetTvTile(ctx *fasthttp.RequestCtx) {

}

func (c *handler) PostBidJson(ctx *fasthttp.RequestCtx) {
	var openRtbBidRequest openrtb2.BidRequest

	postBody := ctx.PostBody()
	if err := c.serializer.UnmarshalJson(postBody, &openRtbBidRequest); err != nil {
		log.Error(
			"bidprotojson unmarshal failed",
			log.Metadata{"error": err},
			log.Metadata{"payload": postBody},
		)
	}
	resp, err := c.bidder.Bid(&openRtbBidRequest)
	if err != nil {
		c.errorResponse(ctx, err)
		return
	}

	respBody, err := c.serializer.MarshalJson(resp)
	if err != nil {
		c.errorResponse(ctx, err)
		return
	}

	ctx.Response.AppendBody(respBody)
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func (c *handler) errorResponse(ctx *fasthttp.RequestCtx, err error) {
	log.Error(
		"bidding service error",
		log.Metadata{"error": err},
	)
	resp, _ := c.serializer.MarshalJson(err)

	ctx.Response.AppendBody(resp)
	ctx.Response.Header.Set(fasthttp.HeaderContentType, "application/json")
	ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
}
