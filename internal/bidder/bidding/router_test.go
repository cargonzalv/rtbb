package bidding

import (
	"github.com/buaazp/fasthttprouter"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	gomock "go.uber.org/mock/gomock"
)

var _ = Describe("Bidding Routes", func() {
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

	Context("bidj route is present", func() {
		It("contains handler for POST /bidj", func() {
			hndlr, _ := fastHttpRouter.Lookup("POST", "/bidj", nil)
			Expect(hndlr).ShouldNot(BeNil())
		})
	})
})
