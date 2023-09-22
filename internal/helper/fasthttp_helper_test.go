package helper

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
)

var _ = Describe("fasthttp_helper", func() {

	It("successfully write data in http ctx", func() {

		body := []byte(`{"Id":10655,"name":"ABC iView"}`)
		ctx := &fasthttp.RequestCtx{}
		WriteJsonBytesResponse(ctx, body)
		Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusOK))
		Expect(string(ctx.Response.Header.ContentType())).To(Equal("application/json"))
		Expect(ctx.Response.Body()).To(Equal(body))
	})
})
