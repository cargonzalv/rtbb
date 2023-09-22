package health

import (
	"context"
	"sync"
	"time"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
)

// Enforce `Observer` interface implementation.
var _ Observer = (*service)(nil)

type HealthStatus struct {
	IsHealthy  bool
	IsCritical bool
}

// The Health Service struct.
type service struct {
	logger       log.Service
	loopInterval time.Duration
	probes       []probe.Probe
	status       map[string]*HealthStatus
	ctx          context.Context
	cancel       context.CancelFunc
}

// mutex for safe reading scan results.
var mutex = &sync.RWMutex{}

// IsLive is liveness check function.
func (s *service) IsLive() bool {
	return true
}

// IsReady is Rreadiness check function.
func (s *service) IsReady() bool {
	return s.isHealthy()
}

// StartMonitors starts health check monitor. The function will process periodical health probe checks.
func (s *service) StartMonitors(ctx context.Context) {
	if s.ctx != nil {
		s.logger.Warn("ignoring to start a new start monitor call as services have already started")
		return
	}
	s.ctx, s.cancel = context.WithCancel(ctx)

	s.logger.Info("helath observer had been started")
	s.check()
	go func() {
		for ctx.Err() == nil {
			updateTimer := time.NewTicker(s.loopInterval)
			for {
				select {
				case <-ctx.Done():
					s.logger.Info("helath observer get cancel signal")
					return
				case <-updateTimer.C:
					s.check()
				}
			}
		}
	}()
}

// Stop health check monitor.
func (s *service) StopMonitors() {
	s.logger.Info("health observer had been stoped")
	if s.ctx != nil {
		s.ctx = nil
		s.cancel()
	}
}

// Process health probes checks.
func (s *service) check() {
	startTime := time.Now()
	mutex.Lock()
	for _, p := range s.probes {
		s.status[p.Name()].IsHealthy = p.Check()
	}
	mutex.Unlock()
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	s.logger.Debug("health service monitor check completed",
		log.Metadata{"start_time": startTime},
		log.Metadata{"end_time": endTime},
		log.Metadata{"duration": duration},
	)
}

// The monitor status consider helthy if all critical probes are helthy.
func (s *service) isHealthy() bool {
	mutex.RLock()
	for _, s := range s.status {
		if s.IsCritical && !s.IsHealthy {
			mutex.RUnlock()
			return false
		}
	}
	mutex.RUnlock()
	return true
}

// Setup initial state of health status.
func initStatus(probes []probe.Probe) (status map[string]*HealthStatus) {
	status = make(map[string]*HealthStatus, len(probes))
	for _, p := range probes {
		status[p.Name()] = &HealthStatus{IsHealthy: false, IsCritical: p.IsCritical()}
	}
	return
}
