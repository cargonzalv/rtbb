package observability

import (
	"github.com/buaazp/fasthttprouter"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Observability Routes", func() {
	var (
		fastHttpRouter *fasthttprouter.Router
	)

	BeforeEach(func() {
		fastHttpRouter = fasthttprouter.New()
		s := server{}
		s.AddRoutes(fastHttpRouter)
	})

	It("contains handler for GET /metrics", func() {
		hndlr, _ := fastHttpRouter.Lookup("GET", "/metrics", nil)
		Expect(hndlr).ShouldNot(BeNil())
	})
})
