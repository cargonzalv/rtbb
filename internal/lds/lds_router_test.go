package lds

import (
	"encoding/json"
	"fmt"

	dsplds "github.com/adgear/dsp-lds/pkg/lds"
	. "github.com/onsi/ginkgo/v2"

	. "github.com/onsi/gomega"
	"github.com/valyala/fasthttp"
	"go.uber.org/mock/gomock"
)

var _ = Describe("LdsRouter", func() {

	var (
		mockCtrl           *gomock.Controller
		mockLocalDataStore *dsplds.MockLocalDataStore
		router             LdsRouter
		ctx                *fasthttp.RequestCtx
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockLocalDataStore = dsplds.NewMockLocalDataStore(mockCtrl)
		router = ProvideLdsRouter(mockLocalDataStore)
		ctx = &fasthttp.RequestCtx{}
	})

	Describe("test GetLdsStateHandler handler", func() {

		When("json marshal successful", func() {

			It("state successful", func() {
				dummyState := make(map[string]interface{})
				dummyState["ad"] = "hello"
				body, _ := json.Marshal(dummyState)
				mockLocalDataStore.EXPECT().GetState(ctx).Return(dummyState, nil)
				router.GetLdsStateHandler(ctx)
				Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusOK))
				Expect(string(ctx.Response.Header.ContentType())).To(Equal("application/json"))
				Expect(ctx.Response.Body()).To(Equal(body))
			})
			It("state failure", func() {
				mockLocalDataStore.EXPECT().GetState(ctx).Return(nil, fmt.Errorf("dummy error test"))
				router.GetLdsStateHandler(ctx)
				Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusInternalServerError))
			})
		})

		When("json marshal error", func() {
			It("chan err pass in dummy state", func() {
				dummyState := make(map[string]interface{})
				dummyState["ad"] = make(chan error)
				mockLocalDataStore.EXPECT().GetState(ctx).Return(dummyState, nil)
				router.GetLdsStateHandler(ctx)
				Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusInternalServerError))
			})
		})
	})

	Describe("test GetLdsValueHandler handler", func() {

		When("json Marshal is successful", func() {

			It("lookup key found", func() {
				body := "testBody"

				jsonBody, _ := json.Marshal(body)
				mockLocalDataStore.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).Return(body, true)
				router.GetLdsValueHandler(ctx)
				Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusOK))
				Expect(string(ctx.Response.Header.ContentType())).To(Equal("application/json"))
				Expect(ctx.Response.Body()).To(Equal(jsonBody))
			})

			It("lookup key not found", func() {
				mockLocalDataStore.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).Return(nil, false)
				router.GetLdsValueHandler(ctx)
				Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusNoContent))
			})

		})

		When("json Marshal is failure", func() {
			It("chan err pass in lookup", func() {
				errChan := make(chan error)
				mockLocalDataStore.EXPECT().Get(ctx, gomock.Any(), gomock.Any()).Return(errChan, true)
				router.GetLdsValueHandler(ctx)
				Expect(ctx.Response.StatusCode()).To(Equal(fasthttp.StatusInternalServerError))
			})
		})

	})

})
