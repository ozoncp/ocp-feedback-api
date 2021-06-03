// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-feedback-api/internal/alarmer (interfaces: Alarmer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAlarmer is a mock of Alarmer interface.
type MockAlarmer struct {
	ctrl     *gomock.Controller
	recorder *MockAlarmerMockRecorder
}

// MockAlarmerMockRecorder is the mock recorder for MockAlarmer.
type MockAlarmerMockRecorder struct {
	mock *MockAlarmer
}

// NewMockAlarmer creates a new mock instance.
func NewMockAlarmer(ctrl *gomock.Controller) *MockAlarmer {
	mock := &MockAlarmer{ctrl: ctrl}
	mock.recorder = &MockAlarmerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlarmer) EXPECT() *MockAlarmerMockRecorder {
	return m.recorder
}

// Alarm mocks base method.
func (m *MockAlarmer) Alarm() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alarm")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Alarm indicates an expected call of Alarm.
func (mr *MockAlarmerMockRecorder) Alarm() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alarm", reflect.TypeOf((*MockAlarmer)(nil).Alarm))
}
