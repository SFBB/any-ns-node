package anynsrpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/anyproto/any-ns-node/cache"
	"github.com/anyproto/any-ns-node/config"
	"github.com/anyproto/any-ns-node/queue"
	"github.com/gogo/protobuf/proto"

	"github.com/anyproto/any-ns-node/verification"
	"github.com/anyproto/any-sync/accountservice"
	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/app/logger"
	"github.com/anyproto/any-sync/net/rpc/server"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"

	contracts "github.com/anyproto/any-ns-node/contracts"
	nsp "github.com/anyproto/any-sync/nameservice/nameserviceproto"
)

const CName = "any-ns.rpc"

var log = logger.NewNamed(CName)

func New() app.Component {
	return &anynsRpc{}
}

type anynsRpc struct {
	cache         cache.CacheService
	confContracts config.Contracts
	confAccount   accountservice.Config
	contracts     contracts.ContractsService
	queue         queue.QueueService

	readFromCache bool
}

func (arpc *anynsRpc) Init(a *app.App) (err error) {
	arpc.cache = a.MustComponent(cache.CName).(cache.CacheService)
	arpc.confContracts = a.MustComponent(config.CName).(*config.Config).GetContracts()
	arpc.confAccount = a.MustComponent(config.CName).(*config.Config).GetAccount()
	arpc.contracts = a.MustComponent(contracts.CName).(contracts.ContractsService)
	arpc.readFromCache = a.MustComponent(config.CName).(*config.Config).ReadFromCache
	arpc.queue = a.MustComponent(queue.CName).(queue.QueueService)

	return nsp.DRPCRegisterAnyns(a.MustComponent(server.CName).(server.DRPCServer), arpc)
}

func (arpc *anynsRpc) Name() (name string) {
	return CName
}

func (arpc *anynsRpc) IsNameAvailable(ctx context.Context, in *nsp.NameAvailableRequest) (*nsp.NameAvailableResponse, error) {
	// 1 - if ReadFromCache is false -> always first read from smart contracts
	// if not, then always just read quickly from cache
	if !arpc.readFromCache {
		log.Debug("EXCPLICIT: read data from smart contracts -> cache", zap.String("FullName", in.FullName))
		err := arpc.cache.UpdateInCache(ctx, &nsp.NameAvailableRequest{
			FullName: in.FullName,
		})

		if err != nil {
			log.Error("failed to update in cache", zap.Error(err))
			return nil, errors.New("failed to update in cache")
		}
	}

	// 2 - check in cache (Mongo)
	return arpc.cache.IsNameAvailable(ctx, in)
}

func (arpc *anynsRpc) GetNameByAddress(ctx context.Context, in *nsp.NameByAddressRequest) (*nsp.NameByAddressResponse, error) {
	// 1 - if ReadFromCache is false -> always first read from smart contracts
	// if not, then always just read quickly from cache
	if !arpc.readFromCache {
		log.Debug("EXCPLICIT: reverse resolve using no cache", zap.String("FullName", in.OwnerScwEthAddress))
		return arpc.getNameByAddressDirectly(ctx, in)
	}

	// check in cache (Mongo)
	return arpc.cache.GetNameByAddress(ctx, in)
}

func (arpc *anynsRpc) getNameByAddressDirectly(ctx context.Context, in *nsp.NameByAddressRequest) (*nsp.NameByAddressResponse, error) {
	// 0 - check parameters
	if !common.IsHexAddress(in.OwnerScwEthAddress) {
		log.Error("invalid ETH address", zap.String("ETH address", in.OwnerScwEthAddress))
		return nil, errors.New("invalid ETH address")
	}

	// 1 - create connection
	conn, err := arpc.contracts.CreateEthConnection()
	if err != nil {
		log.Error("failed to connect to geth", zap.Error(err))
		return nil, errors.New("failed to connect to geth")
	}

	// convert in.OwnerScwEthAddress to common.Address
	var addr = common.HexToAddress(in.OwnerScwEthAddress)

	name, err := arpc.contracts.GetNameByAddress(conn, addr)
	if err != nil {
		log.Error("failed to get name by address", zap.Error(err))
		return nil, errors.New("failed to get name by address")
	}

	// 2 - return results
	var res nsp.NameByAddressResponse

	if name == "" {
		res.Found = false
		return &res, nil
	}

	res.Found = true
	res.Name = name
	return &res, nil
}

func (arpc *anynsRpc) AdminNameRegisterSigned(ctx context.Context, in *nsp.NameRegisterRequestSigned) (*nsp.OperationResponse, error) {
	var resp nsp.OperationResponse

	// 1 - unmarshal the signed request
	var nrr nsp.NameRegisterRequest
	err := proto.Unmarshal(in.Payload, &nrr)
	if err != nil {
		resp.OperationState = nsp.OperationState_Error
		log.Error("can not unmarshal NameRegisterRequest", zap.Error(err))
		return &resp, err
	}

	// 2 - check signature
	err = verification.VerifyAdminIdentity(arpc.confAccount.SigningKey, in.Payload, in.Signature)
	if err != nil {
		resp.OperationState = nsp.OperationState_Error
		log.Error("identity is different", zap.Error(err))
		return &resp, err
	}

	// 3 - check all parameters
	err = verification.CheckRegisterParams(&nrr)
	if err != nil {
		log.Error("invalid parameters", zap.Error(err))
		return nil, err
	}

	// 4 - add to queue
	operationId, err := arpc.queue.AddNewRequest(ctx, &nrr)
	resp.OperationId = fmt.Sprint(operationId)
	resp.OperationState = nsp.OperationState_Pending
	return &resp, err
}
