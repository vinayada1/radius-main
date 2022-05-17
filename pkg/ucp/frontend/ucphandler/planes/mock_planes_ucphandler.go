// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/ucp/frontend/ucphandler/planes (interfaces: PlanesUCPHandler)

// Package planes is a generated GoMock package.
package planes

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	rest "github.com/project-radius/radius/pkg/ucp/rest"
	store "github.com/project-radius/radius/pkg/ucp/store"
)

// MockPlanesUCPHandler is a mock of PlanesUCPHandler interface.
type MockPlanesUCPHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPlanesUCPHandlerMockRecorder
}

// MockPlanesUCPHandlerMockRecorder is the mock recorder for MockPlanesUCPHandler.
type MockPlanesUCPHandlerMockRecorder struct {
	mock *MockPlanesUCPHandler
}

// NewMockPlanesUCPHandler creates a new mock instance.
func NewMockPlanesUCPHandler(ctrl *gomock.Controller) *MockPlanesUCPHandler {
	mock := &MockPlanesUCPHandler{ctrl: ctrl}
	mock.recorder = &MockPlanesUCPHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanesUCPHandler) EXPECT() *MockPlanesUCPHandlerMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockPlanesUCPHandler) CreateOrUpdate(arg0 context.Context, arg1 store.StorageClient, arg2 []byte, arg3 string) (rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockPlanesUCPHandlerMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockPlanesUCPHandler)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// DeleteByID mocks base method.
func (m *MockPlanesUCPHandler) DeleteByID(arg0 context.Context, arg1 store.StorageClient, arg2 string) (rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockPlanesUCPHandlerMockRecorder) DeleteByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockPlanesUCPHandler)(nil).DeleteByID), arg0, arg1, arg2)
}

// GetByID mocks base method.
func (m *MockPlanesUCPHandler) GetByID(arg0 context.Context, arg1 store.StorageClient, arg2 string) (rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockPlanesUCPHandlerMockRecorder) GetByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockPlanesUCPHandler)(nil).GetByID), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockPlanesUCPHandler) List(arg0 context.Context, arg1 store.StorageClient, arg2 string) (rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPlanesUCPHandlerMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlanesUCPHandler)(nil).List), arg0, arg1, arg2)
}

// ProxyRequest mocks base method.
func (m *MockPlanesUCPHandler) ProxyRequest(arg0 context.Context, arg1 store.StorageClient, arg2 http.ResponseWriter, arg3 *http.Request, arg4 string) (rest.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxyRequest", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(rest.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProxyRequest indicates an expected call of ProxyRequest.
func (mr *MockPlanesUCPHandlerMockRecorder) ProxyRequest(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyRequest", reflect.TypeOf((*MockPlanesUCPHandler)(nil).ProxyRequest), arg0, arg1, arg2, arg3, arg4)
}