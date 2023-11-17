// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.33
// source: proto/anyns_api.proto

package anyns_api

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_proto_anyns_api_proto struct{}

func (drpcEncoding_File_proto_anyns_api_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_proto_anyns_api_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_proto_anyns_api_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_proto_anyns_api_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCAnynsClient interface {
	DRPCConn() drpc.Conn

	IsNameAvailable(ctx context.Context, in *NameAvailableRequest) (*NameAvailableResponse, error)
	GetNameByAddress(ctx context.Context, in *NameByAddressRequest) (*NameByAddressResponse, error)
}

type drpcAnynsClient struct {
	cc drpc.Conn
}

func NewDRPCAnynsClient(cc drpc.Conn) DRPCAnynsClient {
	return &drpcAnynsClient{cc}
}

func (c *drpcAnynsClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcAnynsClient) IsNameAvailable(ctx context.Context, in *NameAvailableRequest) (*NameAvailableResponse, error) {
	out := new(NameAvailableResponse)
	err := c.cc.Invoke(ctx, "/Anyns/IsNameAvailable", drpcEncoding_File_proto_anyns_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcAnynsClient) GetNameByAddress(ctx context.Context, in *NameByAddressRequest) (*NameByAddressResponse, error) {
	out := new(NameByAddressResponse)
	err := c.cc.Invoke(ctx, "/Anyns/GetNameByAddress", drpcEncoding_File_proto_anyns_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCAnynsServer interface {
	IsNameAvailable(context.Context, *NameAvailableRequest) (*NameAvailableResponse, error)
	GetNameByAddress(context.Context, *NameByAddressRequest) (*NameByAddressResponse, error)
}

type DRPCAnynsUnimplementedServer struct{}

func (s *DRPCAnynsUnimplementedServer) IsNameAvailable(context.Context, *NameAvailableRequest) (*NameAvailableResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCAnynsUnimplementedServer) GetNameByAddress(context.Context, *NameByAddressRequest) (*NameByAddressResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCAnynsDescription struct{}

func (DRPCAnynsDescription) NumMethods() int { return 2 }

func (DRPCAnynsDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/Anyns/IsNameAvailable", drpcEncoding_File_proto_anyns_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsServer).
					IsNameAvailable(
						ctx,
						in1.(*NameAvailableRequest),
					)
			}, DRPCAnynsServer.IsNameAvailable, true
	case 1:
		return "/Anyns/GetNameByAddress", drpcEncoding_File_proto_anyns_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCAnynsServer).
					GetNameByAddress(
						ctx,
						in1.(*NameByAddressRequest),
					)
			}, DRPCAnynsServer.GetNameByAddress, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterAnyns(mux drpc.Mux, impl DRPCAnynsServer) error {
	return mux.Register(impl, DRPCAnynsDescription{})
}

type DRPCAnyns_IsNameAvailableStream interface {
	drpc.Stream
	SendAndClose(*NameAvailableResponse) error
}

type drpcAnyns_IsNameAvailableStream struct {
	drpc.Stream
}

func (x *drpcAnyns_IsNameAvailableStream) SendAndClose(m *NameAvailableResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_proto_anyns_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCAnyns_GetNameByAddressStream interface {
	drpc.Stream
	SendAndClose(*NameByAddressResponse) error
}

type drpcAnyns_GetNameByAddressStream struct {
	drpc.Stream
}

func (x *drpcAnyns_GetNameByAddressStream) SendAndClose(m *NameByAddressResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_proto_anyns_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
