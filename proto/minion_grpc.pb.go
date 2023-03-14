// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/minion.proto

package proto

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

// MinionServiceClient is the client API for MinionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MinionServiceClient interface {
	// This one is just for testing
	// rpc Create (CreateMinion) returns (MinionCreateResponse) {}
	RegisterNewMinion(ctx context.Context, in *CreateMinion, opts ...grpc.CallOption) (*MinionCreateResponse, error)
	GetCommands(ctx context.Context, opts ...grpc.CallOption) (MinionService_GetCommandsClient, error)
	CmdRun(ctx context.Context, in *CmdRunSend, opts ...grpc.CallOption) (*CmdRunResult, error)
}

type minionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMinionServiceClient(cc grpc.ClientConnInterface) MinionServiceClient {
	return &minionServiceClient{cc}
}

func (c *minionServiceClient) RegisterNewMinion(ctx context.Context, in *CreateMinion, opts ...grpc.CallOption) (*MinionCreateResponse, error) {
	out := new(MinionCreateResponse)
	err := c.cc.Invoke(ctx, "/minion.MinionService/RegisterNewMinion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionServiceClient) GetCommands(ctx context.Context, opts ...grpc.CallOption) (MinionService_GetCommandsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MinionService_ServiceDesc.Streams[0], "/minion.MinionService/GetCommands", opts...)
	if err != nil {
		return nil, err
	}
	x := &minionServiceGetCommandsClient{stream}
	return x, nil
}

type MinionService_GetCommandsClient interface {
	Send(*CommandFromMinion) error
	Recv() (*CommandToMinion, error)
	grpc.ClientStream
}

type minionServiceGetCommandsClient struct {
	grpc.ClientStream
}

func (x *minionServiceGetCommandsClient) Send(m *CommandFromMinion) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minionServiceGetCommandsClient) Recv() (*CommandToMinion, error) {
	m := new(CommandToMinion)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minionServiceClient) CmdRun(ctx context.Context, in *CmdRunSend, opts ...grpc.CallOption) (*CmdRunResult, error) {
	out := new(CmdRunResult)
	err := c.cc.Invoke(ctx, "/minion.MinionService/CmdRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinionServiceServer is the server API for MinionService service.
// All implementations must embed UnimplementedMinionServiceServer
// for forward compatibility
type MinionServiceServer interface {
	// This one is just for testing
	// rpc Create (CreateMinion) returns (MinionCreateResponse) {}
	RegisterNewMinion(context.Context, *CreateMinion) (*MinionCreateResponse, error)
	GetCommands(MinionService_GetCommandsServer) error
	CmdRun(context.Context, *CmdRunSend) (*CmdRunResult, error)
	mustEmbedUnimplementedMinionServiceServer()
}

// UnimplementedMinionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMinionServiceServer struct {
}

func (UnimplementedMinionServiceServer) RegisterNewMinion(context.Context, *CreateMinion) (*MinionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewMinion not implemented")
}
func (UnimplementedMinionServiceServer) GetCommands(MinionService_GetCommandsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommands not implemented")
}
func (UnimplementedMinionServiceServer) CmdRun(context.Context, *CmdRunSend) (*CmdRunResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CmdRun not implemented")
}
func (UnimplementedMinionServiceServer) mustEmbedUnimplementedMinionServiceServer() {}

// UnsafeMinionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MinionServiceServer will
// result in compilation errors.
type UnsafeMinionServiceServer interface {
	mustEmbedUnimplementedMinionServiceServer()
}

func RegisterMinionServiceServer(s grpc.ServiceRegistrar, srv MinionServiceServer) {
	s.RegisterService(&MinionService_ServiceDesc, srv)
}

func _MinionService_RegisterNewMinion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMinion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServiceServer).RegisterNewMinion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minion.MinionService/RegisterNewMinion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServiceServer).RegisterNewMinion(ctx, req.(*CreateMinion))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinionService_GetCommands_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinionServiceServer).GetCommands(&minionServiceGetCommandsServer{stream})
}

type MinionService_GetCommandsServer interface {
	Send(*CommandToMinion) error
	Recv() (*CommandFromMinion, error)
	grpc.ServerStream
}

type minionServiceGetCommandsServer struct {
	grpc.ServerStream
}

func (x *minionServiceGetCommandsServer) Send(m *CommandToMinion) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minionServiceGetCommandsServer) Recv() (*CommandFromMinion, error) {
	m := new(CommandFromMinion)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MinionService_CmdRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdRunSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServiceServer).CmdRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/minion.MinionService/CmdRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServiceServer).CmdRun(ctx, req.(*CmdRunSend))
	}
	return interceptor(ctx, in, info, handler)
}

// MinionService_ServiceDesc is the grpc.ServiceDesc for MinionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MinionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "minion.MinionService",
	HandlerType: (*MinionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewMinion",
			Handler:    _MinionService_RegisterNewMinion_Handler,
		},
		{
			MethodName: "CmdRun",
			Handler:    _MinionService_CmdRun_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCommands",
			Handler:       _MinionService_GetCommands_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/minion.proto",
}
