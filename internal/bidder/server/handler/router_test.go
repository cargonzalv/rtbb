package handler

import (
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/buaazp/fasthttprouter"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info Routes", func() {
	var (
		mockMetricService *metric.MockService
		mockLogger        *log.MockService
		fastHttpRouter    *fasthttprouter.Router
	)

	BeforeEach(func() {
		fastHttpRouter = fasthttprouter.New()
		r := router{
			webServerName: "test-server",
			serviceName:   "rtb-bidder",
			samplingRate:  0,
			logger:        mockLogger,
			metrics:       mockMetricService,
		}
		r.AddRoutes(fastHttpRouter)
	})

	Context("info route is present", func() {
		It("contains handler for GET /info", func() {
			hndlr, _ := fastHttpRouter.Lookup("GET", "/info", nil)
			Expect(hndlr).ShouldNot(BeNil())
		})

		It("contains NotFound handler", func() {
			Expect(fastHttpRouter.NotFound).ShouldNot(BeNil())
		})
	})
})
