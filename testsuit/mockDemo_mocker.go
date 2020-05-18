// Code generated by MockGen. DO NOT EDIT.
// Source: testsuit/mockDemo.go

// Package testsuit is a generated GoMock package.
package testsuit

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPerson is a mock of Person interface
type MockPerson struct {
	ctrl     *gomock.Controller
	recorder *MockPersonMockRecorder
}

// MockPersonMockRecorder is the mock recorder for MockPerson
type MockPersonMockRecorder struct {
	mock *MockPerson
}

// NewMockPerson creates a new mock instance
func NewMockPerson(ctrl *gomock.Controller) *MockPerson {
	mock := &MockPerson{ctrl: ctrl}
	mock.recorder = &MockPersonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPerson) EXPECT() *MockPersonMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockPerson) Get() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockPersonMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPerson)(nil).Get))
}