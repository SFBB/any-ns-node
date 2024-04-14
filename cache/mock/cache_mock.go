// Code generated by MockGen. DO NOT EDIT.
// Source: cache/cache.go
//
// Generated by this command:
//
//	mockgen -source=cache/cache.go
//

// Package mock_cache is a generated GoMock package.
package mock_cache

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	nameserviceproto "github.com/anyproto/any-sync/nameservice/nameserviceproto"
	gomock "go.uber.org/mock/gomock"
)

// MockCacheService is a mock of CacheService interface.
type MockCacheService struct {
	ctrl     *gomock.Controller
	recorder *MockCacheServiceMockRecorder
}

// MockCacheServiceMockRecorder is the mock recorder for MockCacheService.
type MockCacheServiceMockRecorder struct {
	mock *MockCacheService
}

// NewMockCacheService creates a new mock instance.
func NewMockCacheService(ctrl *gomock.Controller) *MockCacheService {
	mock := &MockCacheService{ctrl: ctrl}
	mock.recorder = &MockCacheServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheService) EXPECT() *MockCacheServiceMockRecorder {
	return m.recorder
}

// GetNameByAddress mocks base method.
func (m *MockCacheService) GetNameByAddress(ctx context.Context, in *nameserviceproto.NameByAddressRequest) (*nameserviceproto.NameByAddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNameByAddress", ctx, in)
	ret0, _ := ret[0].(*nameserviceproto.NameByAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNameByAddress indicates an expected call of GetNameByAddress.
func (mr *MockCacheServiceMockRecorder) GetNameByAddress(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNameByAddress", reflect.TypeOf((*MockCacheService)(nil).GetNameByAddress), ctx, in)
}

// GetNameByAnyId mocks base method.
func (m *MockCacheService) GetNameByAnyId(ctx context.Context, in *nameserviceproto.NameByAnyIdRequest) (*nameserviceproto.NameByAddressResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNameByAnyId", ctx, in)
	ret0, _ := ret[0].(*nameserviceproto.NameByAddressResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNameByAnyId indicates an expected call of GetNameByAnyId.
func (mr *MockCacheServiceMockRecorder) GetNameByAnyId(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNameByAnyId", reflect.TypeOf((*MockCacheService)(nil).GetNameByAnyId), ctx, in)
}

// Init mocks base method.
func (m *MockCacheService) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockCacheServiceMockRecorder) Init(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockCacheService)(nil).Init), a)
}

// IsNameAvailable mocks base method.
func (m *MockCacheService) IsNameAvailable(ctx context.Context, in *nameserviceproto.NameAvailableRequest) (*nameserviceproto.NameAvailableResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNameAvailable", ctx, in)
	ret0, _ := ret[0].(*nameserviceproto.NameAvailableResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsNameAvailable indicates an expected call of IsNameAvailable.
func (mr *MockCacheServiceMockRecorder) IsNameAvailable(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNameAvailable", reflect.TypeOf((*MockCacheService)(nil).IsNameAvailable), ctx, in)
}

// Name mocks base method.
func (m *MockCacheService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockCacheServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCacheService)(nil).Name))
}

// UpdateInCache mocks base method.
func (m *MockCacheService) UpdateInCache(ctx context.Context, in *nameserviceproto.NameAvailableRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInCache", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInCache indicates an expected call of UpdateInCache.
func (mr *MockCacheServiceMockRecorder) UpdateInCache(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInCache", reflect.TypeOf((*MockCacheService)(nil).UpdateInCache), ctx, in)
}
