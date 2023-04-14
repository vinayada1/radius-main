// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-radius/radius/pkg/recipes/configloader (interfaces: ConfigurationLoader)

// Package configloader is a generated GoMock package.
package configloader

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	recipes "github.com/project-radius/radius/pkg/recipes"
)

// MockConfigurationLoader is a mock of ConfigurationLoader interface.
type MockConfigurationLoader struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationLoaderMockRecorder
}

// MockConfigurationLoaderMockRecorder is the mock recorder for MockConfigurationLoader.
type MockConfigurationLoaderMockRecorder struct {
	mock *MockConfigurationLoader
}

// NewMockConfigurationLoader creates a new mock instance.
func NewMockConfigurationLoader(ctrl *gomock.Controller) *MockConfigurationLoader {
	mock := &MockConfigurationLoader{ctrl: ctrl}
	mock.recorder = &MockConfigurationLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationLoader) EXPECT() *MockConfigurationLoaderMockRecorder {
	return m.recorder
}

// LoadConfiguration mocks base method.
func (m *MockConfigurationLoader) LoadConfiguration(arg0 context.Context, arg1 recipes.Metadata) (*recipes.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadConfiguration", arg0, arg1)
	ret0, _ := ret[0].(*recipes.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadConfiguration indicates an expected call of LoadConfiguration.
func (mr *MockConfigurationLoaderMockRecorder) LoadConfiguration(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadConfiguration", reflect.TypeOf((*MockConfigurationLoader)(nil).LoadConfiguration), arg0, arg1)
}

// LoadRecipe mocks base method.
func (m *MockConfigurationLoader) LoadRecipe(arg0 context.Context, arg1 recipes.Metadata) (*recipes.Definition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadRecipe", arg0, arg1)
	ret0, _ := ret[0].(*recipes.Definition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadRecipe indicates an expected call of LoadRecipe.
func (mr *MockConfigurationLoaderMockRecorder) LoadRecipe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadRecipe", reflect.TypeOf((*MockConfigurationLoader)(nil).LoadRecipe), arg0, arg1)
}
