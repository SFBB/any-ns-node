package accountabstraction

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"testing"

	"github.com/anyproto/any-sync/app"
	"github.com/anyproto/any-sync/commonspace/object/accountdata"
	"github.com/anyproto/any-sync/net/rpc/rpctest"
	"github.com/anyproto/any-sync/util/crypto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
	"go.uber.org/mock/gomock"

	"github.com/anyproto/any-ns-node/alchemysdk"
	mock_alchemysdk "github.com/anyproto/any-ns-node/alchemysdk/mock"
	"github.com/anyproto/any-ns-node/config"
	"github.com/anyproto/any-ns-node/contracts"
	mock_contracts "github.com/anyproto/any-ns-node/contracts/mock"

	nsp "github.com/anyproto/any-sync/nameservice/nameserviceproto"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var ctx = context.Background()

type fixture struct {
	a         *app.App
	ctrl      *gomock.Controller
	ts        *rpctest.TestServer
	config    *config.Config
	contracts *mock_contracts.MockContractsService
	alchemy   *mock_alchemysdk.MockAlchemyAAService

	*anynsAA
}

func newFixture(t *testing.T) *fixture {
	fx := &fixture{
		a:       new(app.App),
		ctrl:    gomock.NewController(t),
		ts:      rpctest.NewTestServer(),
		config:  new(config.Config),
		anynsAA: New().(*anynsAA),
	}

	fx.config.Contracts = config.Contracts{
		AddrAdmin:     "0x10d5B0e279E5E4c1d1Df5F57DFB7E84813920a51",
		GethUrl:       "https://sepolia.infura.io/v3/68c55936b8534264801fa4bc313ff26f",
		TokenDecimals: 6,
	}
	fx.config.Account.PeerId = "12D3KooWA8EXV3KjBxEU5EnsPfneLx84vMWAtTBQBeyooN82KSuS"
	fx.config.Account.PeerKey = "psqF8Rj52Ci6gsUl5ttwBVhINTP8Yowc2hea73MeFm4Ek9AxedYSB4+r7DYCclDL4WmLggj2caNapFUmsMtn5Q=="
	fx.config.Account.SigningKey = "3MFdA66xRw9PbCWlfa620980P4QccXehFlABnyJ/tfwHbtBVHt+KWuXOfyWSF63Ngi70m+gcWtPAcW5fxCwgVg=="

	fx.config.Aa.NameTokensPerName = 10

	fx.contracts = mock_contracts.NewMockContractsService(fx.ctrl)
	fx.contracts.EXPECT().Name().Return(contracts.CName).AnyTimes()
	fx.contracts.EXPECT().Init(gomock.Any()).AnyTimes()
	fx.contracts.EXPECT().CreateEthConnection().AnyTimes()
	fx.contracts.EXPECT().GenerateAuthOptsForAdmin(gomock.Any()).MaxTimes(2)
	fx.contracts.EXPECT().CalculateTxParams(gomock.Any(), gomock.Any()).AnyTimes()
	fx.contracts.EXPECT().ConnectToController(gomock.Any()).AnyTimes()
	fx.contracts.EXPECT().TxByHash(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	fx.contracts.EXPECT().MakeCommitment(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	fx.contracts.EXPECT().WaitForTxToStartMining(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	fx.contracts.EXPECT().IsContractDeployed(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

	fx.alchemy = mock_alchemysdk.NewMockAlchemyAAService(fx.ctrl)
	fx.alchemy.EXPECT().Name().Return(alchemysdk.CName).AnyTimes()
	fx.alchemy.EXPECT().Init(gomock.Any()).AnyTimes()
	//fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	//fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).AnyTimes()

	fx.a.Register(fx.ts).
		Register(fx.config).
		Register(fx.contracts).
		Register(fx.alchemy).
		Register(fx.anynsAA)

	require.NoError(t, fx.a.Start(ctx))
	return fx
}

func (fx *fixture) finish(t *testing.T) {
	assert.NoError(t, fx.a.Close(ctx))
	fx.ctrl.Finish()
}

func TestAAS_GetSmartWalletAddress(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail if can not connect to smart contract", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		const factoryAddress = "0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789"
		factoryAddressBytes := common.HexToAddress(factoryAddress).Bytes()
		assert.Equal(t, factoryAddressBytes[0], byte(0x5F))

		fa := common.BytesToAddress(factoryAddressBytes)
		assert.Equal(t, factoryAddress, fa.String())

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			return nil, errors.New("fail")
		})

		pctx := context.Background()
		_, err := fx.GetSmartWalletAddress(pctx, common.HexToAddress("0xE34230c1f916e9d628D5F9863Eb3F5667D8FcB37"))
		assert.Error(t, err)
	})

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		})

		pctx := context.Background()
		swa, err := fx.GetSmartWalletAddress(pctx, common.HexToAddress("0xE34230c1f916e9d628D5F9863Eb3F5667D8FcB37"))

		assert.NoError(t, err)
		assert.Equal(t, common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a"), swa)
	})
}

func TestAAS_GetNonceForSmartWalletAddress(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail if can not connect to smart contract", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			return nil, errors.New("fail")
		})

		pctx := context.Background()
		_, err := fx.getNonceForSmartWalletAddress(pctx, common.HexToAddress("0xE34230c1f916e9d628D5F9863Eb3F5667D8FcB37"))
		assert.Error(t, err)
	})

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			out := big.NewInt(6)

			byteArr := out.Bytes()
			return byteArr, nil
		})

		pctx := context.Background()
		nonce, err := fx.getNonceForSmartWalletAddress(pctx, common.HexToAddress("0xE34230c1f916e9d628D5F9863Eb3F5667D8FcB37"))

		assert.NoError(t, err)
		six := big.NewInt(6)
		assert.Equal(t, nonce, six)
	})
}

func TestAAS_GetNamesCountLeft(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail if can not get token balance", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().GetBalanceOf(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, tokenAddress interface{}, scw interface{}, x interface{}) (*big.Int, error) {
			return big.NewInt(0), errors.New("failed to get balance")
		})

		count, err := fx.GetNamesCountLeft(ctx, common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a"))

		assert.Error(t, err)
		assert.Equal(t, uint64(0), count)
	})

	mt.Run("success if no tokens", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().GetBalanceOf(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, tokenAddress interface{}, scw interface{}, x interface{}) (*big.Int, error) {
			return big.NewInt(0), nil
		})

		count, err := fx.GetNamesCountLeft(ctx, common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a"))

		assert.NoError(t, err)
		assert.Equal(t, uint64(0), count)
	})

	mt.Run("success if not enough tokens", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().GetBalanceOf(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, tokenAddress interface{}, scw interface{}, x interface{}) (*big.Int, error) {
			// 10 tokens per name (current testnet settings)
			// 6 decimals
			oneNamePriceWei := big.NewInt(10 * 1000000)

			// divide oneNamePriceWei /2 to get less than 1 name
			out := big.NewInt(0).Div(oneNamePriceWei, big.NewInt(2))
			return out, nil
		})

		count, err := fx.GetNamesCountLeft(ctx, common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a"))

		assert.NoError(t, err)
		assert.Equal(t, uint64(0), count)
	})

	mt.Run("success if got N tokens", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().GetBalanceOf(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, tokenAddress interface{}, scw interface{}, x interface{}) (*big.Int, error) {
			oneNamePriceWei := big.NewInt(10 * 1000000)

			// multiply by 12
			out := big.NewInt(0).Mul(oneNamePriceWei, big.NewInt(12))
			return out, nil
		})

		count, err := fx.GetNamesCountLeft(ctx, common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a"))

		assert.NoError(t, err)
		assert.Equal(t, uint64(12), count)
	})
}
func TestAAS_VerifyAdminIdentity(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)
		// 0 - garbage data test
		err := fx.AdminVerifyIdentity([]byte("payload"), []byte("signature"))
		assert.Error(t, err)

		// 1 - pack some structure
		nrr := nsp.AdminFundUserAccountRequest{
			OwnerEthAddress: "",
			NamesCount:      0,
		}

		marshalled, err := nrr.Marshal()
		require.NoError(t, err)

		// 2 - sign it with some random (wrong) key
		accountKeys, err := accountdata.NewRandom()
		require.NoError(t, err)

		sig, err := accountKeys.SignKey.Sign(marshalled)
		require.NoError(t, err)

		err = fx.AdminVerifyIdentity(marshalled, sig)
		assert.Error(t, err)
	})

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		// 1 - pack some structure
		nrr := nsp.AdminFundUserAccountRequest{
			OwnerEthAddress: "",
			NamesCount:      0,
		}

		marshalled, err := nrr.Marshal()
		require.NoError(t, err)

		// 2 - sign it
		signKey, err := crypto.DecodeKeyFromString(
			fx.config.Account.PeerKey,
			crypto.UnmarshalEd25519PrivateKey,
			nil)
		require.NoError(t, err)

		sig, err := signKey.Sign(marshalled)
		require.NoError(t, err)

		// get associated pub key
		//pubKey := signKey.GetPublic()
		// identity str
		//identityStr := pubKey.Account()
		// A5ommzwhpR5ngp11q9q1P2MMzhUE46Hi421RJbPqswALyoyr
		//log.Info("identity", zap.String("identity", identityStr))

		err = fx.AdminVerifyIdentity(marshalled, sig)
		assert.NoError(t, err)
	})
}

func TestAAS_MintAccessTokens(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail if names count is ZERO", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		// already deployed
		scw := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a")
		_, err := fx.AdminMintAccessTokens(ctx, scw, big.NewInt(0))
		assert.Error(t, err)
	})

	mt.Run("success if SCW was already deployed", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		// nonce is 5
		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			out := big.NewInt(5)

			byteArr := out.Bytes()
			return byteArr, nil
		}).AnyTimes()

		fx.contracts.EXPECT().IsContractDeployed(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, client interface{}, scw interface{}) (bool, error) {
			// deployed!!!
			return true, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		// already deployed
		scw := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a")
		_, err := fx.AdminMintAccessTokens(ctx, scw, big.NewInt(5))
		assert.NoError(t, err)
	})

	mt.Run("success even if SCW is not deployed", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		// nonce is 5
		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			out := big.NewInt(5)

			byteArr := out.Bytes()
			return byteArr, nil
		}).AnyTimes()

		fx.contracts.EXPECT().IsContractDeployed(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, client interface{}, scw interface{}) (bool, error) {
			// not deployed!
			return false, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		// already deployed
		scw := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a")
		_, err := fx.AdminMintAccessTokens(ctx, scw, big.NewInt(5))
		assert.NoError(t, err)
	})
}

func TestAAS_GetDataNameRegister(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail if no fullname specified", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("fail if no any address specified", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("fail if cannot CreateRequestGasAndPaymasterData", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}) (out []byte, err error) {
			return []byte{}, errors.New("fail")
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("fail if cannot SendRequest", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}) (out []byte, err error) {
			return []byte{}, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			return nil, errors.New("fail")
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("fail if SendRequest return wrong JSON", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		// return wrong JSON
		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			byteArr := []byte("123A")

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("fail if SendRequest return error code", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// set error
			response.Error.Code = 123
			response.Error.Message = "Something really bad happened, sorry"

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("fail if CreateRequestStep1 failed", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, errors.New("fail")
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		_, _, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.Error(t, err)
	})

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.contracts.EXPECT().CallContract(gomock.Any(), gomock.Any()).DoAndReturn(func(tokenAddress interface{}, scw interface{}) ([]byte, error) {
			byteArr := common.HexToAddress("0x77d454b313e9D1Acb8cD0cFa140A27544aEC483a").Bytes()
			return byteArr, nil
		}).AnyTimes()

		var req nsp.NameRegisterRequest = nsp.NameRegisterRequest{
			FullName:        "hello.any",
			OwnerEthAddress: "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF",
			OwnerAnyAddress: "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy",
			SpaceId:         "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu",
		}

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep1(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}) (out []byte, uo alchemysdk.UserOperation, err error) {
			var uoOut alchemysdk.UserOperation

			return []byte{}, uoOut, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGasAndPaymasterData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		dataToSign, contextData, err := fx.GetDataNameRegister(context.Background(), &req)
		assert.NoError(t, err)
		assert.NotNil(t, contextData)
		assert.NotNil(t, dataToSign)
	})
}

func TestAAS_GetCallDataForMint(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		smartAccountAddress := common.HexToAddress("0x045F756F248799F4413a026100Ae49e5E7F2031E")

		out, err := GetCallDataForMint(smartAccountAddress, big.NewInt(1), 6)
		outStr := "0x" + hex.EncodeToString(out)

		assert.NoError(t, err)
		assert.Equal(t, outStr, "0x40c10f19000000000000000000000000045f756f248799f4413a026100ae49e5e7f2031e00000000000000000000000000000000000000000000000000000000000f4240")
	})
}

func TestAAS_GetCallDataForAprove(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		from := common.HexToAddress("0x045F756F248799F4413a026100Ae49e5E7F2031E")
		registrarController := common.HexToAddress("0xB6bF17cBe45CbC7609e4f8fA56154c9DeF8590CA")

		out, err := GetCallDataForAprove(from, registrarController, big.NewInt(1), 6)
		outStr := "0x" + hex.EncodeToString(out)

		assert.NoError(t, err)
		assert.Equal(t, outStr, "0x2b991746000000000000000000000000045f756f248799f4413a026100ae49e5e7f2031e000000000000000000000000b6bf17cbe45cbc7609e4f8fa56154c9def8590ca00000000000000000000000000000000000000000000000000000000000f4240")
	})
}

func TestAAS_GetCallDataForBatchExecute(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		erc20tokenAddr := common.HexToAddress("0x8AE88b2b35F15D6320D77ab8EC7E3410F78376F6")

		// just some random data
		data1 := "0x40c10f19000000000000000000000000045f756f248799f4413a026100ae49e5e7f2031e0000000000000000000000000000000000000000000000000000000000000064"
		callDataOriginal1, err := hex.DecodeString(data1[2:])
		assert.NoError(t, err)

		// just some random data
		data2 := "0x40c10f19000000000000000000000000045f756f248799f4413a026100ae49e5e7f2031e0000000000000000000000000000000000000000000000000000000000000064"
		callDataOriginal2, err := hex.DecodeString(data2[2:])
		assert.NoError(t, err)

		// put address and address2 into array
		// both are the same
		addresses := []common.Address{erc20tokenAddr, erc20tokenAddr}
		// put data1 and callDataOriginal2 into array
		datas := [][]byte{callDataOriginal1, callDataOriginal2}

		out, err := GetCallDataForBatchExecute(addresses, datas)
		outStr := "0x" + hex.EncodeToString(out)

		assert.NoError(t, err)
		assert.Equal(t, outStr, "0x18dfb3c7000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000020000000000000000000000008ae88b2b35f15d6320d77ab8ec7e3410f78376f60000000000000000000000008ae88b2b35f15d6320d77ab8ec7e3410f78376f60000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000000004440c10f19000000000000000000000000045f756f248799f4413a026100ae49e5e7f2031e000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004440c10f19000000000000000000000000045f756f248799f4413a026100ae49e5e7f2031e000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000")
	})
}

func TestAAS_GetCallDataForCommit(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		// just some random data
		commitmentStr := "0x1234"
		commitment, err := hex.DecodeString(commitmentStr[2:])
		assert.NoError(t, err)

		// convert commitment to [32]byte
		var commitment32 [32]byte
		copy(commitment32[:], commitment)

		data, err := GetCallDataForCommit(commitment32)
		assert.NoError(t, err)
		assert.Equal(t, "0x"+hex.EncodeToString(data), "0xf14fcbc81234000000000000000000000000000000000000000000000000000000000000")
	})
}

func TestAAS_GetCallDataForRegister(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		secret, err := hex.DecodeString("a4f49c1a7b979dc0ea76cd083a97af07e5983e7041f84bc672134e5b24f21218")
		assert.NoError(t, err)

		// convert secret to [32]byte
		var secret32 [32]byte
		copy(secret32[:], secret)

		ownerAnyAddress := "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy"
		spaceID := "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu"
		callData, _ := contracts.PrepareCallData_SetContentHashSpaceID("xxx123.any", ownerAnyAddress, spaceID)

		var nameFirstPart string = "xxx123"
		var registrantAccount common.Address = common.HexToAddress("0xE34230c1f916e9d628D5F9863Eb3F5667D8FcB37")
		var registrationTime big.Int = *big.NewInt(12324)
		var resolver common.Address = common.HexToAddress("0x8AE88b2b35F15D6320D77ab8EC7E3410F78376F6")
		var isReverseRecord bool = false
		var ownerControlledFuses uint16 = 0

		data, err := GetCallDataForRegister(nameFirstPart, registrantAccount, registrationTime, secret32, resolver, callData, isReverseRecord, ownerControlledFuses)
		assert.NoError(t, err)
		assert.Equal(t, "0x"+hex.EncodeToString(data), "0x74694a2b0000000000000000000000000000000000000000000000000000000000000100000000000000000000000000e34230c1f916e9d628d5f9863eb3f5667d8fcb370000000000000000000000000000000000000000000000000000000000003024a4f49c1a7b979dc0ea76cd083a97af07e5983e7041f84bc672134e5b24f212180000000000000000000000008ae88b2b35f15d6320d77ab8ec7e3410f78376f60000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000006787878313233000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000a4f49c1a7b979dc0ea76cd083a97af07e5983e7041f84bc672134e5b24f212181bf35d6d1b0000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000003b626166796265696273363267717469676e75636b66716c6372376c68686968677a6832766f7278746d633561666d36757868347a64636d7577757500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a4304e6ade979dc0ea76cd083a97af07e5983e7041f84bc672134e5b24f212181bf35d6d1b00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000034313244334b6f6f5750414e7a565a674871414c353743636852483471384e476a6f574470555368566f7642453362686858637a7900000000000000000000000000000000000000000000000000000000000000000000000000000000")
	})
}

func TestAAS_GetCallDataForNameRegister(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fullName := "hello.any"
		ownerAnyAddress := "12D3KooWPANzVZgHqAL57CchRH4q8NGjoWDpUShVovBE3bhhXczy"
		ownerEthAddress := "0xe595e2BA3f0cE990d8037e07250c5C78ce40f8fF"
		spaceID := "bafybeibs62gqtignuckfqlcr7lhhihgzh2vorxtmc5afm6uxh4zdcmuwuu"

		_, err := fx.getCallDataForNameRegister(fullName, ownerAnyAddress, ownerEthAddress, spaceID)
		assert.NoError(t, err)

		// the result has some randomness in it (secret)
		//outStr := "0x" + hex.EncodeToString(data)
		//assert.Equal(t, outStr, "0x6a3b8f2a000")
	})
}

func TestAAS_SendUserOperation(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("fail if context is not a valid JSON", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep2(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		contextData := []byte("123A")
		signedData, _ := hex.DecodeString("12AF")

		_, err := fx.SendUserOperation(ctx, contextData, signedData)
		assert.Error(t, err)
	})

	mt.Run("fail if CreateRequestStep2 failed", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep2(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, scw interface{}, nonce interface{}) (out []byte, err error) {
			return []byte{}, errors.New("fail")
		}).AnyTimes()

		// contextData is a marshalled UserOperation
		var uo alchemysdk.UserOperation
		contextData, err := json.Marshal(uo)
		assert.NoError(t, err)

		signedData, _ := hex.DecodeString("12AF")

		_, err = fx.SendUserOperation(ctx, contextData, signedData)
		assert.Error(t, err)
	})

	mt.Run("fail if SendRequest returns error", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			return nil, errors.New("i cannot")
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep2(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		// contextData is a marshalled UserOperation
		var uo alchemysdk.UserOperation
		contextData, err := json.Marshal(uo)
		assert.NoError(t, err)

		signedData, _ := hex.DecodeString("12AF")

		_, err = fx.SendUserOperation(ctx, contextData, signedData)
		assert.Error(t, err)
	})

	mt.Run("do not fail even if CreateRequestGetUserOperationReceipt fails", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			return nil, errors.New("i cannot")
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep2(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		// contextData is a marshalled UserOperation
		var uo alchemysdk.UserOperation
		contextData, err := json.Marshal(uo)
		assert.NoError(t, err)

		signedData, _ := hex.DecodeString("12AF")

		_, err = fx.SendUserOperation(ctx, contextData, signedData)
		assert.NoError(t, err)
	})

	mt.Run("success", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().DecodeResponseSendRequest(gomock.Any()).DoAndReturn(func(one interface{}) (opHash string, err error) {
			return "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186", nil
		}).AnyTimes()

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			// convert alchemysdk.JSONRPCResponseGasAndPaymaster to []byte array
			response := alchemysdk.JSONRPCResponseGasAndPaymaster{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestAndSign(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}, s interface{}, scw interface{}, nonce interface{}, gasPrice interface{}, x interface{}, y interface{}, z interface{}, xx interface{}, yy interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequest

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestStep2(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()

		// contextData is a marshalled UserOperation
		var uo alchemysdk.UserOperation
		contextData, err := json.Marshal(uo)
		assert.NoError(t, err)

		signedData, _ := hex.DecodeString("12AF")

		op, err := fx.SendUserOperation(ctx, contextData, signedData)
		assert.NoError(t, err)
		assert.Equal(t, op, "0x31b09cc37a91866b493ee9a31980e90b94b09195a85599f5e6d6a246c9e20186")
	})
}

func TestAAS_GetOperationInfo(t *testing.T) {
	var mt = mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("should return NOT-FOUND if error", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			response := alchemysdk.JSONRPCResponseUserOpHash{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		// error here
		fx.alchemy.EXPECT().DecodeResponseGetUserOperationReceipt(gomock.Any()).DoAndReturn(func(one interface{}) (ret *alchemysdk.JSONRPCResponseGetOp, err error) {
			//var out alchemysdk.JSONRPCResponseGetOp
			return nil, errors.New("bad error")
		}).AnyTimes()

		op, err := fx.GetOperation(ctx, "123")
		assert.NoError(t, err)
		assert.Equal(t, op.OperationState, nsp.OperationState_Error)
	})

	mt.Run("should return NOT-FOUND if error field is set", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			response := alchemysdk.JSONRPCResponseUserOpHash{}
			response.Error.Code = 123
			response.Error.Message = "bad error"

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().DecodeResponseGetUserOperationReceipt(gomock.Any()).DoAndReturn(func(one interface{}) (ret *alchemysdk.JSONRPCResponseGetOp, err error) {
			var out alchemysdk.JSONRPCResponseGetOp
			out.Error.Code = 123
			out.Error.Message = "bad error"

			return &out, nil
		}).AnyTimes()

		op, err := fx.GetOperation(ctx, "123")
		assert.NoError(t, err)
		assert.Equal(t, op.OperationState, nsp.OperationState_Error)
	})

	mt.Run("should return PENDING if UserOpHash field is null", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			response := alchemysdk.JSONRPCResponseUserOpHash{}

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().DecodeResponseGetUserOperationReceipt(gomock.Any()).DoAndReturn(func(one interface{}) (ret *alchemysdk.JSONRPCResponseGetOp, err error) {
			var out alchemysdk.JSONRPCResponseGetOp
			out.Result.UserOpHash = "" // here
			out.Error.Code = 0
			out.Error.Message = ""

			return &out, nil
		}).AnyTimes()

		op, err := fx.GetOperation(ctx, "123")
		assert.NoError(t, err)
		assert.Equal(t, op.OperationState, nsp.OperationState_PendingOrNotFound)
	})

	mt.Run("should return error if Success field is false", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			response := alchemysdk.JSONRPCResponseUserOpHash{}
			response.Error.Code = 0
			response.Error.Message = ""

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().DecodeResponseGetUserOperationReceipt(gomock.Any()).DoAndReturn(func(one interface{}) (ret *alchemysdk.JSONRPCResponseGetOp, err error) {
			var out alchemysdk.JSONRPCResponseGetOp
			out.Result.UserOpHash = "123"
			out.Result.Success = false // here
			out.Error.Code = 0
			out.Error.Message = ""

			return &out, nil
		}).AnyTimes()

		op, err := fx.GetOperation(ctx, "123")
		assert.NoError(t, err)
		assert.Equal(t, op.OperationState, nsp.OperationState_Error)
	})

	mt.Run("success if receipt has Success==true", func(mt *mtest.T) {
		fx := newFixture(t)
		defer fx.finish(t)

		fx.alchemy.EXPECT().SendRequest(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			response := alchemysdk.JSONRPCResponseUserOpHash{}
			response.Error.Code = 0
			response.Error.Message = ""

			// convert to JSON
			jsonDATA, err := json.Marshal(response)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().CreateRequestGetUserOperationReceipt(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx interface{}, in interface{}) (out []byte, err error) {
			var req alchemysdk.JSONRPCRequestGetUserOperationReceipt

			// convert to JSON
			jsonDATA, err := json.Marshal(req)
			assert.NoError(t, err)

			// convert to []byte
			byteArr := []byte(jsonDATA)

			return byteArr, nil
		}).AnyTimes()

		fx.alchemy.EXPECT().DecodeResponseGetUserOperationReceipt(gomock.Any()).DoAndReturn(func(one interface{}) (ret *alchemysdk.JSONRPCResponseGetOp, err error) {
			var out alchemysdk.JSONRPCResponseGetOp
			out.Result.UserOpHash = "123" // here
			out.Result.Success = true     // here
			out.Error.Code = 0
			out.Error.Message = ""

			return &out, nil
		}).AnyTimes()

		op, err := fx.GetOperation(ctx, "123")
		assert.NoError(t, err)
		assert.Equal(t, op.OperationState, nsp.OperationState_Completed)
	})
}
