// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/common-fate/granted-approvals/pkg/deploy (interfaces: DeployConfigReader)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	deploy "github.com/common-fate/granted-approvals/pkg/deploy"
	gomock "github.com/golang/mock/gomock"
)

// MockDeployConfigReader is a mock of DeployConfigReader interface.
type MockDeployConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockDeployConfigReaderMockRecorder
}

// MockDeployConfigReaderMockRecorder is the mock recorder for MockDeployConfigReader.
type MockDeployConfigReaderMockRecorder struct {
	mock *MockDeployConfigReader
}

// NewMockDeployConfigReader creates a new mock instance.
func NewMockDeployConfigReader(ctrl *gomock.Controller) *MockDeployConfigReader {
	mock := &MockDeployConfigReader{ctrl: ctrl}
	mock.recorder = &MockDeployConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeployConfigReader) EXPECT() *MockDeployConfigReaderMockRecorder {
	return m.recorder
}

// ReadNotifications mocks base method.
func (m *MockDeployConfigReader) ReadNotifications(arg0 context.Context) (deploy.FeatureMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadNotifications", arg0)
	ret0, _ := ret[0].(deploy.FeatureMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadNotifications indicates an expected call of ReadNotifications.
func (mr *MockDeployConfigReaderMockRecorder) ReadNotifications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadNotifications", reflect.TypeOf((*MockDeployConfigReader)(nil).ReadNotifications), arg0)
}

// ReadProviders mocks base method.
func (m *MockDeployConfigReader) ReadProviders(arg0 context.Context) (deploy.ProviderMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProviders", arg0)
	ret0, _ := ret[0].(deploy.ProviderMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProviders indicates an expected call of ReadProviders.
func (mr *MockDeployConfigReaderMockRecorder) ReadProviders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProviders", reflect.TypeOf((*MockDeployConfigReader)(nil).ReadProviders), arg0)
}