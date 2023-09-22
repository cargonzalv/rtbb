package helper

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func WriteJsonBytesResponse(ctx *fasthttp.RequestCtx, data []byte) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintf(ctx, "%s", data)
}
