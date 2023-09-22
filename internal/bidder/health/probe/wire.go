package probe

import "github.com/adgear/go-commons/pkg/metric"

// Dependency injection Health check probes provider.
func ProvideProbes(metrics metric.Service) []Probe {
	metricProbe := &metricProbe{metrics}
	return []Probe{
		metricProbe,
	}
}
