/*
The package which contains observability related code(metrics, logs, tracing).
*/
package observability

import (
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/adgear/rtb-bidder/config"
)

// ProvideLogger loading logger from go_commons library and set logging configuration.
func ProvideLogger(cfg *config.Config) (logger log.Service) {
	logger = log.GetLogger()
	logger.SetLogParams(cfg.Logger.Format, cfg.App.Environment, cfg.Logger.Level, cfg.App.Name, "")
	return
}

// ProvideMetrics loading metrics from go_commons library.
func ProvideMetrics(cfg *config.Config) (metrics metric.Service) {
	metrics = metric.LoadMetricService(cfg.Metrics.Namespace, cfg.Metrics.Subsystem)
	return
}

// ProvideRouter provides `bidder.Router` implementation.
func ProvideRouter() Router {
	return &server{}
}
