// Code generated by MockGen. DO NOT EDIT.
// Source: contracts/contracts.go

// Package mock_contracts is a generated GoMock package.
package mock_contracts

import (
	context "context"
	big "math/big"
	reflect "reflect"

	anytype_crypto "github.com/anyproto/any-ns-node/anytype_crypto"
	app "github.com/anyproto/any-sync/app"
	ethereum "github.com/ethereum/go-ethereum"
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

// CalculateTxParams mocks base method.
func (m *MockContractsService) CalculateTxParams(conn *ethclient.Client, address common.Address) (*big.Int, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateTxParams", conn, address)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CalculateTxParams indicates an expected call of CalculateTxParams.
func (mr *MockContractsServiceMockRecorder) CalculateTxParams(conn, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateTxParams", reflect.TypeOf((*MockContractsService)(nil).CalculateTxParams), conn, address)
}

// CallContract mocks base method.
func (m *MockContractsService) CallContract(ctx context.Context, msg ethereum.CallMsg) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, msg)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractsServiceMockRecorder) CallContract(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractsService)(nil).CallContract), ctx, msg)
}

// Commit mocks base method.
func (m *MockContractsService) Commit(ctx context.Context, conn *ethclient.Client, opts *bind.TransactOpts, commitment [32]byte, controller *anytype_crypto.AnytypeRegistrarControllerPrivate) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, conn, opts, commitment, controller)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockContractsServiceMockRecorder) Commit(ctx, conn, opts, commitment, controller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockContractsService)(nil).Commit), ctx, conn, opts, commitment, controller)
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

// ConnectToRegistrar mocks base method.
func (m *MockContractsService) ConnectToRegistrar(conn *ethclient.Client) (*anytype_crypto.AnytypeRegistrarImplementation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectToRegistrar", conn)
	ret0, _ := ret[0].(*anytype_crypto.AnytypeRegistrarImplementation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectToRegistrar indicates an expected call of ConnectToRegistrar.
func (mr *MockContractsServiceMockRecorder) ConnectToRegistrar(conn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectToRegistrar", reflect.TypeOf((*MockContractsService)(nil).ConnectToRegistrar), conn)
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
func (m *MockContractsService) GetAdditionalNameInfo(conn *ethclient.Client, currentOwner common.Address, fullName string) (string, string, string, *big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdditionalNameInfo", conn, currentOwner, fullName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(*big.Int)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// GetAdditionalNameInfo indicates an expected call of GetAdditionalNameInfo.
func (mr *MockContractsServiceMockRecorder) GetAdditionalNameInfo(conn, currentOwner, fullName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdditionalNameInfo", reflect.TypeOf((*MockContractsService)(nil).GetAdditionalNameInfo), conn, currentOwner, fullName)
}

// GetBalanceOf mocks base method.
func (m *MockContractsService) GetBalanceOf(client *ethclient.Client, tokenAddress, address common.Address) (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceOf", client, tokenAddress, address)
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceOf indicates an expected call of GetBalanceOf.
func (mr *MockContractsServiceMockRecorder) GetBalanceOf(client, tokenAddress, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceOf", reflect.TypeOf((*MockContractsService)(nil).GetBalanceOf), client, tokenAddress, address)
}

// GetNameByAddress mocks base method.
func (m *MockContractsService) GetNameByAddress(conn *ethclient.Client, address common.Address) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNameByAddress", conn, address)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNameByAddress indicates an expected call of GetNameByAddress.
func (mr *MockContractsServiceMockRecorder) GetNameByAddress(conn, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNameByAddress", reflect.TypeOf((*MockContractsService)(nil).GetNameByAddress), conn, address)
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
func (m *MockContractsService) Register(ctx context.Context, conn *ethclient.Client, authOpts *bind.TransactOpts, nameFirstPart string, registrantAccount common.Address, secret [32]byte, controller *anytype_crypto.AnytypeRegistrarControllerPrivate, fullName, ownerAnyAddr, spaceId string) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, conn, authOpts, nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockContractsServiceMockRecorder) Register(ctx, conn, authOpts, nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockContractsService)(nil).Register), ctx, conn, authOpts, nameFirstPart, registrantAccount, secret, controller, fullName, ownerAnyAddr, spaceId)
}

// RenewName mocks base method.
func (m *MockContractsService) RenewName(ctx context.Context, conn *ethclient.Client, opts *bind.TransactOpts, fullName string, durationSec uint64, controller *anytype_crypto.AnytypeRegistrarControllerPrivate) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewName", ctx, conn, opts, fullName, durationSec, controller)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewName indicates an expected call of RenewName.
func (mr *MockContractsServiceMockRecorder) RenewName(ctx, conn, opts, fullName, durationSec, controller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewName", reflect.TypeOf((*MockContractsService)(nil).RenewName), ctx, conn, opts, fullName, durationSec, controller)
}

// TxByHash mocks base method.
func (m *MockContractsService) TxByHash(ctx context.Context, client *ethclient.Client, txHash common.Hash) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxByHash", ctx, client, txHash)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxByHash indicates an expected call of TxByHash.
func (mr *MockContractsServiceMockRecorder) TxByHash(ctx, client, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxByHash", reflect.TypeOf((*MockContractsService)(nil).TxByHash), ctx, client, txHash)
}

// WaitForTxToStartMining mocks base method.
func (m *MockContractsService) WaitForTxToStartMining(ctx context.Context, conn *ethclient.Client, txHash common.Hash) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForTxToStartMining", ctx, conn, txHash)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForTxToStartMining indicates an expected call of WaitForTxToStartMining.
func (mr *MockContractsServiceMockRecorder) WaitForTxToStartMining(ctx, conn, txHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForTxToStartMining", reflect.TypeOf((*MockContractsService)(nil).WaitForTxToStartMining), ctx, conn, txHash)
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
