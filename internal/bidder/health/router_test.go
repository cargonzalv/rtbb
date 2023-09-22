package health

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Routes", func() {
	var (
		mockCtrl       *gomock.Controller
		mockHandler    *MockHandler
		fastHttpRouter *fasthttprouter.Router
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockHandler = NewMockHandler(mockCtrl)
		fastHttpRouter = fasthttprouter.New()
		r := router{handler: mockHandler}
		r.AddRoutes(fastHttpRouter)
	})

	Context("Readiness and Liveness routes are present", func() {
		It("contains handler for GET /health/liveness", func() {
			hndlr, _ := fastHttpRouter.Lookup("GET", livenessPath, nil)
			Expect(hndlr).ShouldNot(BeNil())
		})

		It("contains handler for GET /health/readiness", func() {
			hndlr, _ := fastHttpRouter.Lookup("GET", readinessPath, nil)
			Expect(hndlr).ShouldNot(BeNil())
		})
	})
})
