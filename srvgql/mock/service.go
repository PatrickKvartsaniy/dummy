// Code generated by MockGen. DO NOT EDIT.
// Source: srvgql/service.go

// Package mock_srvgql is a generated GoMock package.
package mock_srvgql

import (
	models "github.com/akhripko/dummy/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// HealthCheck mocks base method
func (m *MockService) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockServiceMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockService)(nil).HealthCheck))
}

// Hello mocks base method
func (m *MockService) Hello(name string) (*models.HelloMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello", name)
	ret0, _ := ret[0].(*models.HelloMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hello indicates an expected call of Hello
func (mr *MockServiceMockRecorder) Hello(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockService)(nil).Hello), name)
}
