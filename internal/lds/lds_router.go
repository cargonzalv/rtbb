package lds

import (
	"encoding/json"
	"fmt"

	dsplds "github.com/adgear/dsp-lds/pkg/lds"
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/adgear/rtb-bidder/internal/helper"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const (
	KEYSPACE string = "keyspace"
	ID       string = "id"
)

var _ bidder.Router = (*LdsRouter)(nil)
var _ LdsRouterHandler = (*LdsRouter)(nil)

type LdsRouter struct {
	ldsService dsplds.LocalDataStore
}

func ProvideLdsRouter(ldsService dsplds.LocalDataStore) LdsRouter {
	return LdsRouter{ldsService: ldsService}
}

func (l *LdsRouter) AddRoutes(router *fasthttprouter.Router) {
	router.GET("/lds/value", l.GetLdsValueHandler)
	router.GET("/lds/state", l.GetLdsStateHandler)
}

func (l *LdsRouter) GetLdsValueHandler(ctx *fasthttp.RequestCtx) {
	keyspace := ctx.QueryArgs().Peek(KEYSPACE)
	id := ctx.QueryArgs().Peek(ID)

	value, found := l.ldsService.Get(ctx, string(keyspace), string(id))
	if !found {
		ctx.SetStatusCode(fasthttp.StatusNoContent)
		return
	}
	data, err := json.Marshal(value)
	if err != nil {
		log.Error(fmt.Sprintf("Json marshal error: %+v", err))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	helper.WriteJsonBytesResponse(ctx, data)
}

func (l *LdsRouter) GetLdsStateHandler(ctx *fasthttp.RequestCtx) {

	state, err := l.ldsService.GetState(ctx)
	if err != nil {
		log.Error(fmt.Sprintf("Error which getting lds state: %+v", err))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(state)
	if err != nil {
		log.Error(fmt.Sprintf("Json marshal error: %+v", err))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	helper.WriteJsonBytesResponse(ctx, data)
}
