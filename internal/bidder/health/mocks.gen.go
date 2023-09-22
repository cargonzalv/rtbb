// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package health is a generated GoMock package.
package health

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	fasthttp "github.com/valyala/fasthttp"
)

// MockObserver is a mock of Observer interface.
type MockObserver struct {
	ctrl     *gomock.Controller
	recorder *MockObserverMockRecorder
}

// MockObserverMockRecorder is the mock recorder for MockObserver.
type MockObserverMockRecorder struct {
	mock *MockObserver
}

// NewMockObserver creates a new mock instance.
func NewMockObserver(ctrl *gomock.Controller) *MockObserver {
	mock := &MockObserver{ctrl: ctrl}
	mock.recorder = &MockObserverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObserver) EXPECT() *MockObserverMockRecorder {
	return m.recorder
}

// IsLive mocks base method.
func (m *MockObserver) IsLive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLive indicates an expected call of IsLive.
func (mr *MockObserverMockRecorder) IsLive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLive", reflect.TypeOf((*MockObserver)(nil).IsLive))
}

// IsReady mocks base method.
func (m *MockObserver) IsReady() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReady indicates an expected call of IsReady.
func (mr *MockObserverMockRecorder) IsReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockObserver)(nil).IsReady))
}

// StartMonitors mocks base method.
func (m *MockObserver) StartMonitors(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartMonitors", arg0)
}

// StartMonitors indicates an expected call of StartMonitors.
func (mr *MockObserverMockRecorder) StartMonitors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartMonitors", reflect.TypeOf((*MockObserver)(nil).StartMonitors), arg0)
}

// StopMonitors mocks base method.
func (m *MockObserver) StopMonitors() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopMonitors")
}

// StopMonitors indicates an expected call of StopMonitors.
func (mr *MockObserverMockRecorder) StopMonitors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopMonitors", reflect.TypeOf((*MockObserver)(nil).StopMonitors))
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

// GetLivenessHandler mocks base method.
func (m *MockHandler) GetLivenessHandler(ctx *fasthttp.RequestCtx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetLivenessHandler", ctx)
}

// GetLivenessHandler indicates an expected call of GetLivenessHandler.
func (mr *MockHandlerMockRecorder) GetLivenessHandler(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLivenessHandler", reflect.TypeOf((*MockHandler)(nil).GetLivenessHandler), ctx)
}

// GetReadinessHandler mocks base method.
func (m *MockHandler) GetReadinessHandler(ctx *fasthttp.RequestCtx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetReadinessHandler", ctx)
}

// GetReadinessHandler indicates an expected call of GetReadinessHandler.
func (mr *MockHandlerMockRecorder) GetReadinessHandler(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadinessHandler", reflect.TypeOf((*MockHandler)(nil).GetReadinessHandler), ctx)
}
