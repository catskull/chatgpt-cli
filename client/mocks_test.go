// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kardolus/chatgpt-poc/http (interfaces: Caller)

// Package client_test is a generated GoMock package.
package client_test

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCaller is a mock of Caller interface.
type MockCaller struct {
	ctrl     *gomock.Controller
	recorder *MockCallerMockRecorder
}

// MockCallerMockRecorder is the mock recorder for MockCaller.
type MockCallerMockRecorder struct {
	mock *MockCaller
}

// NewMockCaller creates a new mock instance.
func NewMockCaller(ctrl *gomock.Controller) *MockCaller {
	mock := &MockCaller{ctrl: ctrl}
	mock.recorder = &MockCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaller) EXPECT() *MockCallerMockRecorder {
	return m.recorder
}

// Post mocks base method.
func (m *MockCaller) Post(arg0 string, arg1 []byte, arg2 bool) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockCallerMockRecorder) Post(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockCaller)(nil).Post), arg0, arg1, arg2)
}
