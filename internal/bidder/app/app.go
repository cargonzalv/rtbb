/*
Package app package contains the application bootstrap logic.
The dependency loading and wiring generated with package 'wire'.
*/
package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	dsplds "github.com/adgear/dsp-lds/pkg/lds"
	"github.com/adgear/go-commons/pkg/buildinfo"
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/rtb-bidder/internal/bidder/health"
	"github.com/adgear/rtb-bidder/internal/bidder/server"
	"github.com/oklog/run"
)

type App struct {
	logger        log.Service
	metrics       metric.Service
	restServer    server.Server
	healthService health.Observer
	ldsService    dsplds.LocalDataStore
}

// Start the application.
func (a *App) Start() {
	logMetadata := log.Metadata{
		"git_commit": buildinfo.GitCommit,
		"git_tag":    buildinfo.GitDescribeTag,
		"build_time": buildinfo.BuildTime,
	}
	a.logger.Info("building application", logMetadata)
	// Start health monitoring.
	// Health monitoring service is running in separate context.
	ctx := context.Background()

	var g run.Group
	{
		g.Add(func() error {
			a.healthService.StartMonitors(ctx)
			err := a.ldsService.Start(ctx)
			if err != nil {
				a.logger.Error("Error while starting lds, error message", log.Metadata{"error": err})
				return err
			}
			return a.restServer.Start()
		}, func(err error) {
			a.logger.Info("Interrupt rest server run group trigged with error: ", log.Metadata{"error": err})
			a.healthService.StopMonitors()
			serverErr := a.restServer.Shutdown()
			if serverErr == http.ErrServerClosed {
				a.logger.Info("Successfully close server")
			} else {
				a.logger.Info("Unable to close server successfully")
			}
			a.ldsService.Stop()
			a.logger.Info("Done with interrupt function rest server")
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		// On service signal received, will process the
		// service gracefull shutdown.
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			sig := <-c
			log.Info("service signal received", log.Any("signal", sig))
			return nil
		}, func(err error) {
			if err != nil {
				log.Error("server shutdown due to error",
					log.NamedError("error", err),
				)
			}
		})
	}

	a.logger.Info("starting application", logMetadata)

	err := g.Run()
	if err != nil {
		logMetadata["error"] = err
		a.logger.Error("failed to run application", logMetadata)
	}
}
