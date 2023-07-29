// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/contracts.go

// Package mock_contracts is a generated GoMock package.
package mock_contracts

import (
	context "context"
	reflect "reflect"

	anytype_crypto "github.com/anyproto/any-ns-node/anytype_crypto"
	app "github.com/anyproto/any-sync/app"
	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	gomock "go.uber.org/mock/gomock"
)

// MockContractsService is a mock of ContractsService interface.
type MockContractsService struct {
	ctrl     *gomock.Controller
	recorder *MockContractsServiceMockRecorder
}

// MockContractsServiceMockRecorder is the mock recorder for MockContractsService.
type MockContractsServiceMockRecorder struct {
	mock *MockContractsService
}

// NewMockContractsService creates a new mock instance.
func NewMockContractsService(ctrl *gomock.Controller) *MockContractsService {
	mock := &MockContractsService{ctrl: ctrl}
	mock.recorder = &MockContractsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractsService) EXPECT() *MockContractsServiceMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockContractsService) Commit(opts *bind.TransactOpts, commitment [32]byte, controller *anytype_crypto.AnytypeRegistrarControllerPrivate) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", opts, commitment, controller)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockContractsServiceMockRecorder) Commit(opts, commitment, controller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockContractsService)(nil).Commit), opts, commitment, controller)
}

// ConnectToController mocks base method.
func (m *MockContractsService) ConnectToController(conn *ethclient.Client) (*anytype_crypto.AnytypeRegistrarControllerPrivate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToController", conn)
	ret0, _ := ret[0].(*anytype_crypto.AnytypeRegistrarControllerPrivate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToController indicates an expected call of ConnectToController.
func (mr *MockContractsServiceMockRecorder) ConnectToController(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToController", reflect.TypeOf((*MockContractsService)(nil).ConnectToController), conn)
}

// ConnectToNamewrapperContract mocks base method.
func (m *MockContractsService) ConnectToNamewrapperContract(conn *ethclient.Client) (*anytype_crypto.AnytypeNameWrapper, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToNamewrapperContract", conn)
	ret0, _ := ret[0].(*anytype_crypto.AnytypeNameWrapper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToNamewrapperContract indicates an expected call of ConnectToNamewrapperContract.
func (mr *MockContractsServiceMockRecorder) ConnectToNamewrapperContract(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToNamewrapperContract", reflect.TypeOf((*MockContractsService)(nil).ConnectToNamewrapperContract), conn)
}

// ConnectToRegistryContract mocks base method.
func (m *MockContractsService) ConnectToRegistryContract(conn *ethclient.Client) (*anytype_crypto.ENSRegistry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToRegistryContract", conn)
	ret0, _ := ret[0].(*anytype_crypto.ENSRegistry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToRegistryContract indicates an expected call of ConnectToRegistryContract.
func (mr *MockContractsServiceMockRecorder) ConnectToRegistryContract(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToRegistryContract", reflect.TypeOf((*MockContractsService)(nil).ConnectToRegistryContract), conn)
}

// ConnectToResolver mocks base method.
func (m *MockContractsService) ConnectToResolver(conn *ethclient.Client) (*anytype_crypto.AnytypeResolver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToResolver", conn)
	ret0, _ := ret[0].(*anytype_crypto.AnytypeResolver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToResolver indicates an expected call of ConnectToResolver.
func (mr *MockContractsServiceMockRecorder) ConnectToResolver(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToResolver", reflect.TypeOf((*MockContractsService)(nil).ConnectToResolver), conn)
}

// CreateEthConnection mocks base method.
func (m *MockContractsService) CreateEthConnection() (*ethclient.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEthConnection")
	ret0, _ := ret[0].(*ethclient.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEthConnection indicates an expected call of CreateEthConnection.
func (mr *MockContractsServiceMockRecorder) CreateEthConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEthConnection", reflect.TypeOf((*MockContractsService)(nil).CreateEthConnection))
}

// GenerateAuthOptsForAdmin mocks base method.
func (m *MockContractsService) GenerateAuthOptsForAdmin(conn *ethclient.Client) (*bind.TransactOpts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAuthOptsForAdmin", conn)
	ret0, _ := ret[0].(*bind.TransactOpts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAuthOptsForAdmin indicates an expected call of GenerateAuthOptsForAdmin.
func (mr *MockContractsServiceMockRecorder) GenerateAuthOptsForAdmin(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAuthOptsForAdmin", reflect.TypeOf((*MockContractsService)(nil).GenerateAuthOptsForAdmin), conn)
}

// GetAdditionalNameInfo mocks base method.
func (m *MockContractsService) GetAdditionalNameInfo(conn *ethclient.Client, currentOwner common.Address, fullName string) (string, string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdditionalNameInfo", conn, currentOwner, fullName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetAdditionalNameInfo indicates an expected call of GetAdditionalNameInfo.
func (mr *MockContractsServiceMockRecorder) GetAdditionalNameInfo(conn, currentOwner, fullName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdditionalNameInfo", reflect.TypeOf((*MockContractsService)(nil).GetAdditionalNameInfo), conn, currentOwner, fullName)
}

// GetOwnerForNamehash mocks base method.
func (m *MockContractsService) GetOwnerForNamehash(client *ethclient.Client, namehash [32]byte) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwnerForNamehash", client, namehash)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwnerForNamehash indicates an expected call of GetOwnerForNamehash.
func (mr *MockContractsServiceMockRecorder) GetOwnerForNamehash(client, namehash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwnerForNamehash", reflect.TypeOf((*MockContractsService)(nil).GetOwnerForNamehash), client, namehash)
}

// Init mocks base method.
func (m *MockContractsService) Init(a *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockContractsServiceMockRecorder) Init(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockContractsService)(nil).Init), a)
}

// MakeCommitment mocks base method.
func (m *MockContractsService) MakeCommitment(nameFirstPart string, registrantAccount common.Address, secret [32]byte, controller *anytype_crypto.AnytypeRegistrarControllerPrivate, fullName, ownerAnyAddr, spaceId string) ([32]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeCommitment", nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId)
	ret0, _ := ret[0].([32]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeCommitment indicates an expected call of MakeCommitment.
func (mr *MockContractsServiceMockRecorder) MakeCommitment(nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeCommitment", reflect.TypeOf((*MockContractsService)(nil).MakeCommitment), nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId)
}

// Name mocks base method.
func (m *MockContractsService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockContractsServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockContractsService)(nil).Name))
}

// Register mocks base method.
func (m *MockContractsService) Register(authOpts *bind.TransactOpts, nameFirstPart string, registrantAccount common.Address, secret [32]byte, controller *anytype_crypto.AnytypeRegistrarControllerPrivate, fullName, ownerAnyAddr, spaceId string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", authOpts, nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockContractsServiceMockRecorder) Register(authOpts, nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockContractsService)(nil).Register), authOpts, nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId)
}

// WaitMined mocks base method.
func (m *MockContractsService) WaitMined(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitMined", ctx, client, tx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitMined indicates an expected call of WaitMined.
func (mr *MockContractsServiceMockRecorder) WaitMined(ctx, client, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitMined", reflect.TypeOf((*MockContractsService)(nil).WaitMined), ctx, client, tx)
}
