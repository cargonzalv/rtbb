package handler

import (
	"bufio"
	"bytes"
	"github.com/adgear/go-commons/pkg/log"
	"github.com/adgear/go-commons/pkg/metric"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

const host = "greenfield.com"
const infoPath = "/info"
const contentTypeHeader = "Content-Type"

var _ = Describe("Info and Notfound Handler", func() {

	var (
		mockCtrl          *gomock.Controller
		mockMetricService *metric.MockService
		mockLogService    *log.MockService
	)
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockMetricService = metric.NewMockService(mockCtrl)
		mockLogService = log.NewMockService(mockCtrl)
	})

	execRequest := func(endpoint string) (resp fasthttp.Response, err error) {
		var (
			ctx fasthttp.RequestCtx
			req fasthttp.Request
		)

		req.Header.SetHost(host)
		req.SetRequestURI(endpoint)
		ctx.Init(&req, nil, nil)
		if endpoint == infoPath {
			info := infoHandler("")
			info(&ctx)
		} else {
			notFound := notFoundHandler(mockLogService, mockMetricService, "", "", 0.0)
			notFound(&ctx)
		}
		s := ctx.Response.String()
		br := bufio.NewReader(bytes.NewBufferString(s))
		err = resp.Read(br)
		Expect(err).Should(BeNil())
		return
	}

	Context("the `/info` handler", func() {
		It("responses with HTTP status 200", func() {
			resp, _ := execRequest(infoPath)
			Expect(resp.StatusCode()).To(Equal(fasthttp.StatusOK))
			Expect(resp.Header.Peek(contentTypeHeader)).To(Equal([]byte("text/html")))
			Expect(resp.Body()).ShouldNot(BeNil())
			Expect(len(resp.Body())).ShouldNot(BeZero())
		})
	})

	Context("the notFound Handler", func() {
		It("responses with HTTP status 404", func() {
			mockMetricService.EXPECT().Count(gomock.Any(), gomock.Any(), gomock.Any()).Return()
			mockLogService.EXPECT().Warn(gomock.Any(), gomock.Any()).Return()
			resp, _ := execRequest("/missing-route")
			Expect(resp.StatusCode()).To(Equal(fasthttp.StatusNotFound))
		})
	})
})
