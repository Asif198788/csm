// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dell/csm-deployment/store (interfaces: DriverTypeStoreInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/dell/csm-deployment/model"
	gomock "github.com/golang/mock/gomock"
)

// MockDriverTypeStoreInterface is a mock of DriverTypeStoreInterface interface.
type MockDriverTypeStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDriverTypeStoreInterfaceMockRecorder
}

// MockDriverTypeStoreInterfaceMockRecorder is the mock recorder for MockDriverTypeStoreInterface.
type MockDriverTypeStoreInterfaceMockRecorder struct {
	mock *MockDriverTypeStoreInterface
}

// NewMockDriverTypeStoreInterface creates a new mock instance.
func NewMockDriverTypeStoreInterface(ctrl *gomock.Controller) *MockDriverTypeStoreInterface {
	mock := &MockDriverTypeStoreInterface{ctrl: ctrl}
	mock.recorder = &MockDriverTypeStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriverTypeStoreInterface) EXPECT() *MockDriverTypeStoreInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockDriverTypeStoreInterface) GetAll() ([]model.DriverType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.DriverType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDriverTypeStoreInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDriverTypeStoreInterface)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockDriverTypeStoreInterface) GetByID(arg0 uint) (*model.DriverType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*model.DriverType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDriverTypeStoreInterfaceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDriverTypeStoreInterface)(nil).GetByID), arg0)
}
