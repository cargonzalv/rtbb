package health

import (
	"fmt"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/valyala/fasthttp"
)

// Enforce `Handler` interface implementation.
var _ Handler = (*controller)(nil)

type controller struct {
	observer Observer
	logger   log.Service
}

// GetLivenessHandler handles `liveness probe` requests.
func (c *controller) GetLivenessHandler(ctx *fasthttp.RequestCtx) {
	isLive := c.observer.IsLive()
	setResponseStatusCode(ctx, isLive)
	setResponseContext(ctx)
	fmt.Fprintf(ctx, `{"is_live": %t}`, isLive)
}

// GetReadinessHandler handles `readiness probe` requests.
func (c *controller) GetReadinessHandler(ctx *fasthttp.RequestCtx) {
	isReady := c.observer.IsReady()
	setResponseStatusCode(ctx, isReady)
	setResponseContext(ctx)
	fmt.Fprintf(ctx, `{"is_ready": %t}`, isReady)
}

func setResponseStatusCode(ctx *fasthttp.RequestCtx, respValue bool) {
	statusCode := fasthttp.StatusOK
	if !respValue {
		statusCode = fasthttp.StatusServiceUnavailable
	}
	ctx.Response.SetStatusCode(statusCode)
}

func setResponseContext(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Connection", "keep-alive")
	ctx.SetContentType("application/json")
}
