package server

import (
	"crypto/tls"
	"time"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/go-commons/pkg/utils/httputils"
	"github.com/adgear/rtb-bidder/config"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// ProvideFastHttpRouter provides fasthttp Router.
func ProvideFastHttpRouter() *fasthttprouter.Router {
	return fasthttprouter.New()
}

// ProvideRestServer provides Rest Server implementation.
func ProvideRestServer(logger log.Service, metrics metric.Service, fasthttprouter *fasthttprouter.Router, routers *[]bidder.Router, cfg *config.Config) Server {
	params := RestServerParams{
		Port:                cfg.Http.Port,
		MaxConnsPerIp:       cfg.Http.MaxConnsPerIp,
		MaxReqPerConn:       cfg.Http.MaxReqPerConn,
		ReadTimeoutSeconds:  cfg.Http.ReadTimeoutSeconds,
		WriteTimeoutSeconds: cfg.Http.WriteTimeoutSeconds,
		IdleTimeoutSeconds:  cfg.IdleTimeoutSeconds,
		Name:                cfg.Http.Name,
	}

	httpserver := initFasthttpServer(metrics, fasthttprouter, params)

	return &server{
		httpserver:     httpserver,
		fastHttpRouter: fasthttprouter,
		routers:        routers,
		port:           params.Port,
		logger:         logger,
		metrics:        metrics,
	}
}

// initialize, fasthttp server and set metrics middleware.
func initFasthttpServer(metrics metric.Service, router *fasthttprouter.Router, params RestServerParams) (httpserver *fasthttp.Server) {
	handler := httputils.MetricsHandler(metrics, router)
	handler = fasthttp.CompressHandler(handler)
	httpserver = &fasthttp.Server{
		Handler:                            handler,
		Name:                               params.Name,
		Concurrency:                        0,
		ReadBufferSize:                     0,
		WriteBufferSize:                    0,
		ReadTimeout:                        timeDuration(params.ReadTimeoutSeconds),
		WriteTimeout:                       timeDuration(params.WriteTimeoutSeconds),
		IdleTimeout:                        timeDuration(params.IdleTimeoutSeconds),
		MaxConnsPerIP:                      params.MaxConnsPerIp,
		MaxRequestsPerConn:                 params.MaxReqPerConn,
		MaxIdleWorkerDuration:              0,
		TCPKeepalivePeriod:                 0,
		MaxRequestBodySize:                 0,
		DisableKeepalive:                   false,
		TCPKeepalive:                       false,
		ReduceMemoryUsage:                  false,
		GetOnly:                            false,
		DisablePreParseMultipartForm:       false,
		LogAllErrors:                       false,
		SecureErrorLogMessage:              false,
		DisableHeaderNamesNormalizing:      false,
		SleepWhenConcurrencyLimitsExceeded: 0,
		NoDefaultServerHeader:              false,
		NoDefaultDate:                      false,
		NoDefaultContentType:               false,
		KeepHijackedConns:                  false,
		CloseOnShutdown:                    false,
		StreamRequestBody:                  false,
		Logger:                             nil,
		TLSConfig:                          &tls.Config{},
	}
	return
}

// Cast value to time duration in seconds.
func timeDuration(value int) time.Duration {
	return time.Duration(value) * time.Second
}
