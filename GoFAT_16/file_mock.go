// Code generated by MockGen. DO NOT EDIT.
// Source: file.go

// Package gofat is a generated GoMock package.
package gofat

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockfatFileFs is a mock of fatFileFs interface.
type MockfatFileFs struct {
	ctrl     *gomock.Controller
	recorder *MockfatFileFsMockRecorder
}

// MockfatFileFsMockRecorder is the mock recorder for MockfatFileFs.
type MockfatFileFsMockRecorder struct {
	mock *MockfatFileFs
}

// NewMockfatFileFs creates a new mock instance.
func NewMockfatFileFs(ctrl *gomock.Controller) *MockfatFileFs {
	mock := &MockfatFileFs{ctrl: ctrl}
	mock.recorder = &MockfatFileFsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfatFileFs) EXPECT() *MockfatFileFsMockRecorder {
	return m.recorder
}

// readDir mocks base method.
func (m *MockfatFileFs) readDir(cluster fatEntry) ([]ExtendedEntryHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readDir", cluster)
	ret0, _ := ret[0].([]ExtendedEntryHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// readDir indicates an expected call of readDir.
func (mr *MockfatFileFsMockRecorder) readDir(cluster interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readDir", reflect.TypeOf((*MockfatFileFs)(nil).readDir), cluster)
}

// readFileAt mocks base method.
func (m *MockfatFileFs) readFileAt(cluster fatEntry, fileSize, offset, readSize int64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readFileAt", cluster, fileSize, offset, readSize)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// readFileAt indicates an expected call of readFileAt.
func (mr *MockfatFileFsMockRecorder) readFileAt(cluster, fileSize, offset, readSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readFileAt", reflect.TypeOf((*MockfatFileFs)(nil).readFileAt), cluster, fileSize, offset, readSize)
}

// readRoot mocks base method.
func (m *MockfatFileFs) readRoot() ([]ExtendedEntryHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readRoot")
	ret0, _ := ret[0].([]ExtendedEntryHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// readRoot indicates an expected call of readRoot.
func (mr *MockfatFileFsMockRecorder) readRoot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readRoot", reflect.TypeOf((*MockfatFileFs)(nil).readRoot))
}
