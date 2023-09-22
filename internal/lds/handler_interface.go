package lds

import (
	"github.com/valyala/fasthttp"
)

type LdsRouterHandler interface {
	GetLdsValueHandler(ctx *fasthttp.RequestCtx)
	GetLdsStateHandler(ctx *fasthttp.RequestCtx)
}
