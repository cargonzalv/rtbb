package probe

import (
	"time"

	"github.com/adgear/go-commons/pkg/metric"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metrics Probe", func() {
	var (
		mockCtrl    *gomock.Controller
		mockMetrics *metric.MockService
	)

	Context("should have `/health/liveness` and `/health/readiness`", func() {

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockMetrics = metric.NewMockService(mockCtrl)
		})

		It("call to probe function Name() returns `metrics`", func(ctx SpecContext) {
			subject := &metricProbe{
				metrics: mockMetrics,
			}
			name := subject.Name()
			Expect(name).To(Equal("metrics"))

		}, SpecTimeout(time.Second*10))

		It("call to probe function IsCritical() returns `false`", func(ctx SpecContext) {
			subject := &metricProbe{
				metrics: mockMetrics,
			}
			isCritical := subject.IsCritical()
			Expect(isCritical).To(Equal(false))

		}, SpecTimeout(time.Second*10))

		It("when metric.IsReady returns false, call to probe function Check() returns false", func(ctx SpecContext) {
			mockMetrics.EXPECT().IsReady().Return(false)
			subject := &metricProbe{
				metrics: mockMetrics,
			}
			checkResult := subject.Check()
			Expect(checkResult).To(Equal(false))
		}, SpecTimeout(time.Second*10))

		It("when metric.IsReady returns true, call to probe function Check() returns true", func(ctx SpecContext) {
			mockMetrics.EXPECT().IsReady().Return(true)
			subject := &metricProbe{
				metrics: mockMetrics,
			}
			checkResult := subject.Check()
			Expect(checkResult).To(Equal(true))
		}, SpecTimeout(time.Second*10))
	})
})
