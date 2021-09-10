// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dell/csm-deployment/prechecks (interfaces: K8sClientVersionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockK8sClientVersionInterface is a mock of K8sClientVersionInterface interface.
type MockK8sClientVersionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockK8sClientVersionInterfaceMockRecorder
}

// MockK8sClientVersionInterfaceMockRecorder is the mock recorder for MockK8sClientVersionInterface.
type MockK8sClientVersionInterfaceMockRecorder struct {
	mock *MockK8sClientVersionInterface
}

// NewMockK8sClientVersionInterface creates a new mock instance.
func NewMockK8sClientVersionInterface(ctrl *gomock.Controller) *MockK8sClientVersionInterface {
	mock := &MockK8sClientVersionInterface{ctrl: ctrl}
	mock.recorder = &MockK8sClientVersionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockK8sClientVersionInterface) EXPECT() *MockK8sClientVersionInterfaceMockRecorder {
	return m.recorder
}

// GetVersion mocks base method.
func (m *MockK8sClientVersionInterface) GetVersion(arg0 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockK8sClientVersionInterfaceMockRecorder) GetVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockK8sClientVersionInterface)(nil).GetVersion), arg0)
}

// IsOpenShift mocks base method.
func (m *MockK8sClientVersionInterface) IsOpenShift(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOpenShift", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOpenShift indicates an expected call of IsOpenShift.
func (mr *MockK8sClientVersionInterfaceMockRecorder) IsOpenShift(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOpenShift", reflect.TypeOf((*MockK8sClientVersionInterface)(nil).IsOpenShift), arg0)
}
