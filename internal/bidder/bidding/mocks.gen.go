// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package bidding is a generated GoMock package.
package bidding

import (
	reflect "reflect"

	openrtb2 "github.com/prebid/openrtb/v20/openrtb2"
	fasthttp "github.com/valyala/fasthttp"
	gomock "go.uber.org/mock/gomock"
)

// MockBidder is a mock of Bidder interface.
type MockBidder struct {
	ctrl     *gomock.Controller
	recorder *MockBidderMockRecorder
}

// MockBidderMockRecorder is the mock recorder for MockBidder.
type MockBidderMockRecorder struct {
	mock *MockBidder
}

// NewMockBidder creates a new mock instance.
func NewMockBidder(ctrl *gomock.Controller) *MockBidder {
	mock := &MockBidder{ctrl: ctrl}
	mock.recorder = &MockBidderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBidder) EXPECT() *MockBidderMockRecorder {
	return m.recorder
}

// Bid mocks base method.
func (m *MockBidder) Bid(req *openrtb2.BidRequest) (*openrtb2.BidResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bid", req)
	ret0, _ := ret[0].(*openrtb2.BidResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Bid indicates an expected call of Bid.
func (mr *MockBidderMockRecorder) Bid(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bid", reflect.TypeOf((*MockBidder)(nil).Bid), req)
}

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// GetTvTile mocks base method.
func (m *MockHandler) GetTvTile(ctx *fasthttp.RequestCtx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetTvTile", ctx)
}

// GetTvTile indicates an expected call of GetTvTile.
func (mr *MockHandlerMockRecorder) GetTvTile(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTvTile", reflect.TypeOf((*MockHandler)(nil).GetTvTile), ctx)
}

// PostBidJson mocks base method.
func (m *MockHandler) PostBidJson(ctx *fasthttp.RequestCtx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PostBidJson", ctx)
}

// PostBidJson indicates an expected call of PostBidJson.
func (mr *MockHandlerMockRecorder) PostBidJson(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostBidJson", reflect.TypeOf((*MockHandler)(nil).PostBidJson), ctx)
}
