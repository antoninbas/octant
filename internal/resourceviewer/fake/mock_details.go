// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware/octant/internal/resourceviewer (interfaces: Details)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	component "github.com/vmware/octant/pkg/view/component"
	reflect "reflect"
)

// MockDetails is a mock of Details interface
type MockDetails struct {
	ctrl     *gomock.Controller
	recorder *MockDetailsMockRecorder
}

// MockDetailsMockRecorder is the mock recorder for MockDetails
type MockDetailsMockRecorder struct {
	mock *MockDetails
}

// NewMockDetails creates a new mock instance
func NewMockDetails(ctrl *gomock.Controller) *MockDetails {
	mock := &MockDetails{ctrl: ctrl}
	mock.recorder = &MockDetailsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDetails) EXPECT() *MockDetailsMockRecorder {
	return m.recorder
}

// AdjacencyList mocks base method
func (m *MockDetails) AdjacencyList() (*component.AdjList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdjacencyList")
	ret0, _ := ret[0].(*component.AdjList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdjacencyList indicates an expected call of AdjacencyList
func (mr *MockDetailsMockRecorder) AdjacencyList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdjacencyList", reflect.TypeOf((*MockDetails)(nil).AdjacencyList))
}

// Nodes mocks base method
func (m *MockDetails) Nodes(arg0 context.Context) (component.Nodes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nodes", arg0)
	ret0, _ := ret[0].(component.Nodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Nodes indicates an expected call of Nodes
func (mr *MockDetailsMockRecorder) Nodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nodes", reflect.TypeOf((*MockDetails)(nil).Nodes), arg0)
}
