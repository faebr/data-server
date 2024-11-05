// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/datastore/target/netconf/driver.go
//
// Generated by this command:
//
//	mockgen -package=mocknetconf -source=pkg/datastore/target/netconf/driver.go -destination=./mocks/mocknetconf/driver.go
//

// Package mocknetconf is a generated GoMock package.
package mocknetconf

import (
	reflect "reflect"

	types "github.com/sdcio/data-server/pkg/datastore/target/netconf/types"
	gomock "go.uber.org/mock/gomock"
)

// MockDriver is a mock of Driver interface.
type MockDriver struct {
	ctrl     *gomock.Controller
	recorder *MockDriverMockRecorder
	isgomock struct{}
}

// MockDriverMockRecorder is the mock recorder for MockDriver.
type MockDriverMockRecorder struct {
	mock *MockDriver
}

// NewMockDriver creates a new mock instance.
func NewMockDriver(ctrl *gomock.Controller) *MockDriver {
	mock := &MockDriver{ctrl: ctrl}
	mock.recorder = &MockDriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriver) EXPECT() *MockDriverMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDriver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDriverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDriver)(nil).Close))
}

// Commit mocks base method.
func (m *MockDriver) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockDriverMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockDriver)(nil).Commit))
}

// Discard mocks base method.
func (m *MockDriver) Discard() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discard")
	ret0, _ := ret[0].(error)
	return ret0
}

// Discard indicates an expected call of Discard.
func (mr *MockDriverMockRecorder) Discard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discard", reflect.TypeOf((*MockDriver)(nil).Discard))
}

// EditConfig mocks base method.
func (m *MockDriver) EditConfig(target, config string) (*types.NetconfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditConfig", target, config)
	ret0, _ := ret[0].(*types.NetconfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditConfig indicates an expected call of EditConfig.
func (mr *MockDriverMockRecorder) EditConfig(target, config any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditConfig", reflect.TypeOf((*MockDriver)(nil).EditConfig), target, config)
}

// Get mocks base method.
func (m *MockDriver) Get(filter string) (*types.NetconfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", filter)
	ret0, _ := ret[0].(*types.NetconfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockDriverMockRecorder) Get(filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDriver)(nil).Get), filter)
}

// GetConfig mocks base method.
func (m *MockDriver) GetConfig(source, filter string) (*types.NetconfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", source, filter)
	ret0, _ := ret[0].(*types.NetconfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockDriverMockRecorder) GetConfig(source, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockDriver)(nil).GetConfig), source, filter)
}

// IsAlive mocks base method.
func (m *MockDriver) IsAlive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAlive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAlive indicates an expected call of IsAlive.
func (mr *MockDriverMockRecorder) IsAlive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAlive", reflect.TypeOf((*MockDriver)(nil).IsAlive))
}

// Lock mocks base method.
func (m *MockDriver) Lock(target string) (*types.NetconfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", target)
	ret0, _ := ret[0].(*types.NetconfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockDriverMockRecorder) Lock(target any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockDriver)(nil).Lock), target)
}

// Unlock mocks base method.
func (m *MockDriver) Unlock(target string) (*types.NetconfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", target)
	ret0, _ := ret[0].(*types.NetconfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unlock indicates an expected call of Unlock.
func (mr *MockDriverMockRecorder) Unlock(target any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockDriver)(nil).Unlock), target)
}

// Validate mocks base method.
func (m *MockDriver) Validate(source string) (*types.NetconfResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", source)
	ret0, _ := ret[0].(*types.NetconfResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockDriverMockRecorder) Validate(source any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockDriver)(nil).Validate), source)
}
