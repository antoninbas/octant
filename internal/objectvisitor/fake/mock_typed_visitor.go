// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware/octant/internal/objectvisitor (interfaces: TypedVisitor)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	objectvisitor "github.com/vmware/octant/internal/objectvisitor"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	reflect "reflect"
)

// MockTypedVisitor is a mock of TypedVisitor interface
type MockTypedVisitor struct {
	ctrl     *gomock.Controller
	recorder *MockTypedVisitorMockRecorder
}

// MockTypedVisitorMockRecorder is the mock recorder for MockTypedVisitor
type MockTypedVisitorMockRecorder struct {
	mock *MockTypedVisitor
}

// NewMockTypedVisitor creates a new mock instance
func NewMockTypedVisitor(ctrl *gomock.Controller) *MockTypedVisitor {
	mock := &MockTypedVisitor{ctrl: ctrl}
	mock.recorder = &MockTypedVisitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTypedVisitor) EXPECT() *MockTypedVisitorMockRecorder {
	return m.recorder
}

// Supports mocks base method
func (m *MockTypedVisitor) Supports() schema.GroupVersionKind {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supports")
	ret0, _ := ret[0].(schema.GroupVersionKind)
	return ret0
}

// Supports indicates an expected call of Supports
func (mr *MockTypedVisitorMockRecorder) Supports() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supports", reflect.TypeOf((*MockTypedVisitor)(nil).Supports))
}

// Visit mocks base method
func (m *MockTypedVisitor) Visit(arg0 context.Context, arg1 *unstructured.Unstructured, arg2 objectvisitor.ObjectHandler, arg3 objectvisitor.Visitor, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Visit", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Visit indicates an expected call of Visit
func (mr *MockTypedVisitorMockRecorder) Visit(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockTypedVisitor)(nil).Visit), arg0, arg1, arg2, arg3, arg4)
}
