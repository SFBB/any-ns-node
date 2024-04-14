// Code generated by MockGen. DO NOT EDIT.
// Source: account_abstraction/account_abstraction.go
//
// Generated by this command:
//
//	mockgen -source=account_abstraction/account_abstraction.go
//

// Package mock_accountabstraction is a generated GoMock package.
package mock_accountabstraction

import (
	context "context"
	big "math/big"
	reflect "reflect"

	accountabstraction "github.com/anyproto/any-ns-node/account_abstraction"
	app "github.com/anyproto/any-sync/app"
	nameserviceproto "github.com/anyproto/any-sync/nameservice/nameserviceproto"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountAbstractionService is a mock of AccountAbstractionService interface.
type MockAccountAbstractionService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountAbstractionServiceMockRecorder
}

// MockAccountAbstractionServiceMockRecorder is the mock recorder for MockAccountAbstractionService.
type MockAccountAbstractionServiceMockRecorder struct {
	mock *MockAccountAbstractionService
}

// NewMockAccountAbstractionService creates a new mock instance.
func NewMockAccountAbstractionService(ctrl *gomock.Controller) *MockAccountAbstractionService {
	mock := &MockAccountAbstractionService{ctrl: ctrl}
	mock.recorder = &MockAccountAbstractionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountAbstractionService) EXPECT() *MockAccountAbstractionServiceMockRecorder {
	return m.recorder
}

// AdminMintAccessTokens mocks base method.
func (m *MockAccountAbstractionService) AdminMintAccessTokens(ctx context.Context, scw common.Address, amount *big.Int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminMintAccessTokens", ctx, scw, amount)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminMintAccessTokens indicates an expected call of AdminMintAccessTokens.
func (mr *MockAccountAbstractionServiceMockRecorder) AdminMintAccessTokens(ctx, scw, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminMintAccessTokens", reflect.TypeOf((*MockAccountAbstractionService)(nil).AdminMintAccessTokens), ctx, scw, amount)
}

// AdminNameRegister mocks base method.
func (m *MockAccountAbstractionService) AdminNameRegister(ctx context.Context, in *nameserviceproto.NameRegisterRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminNameRegister", ctx, in)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminNameRegister indicates an expected call of AdminNameRegister.
func (mr *MockAccountAbstractionServiceMockRecorder) AdminNameRegister(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminNameRegister", reflect.TypeOf((*MockAccountAbstractionService)(nil).AdminNameRegister), ctx, in)
}

// GetDataNameRegister mocks base method.
func (m *MockAccountAbstractionService) GetDataNameRegister(ctx context.Context, in *nameserviceproto.NameRegisterRequest) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataNameRegister", ctx, in)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDataNameRegister indicates an expected call of GetDataNameRegister.
func (mr *MockAccountAbstractionServiceMockRecorder) GetDataNameRegister(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataNameRegister", reflect.TypeOf((*MockAccountAbstractionService)(nil).GetDataNameRegister), ctx, in)
}

// GetDataNameRegisterForSpace mocks base method.
func (m *MockAccountAbstractionService) GetDataNameRegisterForSpace(ctx context.Context, in *nameserviceproto.NameRegisterForSpaceRequest) ([]byte, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataNameRegisterForSpace", ctx, in)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].([]byte)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetDataNameRegisterForSpace indicates an expected call of GetDataNameRegisterForSpace.
func (mr *MockAccountAbstractionServiceMockRecorder) GetDataNameRegisterForSpace(ctx, in any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataNameRegisterForSpace", reflect.TypeOf((*MockAccountAbstractionService)(nil).GetDataNameRegisterForSpace), ctx, in)
}

// GetNamesCountLeft mocks base method.
func (m *MockAccountAbstractionService) GetNamesCountLeft(ctx context.Context, scw common.Address) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamesCountLeft", ctx, scw)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamesCountLeft indicates an expected call of GetNamesCountLeft.
func (mr *MockAccountAbstractionServiceMockRecorder) GetNamesCountLeft(ctx, scw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamesCountLeft", reflect.TypeOf((*MockAccountAbstractionService)(nil).GetNamesCountLeft), ctx, scw)
}

// GetOperation mocks base method.
func (m *MockAccountAbstractionService) GetOperation(ctx context.Context, operationID string) (*accountabstraction.OperationInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", ctx, operationID)
	ret0, _ := ret[0].(*accountabstraction.OperationInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockAccountAbstractionServiceMockRecorder) GetOperation(ctx, operationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockAccountAbstractionService)(nil).GetOperation), ctx, operationID)
}

// GetSmartWalletAddress mocks base method.
func (m *MockAccountAbstractionService) GetSmartWalletAddress(ctx context.Context, eoa common.Address) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSmartWalletAddress", ctx, eoa)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSmartWalletAddress indicates an expected call of GetSmartWalletAddress.
func (mr *MockAccountAbstractionServiceMockRecorder) GetSmartWalletAddress(ctx, eoa any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSmartWalletAddress", reflect.TypeOf((*MockAccountAbstractionService)(nil).GetSmartWalletAddress), ctx, eoa)
}

// Init mocks base method.
func (m *MockAccountAbstractionService) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockAccountAbstractionServiceMockRecorder) Init(a any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockAccountAbstractionService)(nil).Init), a)
}

// IsScwDeployed mocks base method.
func (m *MockAccountAbstractionService) IsScwDeployed(ctx context.Context, scw common.Address) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsScwDeployed", ctx, scw)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsScwDeployed indicates an expected call of IsScwDeployed.
func (mr *MockAccountAbstractionServiceMockRecorder) IsScwDeployed(ctx, scw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsScwDeployed", reflect.TypeOf((*MockAccountAbstractionService)(nil).IsScwDeployed), ctx, scw)
}

// Name mocks base method.
func (m *MockAccountAbstractionService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAccountAbstractionServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAccountAbstractionService)(nil).Name))
}

// SendUserOperation mocks base method.
func (m *MockAccountAbstractionService) SendUserOperation(ctx context.Context, contextData, signedByUserData []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUserOperation", ctx, contextData, signedByUserData)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendUserOperation indicates an expected call of SendUserOperation.
func (mr *MockAccountAbstractionServiceMockRecorder) SendUserOperation(ctx, contextData, signedByUserData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUserOperation", reflect.TypeOf((*MockAccountAbstractionService)(nil).SendUserOperation), ctx, contextData, signedByUserData)
}
