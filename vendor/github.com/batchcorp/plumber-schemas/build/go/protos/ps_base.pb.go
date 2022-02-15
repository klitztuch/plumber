// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_base.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("ps_base.proto", fileDescriptor_ef931efdbd582aee) }

var fileDescriptor_ef931efdbd582aee = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xef, 0x44, 0xc7, 0xed, 0xba, 0xce, 0x8d, 0x9a, 0xad, 0xee, 0xfa, 0x00, 0x36, 0xa0,
	0xe2, 0x9d, 0xe0, 0x5a, 0x31, 0x08, 0xc2, 0x4a, 0xb7, 0x05, 0xf1, 0x26, 0x24, 0xe9, 0xa1, 0x2d,
	0x4c, 0x32, 0x63, 0xce, 0x44, 0xf0, 0x31, 0x7d, 0x23, 0x69, 0xce, 0xa4, 0xf3, 0x2f, 0xd9, 0xab,
	0x85, 0xef, 0x77, 0xf2, 0xe3, 0xdb, 0x9c, 0xe9, 0x84, 0xcd, 0x14, 0xe6, 0x65, 0x81, 0xb0, 0x50,
	0xad, 0xd4, 0x92, 0x3f, 0xe8, 0xff, 0x60, 0x72, 0xa1, 0x30, 0xaf, 0x64, 0xd3, 0x40, 0xa5, 0x89,
	0x24, 0x4f, 0x14, 0xe6, 0xba, 0x6b, 0x1a, 0x10, 0x26, 0x38, 0x57, 0x98, 0xb7, 0x20, 0x8a, 0xbf,
	0xce, 0x00, 0x42, 0xfb, 0x07, 0x5a, 0x0a, 0xde, 0xfe, 0x3b, 0x63, 0xb3, 0x1f, 0xa2, 0xab, 0x4b,
	0x68, 0xef, 0xfa, 0x9c, 0xff, 0x64, 0x4f, 0x33, 0xd0, 0x37, 0x42, 0x2c, 0x49, 0x7d, 0x90, 0x0d,
	0xf2, 0x6b, 0x1a, 0xc7, 0x45, 0x84, 0x56, 0xf0, 0xbb, 0x03, 0xd4, 0xc9, 0xeb, 0x7b, 0x26, 0x50,
	0xc9, 0x06, 0x81, 0x7f, 0x67, 0xb3, 0x0c, 0xb4, 0x25, 0x7c, 0xee, 0x3c, 0x63, 0xe3, 0xc1, 0xf8,
	0x72, 0x82, 0x1a, 0xdb, 0x86, 0x5d, 0x2c, 0x5b, 0x28, 0x34, 0x38, 0xc2, 0xab, 0xe1, 0x91, 0x90,
	0x0c, 0xce, 0xeb, 0xe9, 0x01, 0xa3, 0xbd, 0x65, 0xe7, 0x6b, 0x40, 0xb7, 0xe5, 0xa9, 0x87, 0x9f,
	0x0f, 0xca, 0x57, 0x53, 0xd8, 0xf6, 0xdc, 0xa8, 0xed, 0x44, 0xcf, 0x90, 0x44, 0x3d, 0xe3, 0x01,
	0xab, 0xfd, 0x02, 0x02, 0xc6, 0xb5, 0x21, 0x89, 0xb4, 0xf1, 0x80, 0xd1, 0x7e, 0x63, 0x67, 0xb4,
	0xc0, 0xd5, 0xf1, 0xd4, 0x20, 0xbf, 0xf4, 0xd7, 0x4a, 0xe9, 0xa0, 0x9b, 0x8f, 0x43, 0xa3, 0xfa,
	0xc8, 0x1e, 0x66, 0xa0, 0xfb, 0x90, 0x3f, 0x73, 0x26, 0xfb, 0x64, 0x50, 0x3c, 0x8f, 0x81, 0x79,
	0xfc, 0x2b, 0x7b, 0x4c, 0x4b, 0x22, 0x43, 0xe2, 0x6f, 0xce, 0x93, 0x5c, 0x8e, 0x32, 0xeb, 0xa1,
	0x97, 0x18, 0x78, 0x9c, 0x30, 0xf2, 0x78, 0xcc, 0x7a, 0x56, 0x80, 0x5d, 0x1d, 0x7a, 0x9c, 0x30,
	0xf2, 0x78, 0xcc, 0x78, 0x3e, 0xb1, 0x47, 0x77, 0x5a, 0x2a, 0xb2, 0x9c, 0xfe, 0xfd, 0x53, 0x34,
	0x38, 0x5e, 0x8c, 0x10, 0xdb, 0x84, 0xf6, 0x17, 0x34, 0x71, 0xc2, 0xa8, 0x89, 0xc7, 0x6c, 0x93,
	0x0c, 0xf4, 0xba, 0xbf, 0x2f, 0xb8, 0xbb, 0x08, 0x8a, 0xa2, 0x26, 0x0e, 0xf1, 0x7e, 0xd1, 0x37,
	0x42, 0x50, 0x8e, 0x3c, 0x38, 0x11, 0x26, 0x1e, 0xfb, 0x45, 0xbb, 0xd4, 0x9e, 0x3d, 0x5a, 0xa0,
	0xa9, 0x14, 0xac, 0xd5, 0x6f, 0x35, 0x1f, 0x87, 0x46, 0xb5, 0x64, 0xec, 0xf8, 0xde, 0x8c, 0xc8,
	0x7b, 0x97, 0xbe, 0x26, 0x19, 0x43, 0xb6, 0x0f, 0x2d, 0x30, 0xec, 0xe3, 0xa6, 0x51, 0x1f, 0x1f,
	0x5a, 0x15, 0x9d, 0xa9, 0x50, 0xe5, 0xa6, 0x91, 0xca, 0x87, 0x56, 0x45, 0xcb, 0x0c, 0x55, 0x6e,
	0x1a, 0xa9, 0x7c, 0x68, 0xef, 0x90, 0x0c, 0x34, 0xdd, 0xfb, 0xb7, 0x8a, 0x6e, 0xfa, 0x2b, 0x67,
	0x47, 0x1e, 0x89, 0xee, 0x90, 0x78, 0x80, 0xb4, 0x9f, 0x3f, 0xfc, 0x7a, 0xbf, 0x3b, 0xe8, 0x7d,
	0x57, 0x2e, 0x2a, 0x59, 0xa7, 0x65, 0xa1, 0xab, 0x7d, 0x25, 0x5b, 0x95, 0x2a, 0xfa, 0xce, 0xbc,
	0xc1, 0x6a, 0x0f, 0x75, 0x81, 0x69, 0xd9, 0x1d, 0xc4, 0x36, 0xdd, 0xc9, 0x94, 0x84, 0x25, 0x7d,
	0xd7, 0xde, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x3c, 0xc4, 0x54, 0xef, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PlumberServerClient is the client API for PlumberServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlumberServerClient interface {
	// List configured/known connections
	GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error)
	GetAllRelays(ctx context.Context, in *GetAllRelaysRequest, opts ...grpc.CallOption) (*GetAllRelaysResponse, error)
	GetRelay(ctx context.Context, in *GetRelayRequest, opts ...grpc.CallOption) (*GetRelayResponse, error)
	CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error)
	UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error)
	ResumeRelay(ctx context.Context, in *ResumeRelayRequest, opts ...grpc.CallOption) (*ResumeRelayResponse, error)
	StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error)
	DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error)
	GetTunnel(ctx context.Context, in *GetTunnelRequest, opts ...grpc.CallOption) (*GetTunnelResponse, error)
	GetAllTunnels(ctx context.Context, in *GetAllTunnelsRequest, opts ...grpc.CallOption) (*GetAllTunnelsResponse, error)
	CreateTunnel(ctx context.Context, in *CreateTunnelRequest, opts ...grpc.CallOption) (*CreateTunnelResponse, error)
	StopTunnel(ctx context.Context, in *StopTunnelRequest, opts ...grpc.CallOption) (*StopTunnelResponse, error)
	ResumeTunnel(ctx context.Context, in *ResumeTunnelRequest, opts ...grpc.CallOption) (*ResumeTunnelResponse, error)
	UpdateTunnel(ctx context.Context, in *UpdateTunnelRequest, opts ...grpc.CallOption) (*UpdateTunnelResponse, error)
	DeleteTunnel(ctx context.Context, in *DeleteTunnelRequest, opts ...grpc.CallOption) (*DeleteTunnelResponse, error)
	GetServerOptions(ctx context.Context, in *GetServerOptionsRequest, opts ...grpc.CallOption) (*GetServerOptionsResponse, error)
}

type plumberServerClient struct {
	cc *grpc.ClientConn
}

func NewPlumberServerClient(cc *grpc.ClientConn) PlumberServerClient {
	return &plumberServerClient{cc}
}

func (c *plumberServerClient) GetAllConnections(ctx context.Context, in *GetAllConnectionsRequest, opts ...grpc.CallOption) (*GetAllConnectionsResponse, error) {
	out := new(GetAllConnectionsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetConnection(ctx context.Context, in *GetConnectionRequest, opts ...grpc.CallOption) (*GetConnectionResponse, error) {
	out := new(GetConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateConnection(ctx context.Context, in *CreateConnectionRequest, opts ...grpc.CallOption) (*CreateConnectionResponse, error) {
	out := new(CreateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) TestConnection(ctx context.Context, in *TestConnectionRequest, opts ...grpc.CallOption) (*TestConnectionResponse, error) {
	out := new(TestConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/TestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateConnection(ctx context.Context, in *UpdateConnectionRequest, opts ...grpc.CallOption) (*UpdateConnectionResponse, error) {
	out := new(UpdateConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error) {
	out := new(DeleteConnectionResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetAllRelays(ctx context.Context, in *GetAllRelaysRequest, opts ...grpc.CallOption) (*GetAllRelaysResponse, error) {
	out := new(GetAllRelaysResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllRelays", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetRelay(ctx context.Context, in *GetRelayRequest, opts ...grpc.CallOption) (*GetRelayResponse, error) {
	out := new(GetRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateRelay(ctx context.Context, in *CreateRelayRequest, opts ...grpc.CallOption) (*CreateRelayResponse, error) {
	out := new(CreateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateRelay(ctx context.Context, in *UpdateRelayRequest, opts ...grpc.CallOption) (*UpdateRelayResponse, error) {
	out := new(UpdateRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) ResumeRelay(ctx context.Context, in *ResumeRelayRequest, opts ...grpc.CallOption) (*ResumeRelayResponse, error) {
	out := new(ResumeRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/ResumeRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopRelay(ctx context.Context, in *StopRelayRequest, opts ...grpc.CallOption) (*StopRelayResponse, error) {
	out := new(StopRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteRelay(ctx context.Context, in *DeleteRelayRequest, opts ...grpc.CallOption) (*DeleteRelayResponse, error) {
	out := new(DeleteRelayResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteRelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetTunnel(ctx context.Context, in *GetTunnelRequest, opts ...grpc.CallOption) (*GetTunnelResponse, error) {
	out := new(GetTunnelResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetAllTunnels(ctx context.Context, in *GetAllTunnelsRequest, opts ...grpc.CallOption) (*GetAllTunnelsResponse, error) {
	out := new(GetAllTunnelsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetAllTunnels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) CreateTunnel(ctx context.Context, in *CreateTunnelRequest, opts ...grpc.CallOption) (*CreateTunnelResponse, error) {
	out := new(CreateTunnelResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/CreateTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) StopTunnel(ctx context.Context, in *StopTunnelRequest, opts ...grpc.CallOption) (*StopTunnelResponse, error) {
	out := new(StopTunnelResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/StopTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) ResumeTunnel(ctx context.Context, in *ResumeTunnelRequest, opts ...grpc.CallOption) (*ResumeTunnelResponse, error) {
	out := new(ResumeTunnelResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/ResumeTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) UpdateTunnel(ctx context.Context, in *UpdateTunnelRequest, opts ...grpc.CallOption) (*UpdateTunnelResponse, error) {
	out := new(UpdateTunnelResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/UpdateTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) DeleteTunnel(ctx context.Context, in *DeleteTunnelRequest, opts ...grpc.CallOption) (*DeleteTunnelResponse, error) {
	out := new(DeleteTunnelResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/DeleteTunnel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *plumberServerClient) GetServerOptions(ctx context.Context, in *GetServerOptionsRequest, opts ...grpc.CallOption) (*GetServerOptionsResponse, error) {
	out := new(GetServerOptionsResponse)
	err := c.cc.Invoke(ctx, "/protos.PlumberServer/GetServerOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlumberServerServer is the server API for PlumberServer service.
type PlumberServerServer interface {
	// List configured/known connections
	GetAllConnections(context.Context, *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error)
	// Fetch a specific connection by ID
	GetConnection(context.Context, *GetConnectionRequest) (*GetConnectionResponse, error)
	// Create a connection in plumber
	CreateConnection(context.Context, *CreateConnectionRequest) (*CreateConnectionResponse, error)
	// Test a connection before saving its configuration
	TestConnection(context.Context, *TestConnectionRequest) (*TestConnectionResponse, error)
	// Any active connections will be restarted
	UpdateConnection(context.Context, *UpdateConnectionRequest) (*UpdateConnectionResponse, error)
	// If there are any active connections, delete will cause them to get closed
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error)
	GetAllRelays(context.Context, *GetAllRelaysRequest) (*GetAllRelaysResponse, error)
	GetRelay(context.Context, *GetRelayRequest) (*GetRelayResponse, error)
	CreateRelay(context.Context, *CreateRelayRequest) (*CreateRelayResponse, error)
	UpdateRelay(context.Context, *UpdateRelayRequest) (*UpdateRelayResponse, error)
	ResumeRelay(context.Context, *ResumeRelayRequest) (*ResumeRelayResponse, error)
	StopRelay(context.Context, *StopRelayRequest) (*StopRelayResponse, error)
	DeleteRelay(context.Context, *DeleteRelayRequest) (*DeleteRelayResponse, error)
	GetTunnel(context.Context, *GetTunnelRequest) (*GetTunnelResponse, error)
	GetAllTunnels(context.Context, *GetAllTunnelsRequest) (*GetAllTunnelsResponse, error)
	CreateTunnel(context.Context, *CreateTunnelRequest) (*CreateTunnelResponse, error)
	StopTunnel(context.Context, *StopTunnelRequest) (*StopTunnelResponse, error)
	ResumeTunnel(context.Context, *ResumeTunnelRequest) (*ResumeTunnelResponse, error)
	UpdateTunnel(context.Context, *UpdateTunnelRequest) (*UpdateTunnelResponse, error)
	DeleteTunnel(context.Context, *DeleteTunnelRequest) (*DeleteTunnelResponse, error)
	GetServerOptions(context.Context, *GetServerOptionsRequest) (*GetServerOptionsResponse, error)
}

// UnimplementedPlumberServerServer can be embedded to have forward compatible implementations.
type UnimplementedPlumberServerServer struct {
}

func (*UnimplementedPlumberServerServer) GetAllConnections(ctx context.Context, req *GetAllConnectionsRequest) (*GetAllConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConnections not implemented")
}
func (*UnimplementedPlumberServerServer) GetConnection(ctx context.Context, req *GetConnectionRequest) (*GetConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (*UnimplementedPlumberServerServer) CreateConnection(ctx context.Context, req *CreateConnectionRequest) (*CreateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) TestConnection(ctx context.Context, req *TestConnectionRequest) (*TestConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestConnection not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateConnection(ctx context.Context, req *UpdateConnectionRequest) (*UpdateConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConnection not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteConnection(ctx context.Context, req *DeleteConnectionRequest) (*DeleteConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (*UnimplementedPlumberServerServer) GetAllRelays(ctx context.Context, req *GetAllRelaysRequest) (*GetAllRelaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRelays not implemented")
}
func (*UnimplementedPlumberServerServer) GetRelay(ctx context.Context, req *GetRelayRequest) (*GetRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelay not implemented")
}
func (*UnimplementedPlumberServerServer) CreateRelay(ctx context.Context, req *CreateRelayRequest) (*CreateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateRelay(ctx context.Context, req *UpdateRelayRequest) (*UpdateRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRelay not implemented")
}
func (*UnimplementedPlumberServerServer) ResumeRelay(ctx context.Context, req *ResumeRelayRequest) (*ResumeRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeRelay not implemented")
}
func (*UnimplementedPlumberServerServer) StopRelay(ctx context.Context, req *StopRelayRequest) (*StopRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopRelay not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteRelay(ctx context.Context, req *DeleteRelayRequest) (*DeleteRelayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelay not implemented")
}
func (*UnimplementedPlumberServerServer) GetTunnel(ctx context.Context, req *GetTunnelRequest) (*GetTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTunnel not implemented")
}
func (*UnimplementedPlumberServerServer) GetAllTunnels(ctx context.Context, req *GetAllTunnelsRequest) (*GetAllTunnelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTunnels not implemented")
}
func (*UnimplementedPlumberServerServer) CreateTunnel(ctx context.Context, req *CreateTunnelRequest) (*CreateTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTunnel not implemented")
}
func (*UnimplementedPlumberServerServer) StopTunnel(ctx context.Context, req *StopTunnelRequest) (*StopTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTunnel not implemented")
}
func (*UnimplementedPlumberServerServer) ResumeTunnel(ctx context.Context, req *ResumeTunnelRequest) (*ResumeTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeTunnel not implemented")
}
func (*UnimplementedPlumberServerServer) UpdateTunnel(ctx context.Context, req *UpdateTunnelRequest) (*UpdateTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTunnel not implemented")
}
func (*UnimplementedPlumberServerServer) DeleteTunnel(ctx context.Context, req *DeleteTunnelRequest) (*DeleteTunnelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTunnel not implemented")
}
func (*UnimplementedPlumberServerServer) GetServerOptions(ctx context.Context, req *GetServerOptionsRequest) (*GetServerOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerOptions not implemented")
}

func RegisterPlumberServerServer(s *grpc.Server, srv PlumberServerServer) {
	s.RegisterService(&_PlumberServer_serviceDesc, srv)
}

func _PlumberServer_GetAllConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllConnections(ctx, req.(*GetAllConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetConnection(ctx, req.(*GetConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateConnection(ctx, req.(*CreateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_TestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).TestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/TestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).TestConnection(ctx, req.(*TestConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateConnection(ctx, req.(*UpdateConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetAllRelays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRelaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllRelays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllRelays",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllRelays(ctx, req.(*GetAllRelaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetRelay(ctx, req.(*GetRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateRelay(ctx, req.(*CreateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateRelay(ctx, req.(*UpdateRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_ResumeRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).ResumeRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/ResumeRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).ResumeRelay(ctx, req.(*ResumeRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopRelay(ctx, req.(*StopRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteRelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteRelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteRelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteRelay(ctx, req.(*DeleteRelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetTunnel(ctx, req.(*GetTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetAllTunnels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTunnelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetAllTunnels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetAllTunnels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetAllTunnels(ctx, req.(*GetAllTunnelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_CreateTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).CreateTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/CreateTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).CreateTunnel(ctx, req.(*CreateTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_StopTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).StopTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/StopTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).StopTunnel(ctx, req.(*StopTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_ResumeTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).ResumeTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/ResumeTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).ResumeTunnel(ctx, req.(*ResumeTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_UpdateTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).UpdateTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/UpdateTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).UpdateTunnel(ctx, req.(*UpdateTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_DeleteTunnel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTunnelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).DeleteTunnel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/DeleteTunnel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).DeleteTunnel(ctx, req.(*DeleteTunnelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlumberServer_GetServerOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlumberServerServer).GetServerOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PlumberServer/GetServerOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlumberServerServer).GetServerOptions(ctx, req.(*GetServerOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlumberServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PlumberServer",
	HandlerType: (*PlumberServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllConnections",
			Handler:    _PlumberServer_GetAllConnections_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _PlumberServer_GetConnection_Handler,
		},
		{
			MethodName: "CreateConnection",
			Handler:    _PlumberServer_CreateConnection_Handler,
		},
		{
			MethodName: "TestConnection",
			Handler:    _PlumberServer_TestConnection_Handler,
		},
		{
			MethodName: "UpdateConnection",
			Handler:    _PlumberServer_UpdateConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _PlumberServer_DeleteConnection_Handler,
		},
		{
			MethodName: "GetAllRelays",
			Handler:    _PlumberServer_GetAllRelays_Handler,
		},
		{
			MethodName: "GetRelay",
			Handler:    _PlumberServer_GetRelay_Handler,
		},
		{
			MethodName: "CreateRelay",
			Handler:    _PlumberServer_CreateRelay_Handler,
		},
		{
			MethodName: "UpdateRelay",
			Handler:    _PlumberServer_UpdateRelay_Handler,
		},
		{
			MethodName: "ResumeRelay",
			Handler:    _PlumberServer_ResumeRelay_Handler,
		},
		{
			MethodName: "StopRelay",
			Handler:    _PlumberServer_StopRelay_Handler,
		},
		{
			MethodName: "DeleteRelay",
			Handler:    _PlumberServer_DeleteRelay_Handler,
		},
		{
			MethodName: "GetTunnel",
			Handler:    _PlumberServer_GetTunnel_Handler,
		},
		{
			MethodName: "GetAllTunnels",
			Handler:    _PlumberServer_GetAllTunnels_Handler,
		},
		{
			MethodName: "CreateTunnel",
			Handler:    _PlumberServer_CreateTunnel_Handler,
		},
		{
			MethodName: "StopTunnel",
			Handler:    _PlumberServer_StopTunnel_Handler,
		},
		{
			MethodName: "ResumeTunnel",
			Handler:    _PlumberServer_ResumeTunnel_Handler,
		},
		{
			MethodName: "UpdateTunnel",
			Handler:    _PlumberServer_UpdateTunnel_Handler,
		},
		{
			MethodName: "DeleteTunnel",
			Handler:    _PlumberServer_DeleteTunnel_Handler,
		},
		{
			MethodName: "GetServerOptions",
			Handler:    _PlumberServer_GetServerOptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ps_base.proto",
}
