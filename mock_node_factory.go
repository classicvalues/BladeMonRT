// Code generated by MockGen. DO NOT EDIT.
// Source: ./node_factory.go

// Package main is a generated GoMock package.
package main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nodes "github.com/microsoft/BladeMonRT/nodes"
)

// MockInterfaceNodeFactory is a mock of InterfaceNodeFactory interface.
type MockInterfaceNodeFactory struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceNodeFactoryMockRecorder
}

// MockInterfaceNodeFactoryMockRecorder is the mock recorder for MockInterfaceNodeFactory.
type MockInterfaceNodeFactoryMockRecorder struct {
	mock *MockInterfaceNodeFactory
}

// NewMockInterfaceNodeFactory creates a new mock instance.
func NewMockInterfaceNodeFactory(ctrl *gomock.Controller) *MockInterfaceNodeFactory {
	mock := &MockInterfaceNodeFactory{ctrl: ctrl}
	mock.recorder = &MockInterfaceNodeFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterfaceNodeFactory) EXPECT() *MockInterfaceNodeFactoryMockRecorder {
	return m.recorder
}

// constructNode mocks base method.
func (m *MockInterfaceNodeFactory) constructNode(typeName string) nodes.InterfaceNode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "constructNode", typeName)
	ret0, _ := ret[0].(nodes.InterfaceNode)
	return ret0
}

// constructNode indicates an expected call of constructNode.
func (mr *MockInterfaceNodeFactoryMockRecorder) constructNode(typeName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "constructNode", reflect.TypeOf((*MockInterfaceNodeFactory)(nil).constructNode), typeName)
}
