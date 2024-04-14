// Code generated by MockGen. DO NOT EDIT.
// Source: queue/queue.go
//
// Generated by this command:
//
//	mockgen -source=queue/queue.go
//

// Package mock_queue is a generated GoMock package.
package mock_queue

import (
	context "context"
	reflect "reflect"

	queue "github.com/anyproto/any-ns-node/queue"
	app "github.com/anyproto/any-sync/app"
	nameserviceproto "github.com/anyproto/any-sync/nameservice/nameserviceproto"
	gomock "go.uber.org/mock/gomock"
)

// MockQueueService is a mock of QueueService interface.
type MockQueueService struct {
	ctrl     *gomock.Controller
	recorder *MockQueueServiceMockRecorder
}

// MockQueueServiceMockRecorder is the mock recorder for MockQueueService.
type MockQueueServiceMockRecorder struct {
	mock *MockQueueService
}

// NewMockQueueService creates a new mock instance.
func NewMockQueueService(ctrl *gomock.Controller) *MockQueueService {
	mock := &MockQueueService{ctrl: ctrl}
	mock.recorder = &MockQueueServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueueService) EXPECT() *MockQueueServiceMockRecorder {
	return m.recorder
}

// AddNewRequest mocks base method.
func (m *MockQueueService) AddNewRequest(ctx context.Context, req *nameserviceproto.NameRegisterRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewRequest", ctx, req)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewRequest indicates an expected call of AddNewRequest.
func (mr *MockQueueServiceMockRecorder) AddNewRequest(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewRequest", reflect.TypeOf((*MockQueueService)(nil).AddNewRequest), ctx, req)
}

// Close mocks base method.
func (m *MockQueueService) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockQueueServiceMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockQueueService)(nil).Close), ctx)
}

// FindAndProcessAllItemsInDb mocks base method.
func (m *MockQueueService) FindAndProcessAllItemsInDb(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAndProcessAllItemsInDb", ctx)
}

// FindAndProcessAllItemsInDb indicates an expected call of FindAndProcessAllItemsInDb.
func (mr *MockQueueServiceMockRecorder) FindAndProcessAllItemsInDb(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAndProcessAllItemsInDb", reflect.TypeOf((*MockQueueService)(nil).FindAndProcessAllItemsInDb), ctx)
}

// FindAndProcessAllItemsInDbWithStatus mocks base method.
func (m *MockQueueService) FindAndProcessAllItemsInDbWithStatus(ctx context.Context, status queue.QueueItemStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindAndProcessAllItemsInDbWithStatus", ctx, status)
}

// FindAndProcessAllItemsInDbWithStatus indicates an expected call of FindAndProcessAllItemsInDbWithStatus.
func (mr *MockQueueServiceMockRecorder) FindAndProcessAllItemsInDbWithStatus(ctx, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAndProcessAllItemsInDbWithStatus", reflect.TypeOf((*MockQueueService)(nil).FindAndProcessAllItemsInDbWithStatus), ctx, status)
}

// GetRequestStatus mocks base method.
func (m *MockQueueService) GetRequestStatus(ctx context.Context, operationId int64) (nameserviceproto.OperationState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestStatus", ctx, operationId)
	ret0, _ := ret[0].(nameserviceproto.OperationState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestStatus indicates an expected call of GetRequestStatus.
func (mr *MockQueueServiceMockRecorder) GetRequestStatus(ctx, operationId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestStatus", reflect.TypeOf((*MockQueueService)(nil).GetRequestStatus), ctx, operationId)
}

// Init mocks base method.
func (m *MockQueueService) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockQueueServiceMockRecorder) Init(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockQueueService)(nil).Init), a)
}

// Name mocks base method.
func (m *MockQueueService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockQueueServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockQueueService)(nil).Name))
}

// NameRegisterMoveStateNext mocks base method.
func (m *MockQueueService) NameRegisterMoveStateNext(ctx context.Context, queueItem *queue.QueueItem) (queue.QueueItemStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NameRegisterMoveStateNext", ctx, queueItem)
	ret0, _ := ret[0].(queue.QueueItemStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NameRegisterMoveStateNext indicates an expected call of NameRegisterMoveStateNext.
func (mr *MockQueueServiceMockRecorder) NameRegisterMoveStateNext(ctx, queueItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NameRegisterMoveStateNext", reflect.TypeOf((*MockQueueService)(nil).NameRegisterMoveStateNext), ctx, queueItem)
}

// ProcessItem mocks base method.
func (m *MockQueueService) ProcessItem(ctx context.Context, queueItem *queue.QueueItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessItem", ctx, queueItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessItem indicates an expected call of ProcessItem.
func (mr *MockQueueServiceMockRecorder) ProcessItem(ctx, queueItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessItem", reflect.TypeOf((*MockQueueService)(nil).ProcessItem), ctx, queueItem)
}

// Run mocks base method.
func (m *MockQueueService) Run(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockQueueServiceMockRecorder) Run(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockQueueService)(nil).Run), ctx)
}

// SaveItemToDb mocks base method.
func (m *MockQueueService) SaveItemToDb(ctx context.Context, queueItem *queue.QueueItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveItemToDb", ctx, queueItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveItemToDb indicates an expected call of SaveItemToDb.
func (mr *MockQueueServiceMockRecorder) SaveItemToDb(ctx, queueItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveItemToDb", reflect.TypeOf((*MockQueueService)(nil).SaveItemToDb), ctx, queueItem)
}
