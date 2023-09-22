package handler

import (
	"math/rand"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/valyala/fasthttp"
)

// notFoundHandler return handler for requests which does not match any registered routes.
func notFoundHandler(logger log.Service, metrics metric.Service, webServerName string, serviceName string, samplingRate float32) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		//publish counter
		labels := map[string]string{
			"name":       webServerName,
			"service":    serviceName,
			"error_type": "route_not_found",
		}

		metrics.Count("route_not_found", 1, labels)

		//log samples of Not found request
		if rand.Float32() > float32(samplingRate) {
			h := &ctx.Request.Header

			logger.Warn("route not found",
				log.Metadata{"req": string(ctx.Request.URI().RequestURI())},
				log.Metadata{"error_type": "route_not_found"},
				log.Metadata{"origin": string(h.Peek(fasthttp.HeaderOrigin))},
				log.Metadata{"referrer": string(h.Referer())},
				log.Metadata{"ua": string(h.UserAgent())},
				log.Metadata{"dnt": string(h.Peek(fasthttp.HeaderDNT))},
				log.Metadata{"remote_ip": ctx.RemoteIP().String()},
				log.Metadata{"x_fwd_for": string(h.Peek(fasthttp.HeaderXForwardedFor))},
				log.Metadata{"x_fwd_proto": string(h.Peek(fasthttp.HeaderXForwardedProto))},
			)
		}
		ctx.Response.SetStatusCode(fasthttp.StatusNotFound)
	}
}
