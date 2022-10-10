// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/sync/interface.go

// Package sync is a generated GoMock package.
package sync

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// SyncFiles mocks base method.
func (m *MockClient) SyncFiles(syncParameters SyncParameters) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncFiles", syncParameters)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncFiles indicates an expected call of SyncFiles.
func (mr *MockClientMockRecorder) SyncFiles(syncParameters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncFiles", reflect.TypeOf((*MockClient)(nil).SyncFiles), syncParameters)
}
