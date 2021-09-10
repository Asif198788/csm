// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dell/csm-deployment/store (interfaces: ModuleTypeStoreInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/dell/csm-deployment/model"
	gomock "github.com/golang/mock/gomock"
)

// MockModuleTypeStoreInterface is a mock of ModuleTypeStoreInterface interface.
type MockModuleTypeStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockModuleTypeStoreInterfaceMockRecorder
}

// MockModuleTypeStoreInterfaceMockRecorder is the mock recorder for MockModuleTypeStoreInterface.
type MockModuleTypeStoreInterfaceMockRecorder struct {
	mock *MockModuleTypeStoreInterface
}

// NewMockModuleTypeStoreInterface creates a new mock instance.
func NewMockModuleTypeStoreInterface(ctrl *gomock.Controller) *MockModuleTypeStoreInterface {
	mock := &MockModuleTypeStoreInterface{ctrl: ctrl}
	mock.recorder = &MockModuleTypeStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModuleTypeStoreInterface) EXPECT() *MockModuleTypeStoreInterfaceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockModuleTypeStoreInterface) GetAll() ([]model.ModuleType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.ModuleType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockModuleTypeStoreInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockModuleTypeStoreInterface)(nil).GetAll))
}

// GetAllByID mocks base method.
func (m *MockModuleTypeStoreInterface) GetAllByID(arg0 ...uint) ([]model.ModuleType, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAllByID", varargs...)
	ret0, _ := ret[0].([]model.ModuleType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByID indicates an expected call of GetAllByID.
func (mr *MockModuleTypeStoreInterfaceMockRecorder) GetAllByID(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByID", reflect.TypeOf((*MockModuleTypeStoreInterface)(nil).GetAllByID), arg0...)
}

// GetByID mocks base method.
func (m *MockModuleTypeStoreInterface) GetByID(arg0 uint) (*model.ModuleType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*model.ModuleType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockModuleTypeStoreInterfaceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockModuleTypeStoreInterface)(nil).GetByID), arg0)
}