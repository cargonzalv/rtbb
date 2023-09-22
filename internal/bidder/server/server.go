/*
The server package has logic related to the Rest Http server.
*/
package server

import (
	"fmt"
	"net"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/rtb-bidder/internal/bidder"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

// The web server struct.
type server struct {
	// Web server port.
	port int

	fastHttpRouter *fasthttprouter.Router

	routers *[]bidder.Router

	// Web server pointer.
	httpserver *fasthttp.Server

	// Logger instance.
	logger log.Service

	// Metrics instance.
	metrics metric.Service
}

// Enforce `Service` interface implementation.
var _ Server = (*server)(nil)

// Start the new RestServer.
func (s *server) Start() error {
	// add routes from all packages
	for _, r := range *s.routers {
		r.AddRoutes(s.fastHttpRouter)
	}

	addr := "0.0.0.0:" + fmt.Sprint(s.port)
	ln, err := net.Listen("tcp4", addr)

	if err != nil {
		s.logger.Error("Error creating listener", log.Metadata{"error": err, "addr": addr})
		return err
	}

	s.logger.Info("Web server starting", log.Metadata{
		"addr":                         addr,
		"http_read_timeout":            s.httpserver.ReadTimeout,
		"http_write_timeout":           s.httpserver.WriteTimeout,
		"http_max_keep_alive_duration": s.httpserver.IdleTimeout,
		"http_max_conns_per_ip":        s.httpserver.MaxConnsPerIP,
		"http_max_requests_per_conn":   s.httpserver.MaxRequestsPerConn,
	})
	return s.httpserver.Serve(ln)
}

func (s *server) Shutdown() error {
	return s.httpserver.Shutdown()
}
