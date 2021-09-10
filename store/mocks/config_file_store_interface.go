// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dell/csm-deployment/store (interfaces: ConfigFileStoreInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/dell/csm-deployment/model"
	gomock "github.com/golang/mock/gomock"
)

// MockConfigFileStoreInterface is a mock of ConfigFileStoreInterface interface.
type MockConfigFileStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConfigFileStoreInterfaceMockRecorder
}

// MockConfigFileStoreInterfaceMockRecorder is the mock recorder for MockConfigFileStoreInterface.
type MockConfigFileStoreInterfaceMockRecorder struct {
	mock *MockConfigFileStoreInterface
}

// NewMockConfigFileStoreInterface creates a new mock instance.
func NewMockConfigFileStoreInterface(ctrl *gomock.Controller) *MockConfigFileStoreInterface {
	mock := &MockConfigFileStoreInterface{ctrl: ctrl}
	mock.recorder = &MockConfigFileStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigFileStoreInterface) EXPECT() *MockConfigFileStoreInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockConfigFileStoreInterface) Create(arg0 *model.ConfigFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockConfigFileStoreInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConfigFileStoreInterface)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockConfigFileStoreInterface) Delete(arg0 *model.ConfigFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockConfigFileStoreInterfaceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockConfigFileStoreInterface)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockConfigFileStoreInterface) GetAll() ([]model.ConfigFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.ConfigFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockConfigFileStoreInterfaceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockConfigFileStoreInterface)(nil).GetAll))
}

// GetAllByName mocks base method.
func (m *MockConfigFileStoreInterface) GetAllByName(arg0 string) ([]model.ConfigFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByName", arg0)
	ret0, _ := ret[0].([]model.ConfigFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByName indicates an expected call of GetAllByName.
func (mr *MockConfigFileStoreInterfaceMockRecorder) GetAllByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByName", reflect.TypeOf((*MockConfigFileStoreInterface)(nil).GetAllByName), arg0)
}

// GetByID mocks base method.
func (m *MockConfigFileStoreInterface) GetByID(arg0 uint) (*model.ConfigFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*model.ConfigFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockConfigFileStoreInterfaceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockConfigFileStoreInterface)(nil).GetByID), arg0)
}

// Update mocks base method.
func (m *MockConfigFileStoreInterface) Update(arg0 *model.ConfigFile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockConfigFileStoreInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockConfigFileStoreInterface)(nil).Update), arg0)
}