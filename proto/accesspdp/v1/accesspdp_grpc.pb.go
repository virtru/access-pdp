// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: accesspdp/v1/accesspdp.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/accesspdp.v1.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/accesspdp.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accesspdp.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accesspdp/v1/accesspdp.proto",
}

// AccessPDPEndpointClient is the client API for AccessPDPEndpoint service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessPDPEndpointClient interface {
	DetermineAccess(ctx context.Context, in *DetermineAccessRequest, opts ...grpc.CallOption) (AccessPDPEndpoint_DetermineAccessClient, error)
}

type accessPDPEndpointClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessPDPEndpointClient(cc grpc.ClientConnInterface) AccessPDPEndpointClient {
	return &accessPDPEndpointClient{cc}
}

func (c *accessPDPEndpointClient) DetermineAccess(ctx context.Context, in *DetermineAccessRequest, opts ...grpc.CallOption) (AccessPDPEndpoint_DetermineAccessClient, error) {
	stream, err := c.cc.NewStream(ctx, &AccessPDPEndpoint_ServiceDesc.Streams[0], "/accesspdp.v1.AccessPDPEndpoint/DetermineAccess", opts...)
	if err != nil {
		return nil, err
	}
	x := &accessPDPEndpointDetermineAccessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccessPDPEndpoint_DetermineAccessClient interface {
	Recv() (*DetermineAccessResponse, error)
	grpc.ClientStream
}

type accessPDPEndpointDetermineAccessClient struct {
	grpc.ClientStream
}

func (x *accessPDPEndpointDetermineAccessClient) Recv() (*DetermineAccessResponse, error) {
	m := new(DetermineAccessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccessPDPEndpointServer is the server API for AccessPDPEndpoint service.
// All implementations must embed UnimplementedAccessPDPEndpointServer
// for forward compatibility
type AccessPDPEndpointServer interface {
	DetermineAccess(*DetermineAccessRequest, AccessPDPEndpoint_DetermineAccessServer) error
	mustEmbedUnimplementedAccessPDPEndpointServer()
}

// UnimplementedAccessPDPEndpointServer must be embedded to have forward compatible implementations.
type UnimplementedAccessPDPEndpointServer struct {
}

func (UnimplementedAccessPDPEndpointServer) DetermineAccess(*DetermineAccessRequest, AccessPDPEndpoint_DetermineAccessServer) error {
	return status.Errorf(codes.Unimplemented, "method DetermineAccess not implemented")
}
func (UnimplementedAccessPDPEndpointServer) mustEmbedUnimplementedAccessPDPEndpointServer() {}

// UnsafeAccessPDPEndpointServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessPDPEndpointServer will
// result in compilation errors.
type UnsafeAccessPDPEndpointServer interface {
	mustEmbedUnimplementedAccessPDPEndpointServer()
}

func RegisterAccessPDPEndpointServer(s grpc.ServiceRegistrar, srv AccessPDPEndpointServer) {
	s.RegisterService(&AccessPDPEndpoint_ServiceDesc, srv)
}

func _AccessPDPEndpoint_DetermineAccess_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DetermineAccessRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccessPDPEndpointServer).DetermineAccess(m, &accessPDPEndpointDetermineAccessServer{stream})
}

type AccessPDPEndpoint_DetermineAccessServer interface {
	Send(*DetermineAccessResponse) error
	grpc.ServerStream
}

type accessPDPEndpointDetermineAccessServer struct {
	grpc.ServerStream
}

func (x *accessPDPEndpointDetermineAccessServer) Send(m *DetermineAccessResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AccessPDPEndpoint_ServiceDesc is the grpc.ServiceDesc for AccessPDPEndpoint service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessPDPEndpoint_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accesspdp.v1.AccessPDPEndpoint",
	HandlerType: (*AccessPDPEndpointServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DetermineAccess",
			Handler:       _AccessPDPEndpoint_DetermineAccess_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "accesspdp/v1/accesspdp.proto",
}
