package probe

import "github.com/adgear/go-commons/pkg/metric"

// Enforce Probe interface implementation.
var _ Probe = (*metricProbe)(nil)

// The metric probe struct.
type metricProbe struct {
	metrics metric.Service
}

// The function returns metric probe name.
func (*metricProbe) Name() string {
	return "metrics"
}

// The function process metric service health check.
func (p *metricProbe) Check() bool {
	return p.metrics.IsReady()
}

func (p *metricProbe) IsCritical() bool {
	return false
}
