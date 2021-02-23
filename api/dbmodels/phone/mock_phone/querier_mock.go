// Code generated by MockGen. DO NOT EDIT.
// Source: dbmodels/phone/querier.go

// Package mock_phone is a generated GoMock package.
package mock_phone

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	phone "github.com/rickypai/web-template/api/dbmodels/phone"
	reflect "reflect"
)

// MockQuerier is a mock of Querier interface
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetByID mocks base method
func (m *MockQuerier) GetByID(ctx context.Context, id int64) (phone.Phone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(phone.Phone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockQuerierMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockQuerier)(nil).GetByID), ctx, id)
}

// ListPhones mocks base method
func (m *MockQuerier) ListPhones(ctx context.Context) ([]phone.Phone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPhones", ctx)
	ret0, _ := ret[0].([]phone.Phone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPhones indicates an expected call of ListPhones
func (mr *MockQuerierMockRecorder) ListPhones(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPhones", reflect.TypeOf((*MockQuerier)(nil).ListPhones), ctx)
}