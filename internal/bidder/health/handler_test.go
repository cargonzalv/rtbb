package health

import (
	"bufio"
	"bytes"

	"github.com/adgear/go-commons/pkg/log"
	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

const host = "greenfield.com"
const livenessPath = "/health/liveness"
const readinessPath = "/health/readiness"
const contentTypeHeader = "Content-Type"
const jsonType = "application/json"

var _ = Describe("Handler", func() {
	var (
		mockCtrl     *gomock.Controller
		mockLogger   *log.MockService
		mockObserver *MockObserver
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockLogger = log.NewMockService(mockCtrl)
		mockObserver = NewMockObserver(mockCtrl)
	})

	execRequest := func(endpoint string) (resp fasthttp.Response, err error) {
		var (
			ctx fasthttp.RequestCtx
			req fasthttp.Request
		)
		subject := &controller{
			observer: mockObserver,
			logger:   mockLogger,
		}

		req.Header.SetHost(host)
		req.SetRequestURI(endpoint)
		ctx.Init(&req, nil, nil)
		if endpoint == livenessPath {
			subject.GetLivenessHandler(&ctx)
		} else {
			subject.GetReadinessHandler(&ctx)
		}
		s := ctx.Response.String()
		br := bufio.NewReader(bytes.NewBufferString(s))
		err = resp.Read(br)
		Expect(err).Should(BeNil())
		return
	}

	Context("the `/health/liveness`", func() {
		When("the obeserver.IsLive() returns true", func() {
			BeforeEach(func() {
				mockObserver.EXPECT().IsLive().Return(true)
			})

			It("responses with HTTP status 200", func() {
				resp, _ := execRequest(livenessPath)

				Expect(resp.StatusCode()).To(Equal(fasthttp.StatusOK))
			})

			It("responses with `{is_live: true}` body", func() {
				resp, _ := execRequest(livenessPath)

				expectedBody := []byte(`{"is_live": true}`)
				Expect(resp.Header.Peek(contentTypeHeader)).To(Equal([]byte(jsonType)))
				Expect(resp.Body()).To(Equal(expectedBody))
			})
		})

		When("the obeserver.IsLive() returns false", func() {
			BeforeEach(func() {
				mockObserver.EXPECT().IsLive().Return(false)
			})

			It("responses with HTTP status 503", func() {
				resp, _ := execRequest(livenessPath)

				Expect(resp.StatusCode()).To(Equal(fasthttp.StatusServiceUnavailable))
			})

			It("responses with `{is_live: false}` body", func() {
				resp, _ := execRequest(livenessPath)

				expectedBody := []byte(`{"is_live": false}`)
				Expect(resp.Header.Peek(contentTypeHeader)).To(Equal([]byte(jsonType)))
				Expect(resp.Body()).To(Equal(expectedBody))
			})
		})
	})

	Context("the `/health/readiness`", func() {
		When("the obeserver.IsReady() returns true", func() {
			BeforeEach(func() {
				mockObserver.EXPECT().IsReady().Return(true)
			})

			It("responses with HTTP status 200", func() {
				resp, _ := execRequest(readinessPath)

				Expect(resp.StatusCode()).To(Equal(fasthttp.StatusOK))
			})

			It("responses with `{is_ready: true}` body", func() {
				resp, _ := execRequest(readinessPath)

				expectedBody := []byte(`{"is_ready": true}`)
				Expect(resp.Header.Peek(contentTypeHeader)).To(Equal([]byte(jsonType)))
				Expect(resp.Body()).To(Equal(expectedBody))
			})
		})

		When("the obeserver.IsReady() returns false", func() {
			BeforeEach(func() {
				mockObserver.EXPECT().IsReady().Return(false)
			})

			It("responses with HTTP status 503", func() {
				resp, _ := execRequest(readinessPath)

				Expect(resp.StatusCode()).To(Equal(fasthttp.StatusServiceUnavailable))
			})

			It("responses with `{is_ready: false}` body", func() {
				resp, _ := execRequest(readinessPath)

				expectedBody := []byte(`{"is_ready": false}`)
				Expect(resp.Header.Peek(contentTypeHeader)).To(Equal([]byte(jsonType)))
				Expect(resp.Body()).To(Equal(expectedBody))
			})
		})
	})
})
