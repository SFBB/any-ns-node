package client

import (
	"context"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/net/pool"
	"github.com/anyproto/any-sync/net/rpc/rpcerr"
	"github.com/anyproto/any-sync/nodeconf"

	as "github.com/anyproto/any-ns-node/pb/anyns_api"
)

const CName = "any-ns.anynsclient"

//var log = logger.NewNamed(CName)

/*
 * This client component can be used to access the Any Naming System (any-ns)
 * from other components.
 */
type AnyNsClientService interface {
	IsNameAvailable(ctx context.Context, in *as.NameAvailableRequest) (out *as.NameAvailableResponse, err error)
	GetOperationStatus(ctx context.Context, in *as.GetOperationStatusRequest) (out *as.OperationResponse, err error)
	// TODO: remove and use only NameRegisterSigned
	NameRegister(ctx context.Context, in *as.NameRegisterRequest) (out *as.OperationResponse, err error)
	NameRegisterSigned(ctx context.Context, in *as.NameRegisterSignedRequest) (out *as.OperationResponse, err error)

	NameRenew(ctx context.Context, in *as.NameRenewRequest) (out *as.OperationResponse, err error)

	// reverse resolve
	GetNameByAddress(ctx context.Context, in *as.NameByAddressRequest) (out *as.NameByAddressResponse, err error)

	// AccountAbstractions interface
	GetUserAccount(ctx context.Context, in *as.GetUserAccountRequest) (out *as.UserAccount, err error)
	AdminFundUserAccount(ctx context.Context, in *as.AdminFundUserAccountRequestSigned) (out *as.OperationResponse, err error)
	GetOperation(ctx context.Context, in *as.GetOperationStatusRequest) (out *as.OperationResponse, err error)

	app.ComponentRunnable
}

type service struct {
	pool     pool.Pool
	nodeconf nodeconf.Service
	close    chan struct{}
}

func (s *service) Init(a *app.App) (err error) {
	s.pool = a.MustComponent(pool.CName).(pool.Pool)
	s.nodeconf = a.MustComponent(nodeconf.CName).(nodeconf.Service)
	s.close = make(chan struct{})
	return nil
}

func (s *service) Name() (name string) {
	return CName
}

func New() AnyNsClientService {
	return new(service)
}

func (s *service) Run(_ context.Context) error {
	return nil
}

func (s *service) Close(_ context.Context) error {
	select {
	case <-s.close:
	default:
		close(s.close)
	}
	return nil
}

func (s *service) doClient(ctx context.Context, fn func(cl as.DRPCAnynsClient) error) error {
	// it will try to connect to the Naming Node
	// please enable "namingNode" type of node in the config (in the network.nodes array)
	peer, err := s.pool.Get(ctx, s.nodeconf.NamingNodePeers()[0])

	if err != nil {
		return err
	}

	dc, err := peer.AcquireDrpcConn(ctx)
	if err != nil {
		return err
	}
	defer peer.ReleaseDrpcConn(dc)

	return fn(as.NewDRPCAnynsClient(dc))
}

func (s *service) doClientAA(ctx context.Context, fn func(cl as.DRPCAnynsAccountAbstractionClient) error) error {
	// it will try to connect to the Naming Node
	// please enable "namingNode" type of node in the config (in the network.nodes array)
	peer, err := s.pool.Get(ctx, s.nodeconf.NamingNodePeers()[0])

	if err != nil {
		return err
	}

	dc, err := peer.AcquireDrpcConn(ctx)
	if err != nil {
		return err
	}
	defer peer.ReleaseDrpcConn(dc)

	return fn(as.NewDRPCAnynsAccountAbstractionClient(dc))
}

func (s *service) IsNameAvailable(ctx context.Context, in *as.NameAvailableRequest) (out *as.NameAvailableResponse, err error) {
	err = s.doClient(ctx, func(cl as.DRPCAnynsClient) error {
		if out, err = cl.IsNameAvailable(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) GetOperationStatus(ctx context.Context, in *as.GetOperationStatusRequest) (out *as.OperationResponse, err error) {
	err = s.doClient(ctx, func(cl as.DRPCAnynsClient) error {
		if out, err = cl.GetOperationStatus(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) NameRegister(ctx context.Context, in *as.NameRegisterRequest) (out *as.OperationResponse, err error) {
	err = s.doClient(ctx, func(cl as.DRPCAnynsClient) error {
		if out, err = cl.NameRegister(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) NameRegisterSigned(ctx context.Context, in *as.NameRegisterSignedRequest) (out *as.OperationResponse, err error) {
	err = s.doClient(ctx, func(cl as.DRPCAnynsClient) error {
		if out, err = cl.NameRegisterSigned(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) NameRenew(ctx context.Context, in *as.NameRenewRequest) (out *as.OperationResponse, err error) {
	err = s.doClient(ctx, func(cl as.DRPCAnynsClient) error {
		if out, err = cl.NameRenew(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) GetNameByAddress(ctx context.Context, in *as.NameByAddressRequest) (out *as.NameByAddressResponse, err error) {
	err = s.doClient(ctx, func(cl as.DRPCAnynsClient) error {
		if out, err = cl.GetNameByAddress(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

// AA
func (s *service) GetUserAccount(ctx context.Context, in *as.GetUserAccountRequest) (out *as.UserAccount, err error) {
	err = s.doClientAA(ctx, func(cl as.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.GetUserAccount(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) AdminFundUserAccount(ctx context.Context, in *as.AdminFundUserAccountRequestSigned) (out *as.OperationResponse, err error) {
	err = s.doClientAA(ctx, func(cl as.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.AdminFundUserAccount(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}

func (s *service) GetOperation(ctx context.Context, in *as.GetOperationStatusRequest) (out *as.OperationResponse, err error) {
	err = s.doClientAA(ctx, func(cl as.DRPCAnynsAccountAbstractionClient) error {
		if out, err = cl.GetOperation(ctx, in); err != nil {
			return rpcerr.Unwrap(err)
		}
		return nil
	})
	return
}
