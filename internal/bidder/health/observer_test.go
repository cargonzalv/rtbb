package health

import (
	"sync"
	"time"

	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/rtb-bidder/internal/bidder/health/probe"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Observer", func() {
	var (
		mockCtrl   *gomock.Controller
		mockLogger *log.MockService
		mockProbe1 *probe.MockProbe
		mockProbe2 *probe.MockProbe
		wg         sync.WaitGroup
	)

	Context("should start and stop monitors", func() {

		BeforeEach(func() {
			cntr1, cntr2 := 2, 2
			mockCtrl = gomock.NewController(GinkgoT())
			mockLogger = log.NewMockService(mockCtrl)
			mockProbe1 = probe.NewMockProbe(mockCtrl)

			mockProbe1.EXPECT().Name().Return("probe1").AnyTimes()
			mockProbe1.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr1--
				if cntr1 == 0 {
					wg.Done()
				}
				return true
			}).MinTimes(1)
			mockProbe1.EXPECT().IsCritical().Return(false).MinTimes(1)

			mockProbe2 = probe.NewMockProbe(mockCtrl)

			mockProbe2.EXPECT().Name().Return("probe2").AnyTimes()
			mockProbe2.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr2--
				if cntr2 == 0 {
					wg.Done()
				}
				return true
			}).MinTimes(1)

			mockProbe2.EXPECT().IsCritical().Return(false).MinTimes(1)

		})

		It("should start monitoring and process check probes", func(ctx SpecContext) {
			wg.Add(2)

			mockLogger.EXPECT().Debug(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

			// start monitoring should log info message.
			mockLogger.EXPECT().Info(gomock.Any()).Times(1)
			mockLogger.EXPECT().Warn(gomock.Any()).Times(0)

			probes := []probe.Probe{mockProbe1, mockProbe2}

			subject := &service{
				logger:       mockLogger,
				loopInterval: 1,
				probes:       probes,
				status:       initStatus(probes),
			}

			subject.StartMonitors(ctx)

			// wait until probe checks will be executed 2 times
			wg.Wait()

			mockLogger.EXPECT().Info(gomock.Any()).AnyTimes()
			subject.StopMonitors()

			Expect(subject.IsLive()).To(Equal(true))
			Expect(subject.IsReady()).To(Equal(true))

		}, SpecTimeout(time.Second*10))

		It("should log warning when start running monitoring process", func(ctx SpecContext) {
			wg.Add(2)

			mockLogger.EXPECT().Debug(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

			mockLogger.EXPECT().Info(gomock.Any()).Times(1)

			probes := []probe.Probe{mockProbe1, mockProbe2}

			subject := &service{
				logger:       mockLogger,
				loopInterval: 1,
				probes:       probes,
				status:       initStatus(probes),
			}
			subject.StartMonitors(ctx)

			mockLogger.EXPECT().Warn(gomock.Any()).Times(1)

			subject.StartMonitors(ctx)

			wg.Wait()

			mockLogger.EXPECT().Info(gomock.Any()).AnyTimes()
			subject.StopMonitors()

			Expect(subject.IsLive()).To(Equal(true))
			Expect(subject.IsReady()).To(Equal(true))

		}, SpecTimeout(time.Second*10))
	})

	Context("should report IsReady", func() {

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockLogger = log.NewMockService(mockCtrl)
			mockProbe1 = probe.NewMockProbe(mockCtrl)
			mockProbe2 = probe.NewMockProbe(mockCtrl)

			mockLogger.EXPECT().Debug(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
			mockLogger.EXPECT().Info(gomock.Any()).AnyTimes()
			mockLogger.EXPECT().Warn(gomock.Any()).AnyTimes()
		})

		It("when all probes Check() return true, IsReady returns true", func(ctx SpecContext) {
			cntr := 2
			cntr1, cntr2 := cntr, cntr
			wg.Add(cntr)

			mockProbe1.EXPECT().Name().Return("probe1").AnyTimes()
			// "probe1" is healthy
			mockProbe1.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr1--
				if cntr1 == 0 {
					wg.Done()
				}
				return true
			}).MinTimes(1)
			// "probe1" is critical
			mockProbe1.EXPECT().IsCritical().Return(true).MinTimes(1)

			mockProbe2.EXPECT().Name().Return("probe2").AnyTimes()
			// "probe1" is healthy
			mockProbe2.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr2--
				if cntr2 == 0 {
					wg.Done()
				}
				return true
			}).MinTimes(1)
			mockProbe2.EXPECT().IsCritical().Return(false).MinTimes(1)

			probes := []probe.Probe{mockProbe1, mockProbe2}

			subject := &service{
				logger:       mockLogger,
				loopInterval: 1,
				probes:       probes,
				status:       initStatus(probes),
			}

			subject.StartMonitors(ctx)

			// wait until probe checks will be executed 2 times
			wg.Wait()

			subject.StopMonitors()

			Expect(subject.IsLive()).To(Equal(true))
			// IsReady retruns true
			Expect(subject.IsReady()).To(Equal(true))

		}, SpecTimeout(time.Second*10))

		It("when any probe with IsCritical = true and Check() returns false, IsReady returns false", func(ctx SpecContext) {
			cntr := 2
			cntr1, cntr2 := cntr, cntr
			wg.Add(cntr)

			mockProbe1.EXPECT().Name().Return("probe1").AnyTimes()
			// "probe1" is failing
			mockProbe1.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr1--
				if cntr1 == 0 {
					wg.Done()
				}
				return false
			}).MinTimes(1)
			// "probe1" is critical
			mockProbe1.EXPECT().IsCritical().Return(true).MinTimes(1)

			mockProbe2.EXPECT().Name().Return("probe2").AnyTimes()
			// "probe1" is healthy
			mockProbe2.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr2--
				if cntr2 == 0 {
					wg.Done()
				}
				return true
			}).MinTimes(1)
			mockProbe2.EXPECT().IsCritical().Return(false).MinTimes(1)

			probes := []probe.Probe{mockProbe1, mockProbe2}

			subject := &service{
				logger:       mockLogger,
				loopInterval: 1,
				probes:       probes,
				status:       initStatus(probes),
			}

			subject.StartMonitors(ctx)

			// wait until probe checks will be executed 2 times
			wg.Wait()

			subject.StopMonitors()

			Expect(subject.IsLive()).To(Equal(true))
			// IsReady retruns false
			Expect(subject.IsReady()).To(Equal(false))

		}, SpecTimeout(time.Second*10))

		It("when any probe with IsCritical = false and Check() returns false, IsReady returns true", func(ctx SpecContext) {
			cntr := 2
			cntr1, cntr2 := cntr, cntr
			wg.Add(cntr)

			mockProbe1.EXPECT().Name().Return("probe1").AnyTimes()
			// "probe1" is failing
			mockProbe1.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr1--
				if cntr1 == 0 {
					wg.Done()
				}
				return false
			}).MinTimes(1)
			// "probe1" is non critical
			mockProbe1.EXPECT().IsCritical().Return(false).MinTimes(1)

			mockProbe2.EXPECT().Name().Return("probe2").AnyTimes()
			// "probe1" is healthy
			mockProbe2.EXPECT().Check().DoAndReturn(func() bool {
				// Every time monitor runs `check()` substruct 1 from waith group counter
				cntr2--
				if cntr2 == 0 {
					wg.Done()
				}
				return true
			}).MinTimes(1)
			mockProbe2.EXPECT().IsCritical().Return(false).MinTimes(1)

			probes := []probe.Probe{mockProbe1, mockProbe2}

			subject := &service{
				logger:       mockLogger,
				loopInterval: 1,
				probes:       probes,
				status:       initStatus(probes),
			}

			subject.StartMonitors(ctx)

			// wait until probe checks will be executed 2 times
			wg.Wait()

			subject.StopMonitors()

			Expect(subject.IsLive()).To(Equal(true))
			// IsReady retruns true
			Expect(subject.IsReady()).To(Equal(true))

		}, SpecTimeout(time.Second*10))

	})
})
