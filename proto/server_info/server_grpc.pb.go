// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server_info

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

// PushToServerClient is the client API for PushToServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushToServerClient interface {
	PushTask(ctx context.Context, in *CommandItem, opts ...grpc.CallOption) (*CommandItem, error)
}

type pushToServerClient struct {
	cc grpc.ClientConnInterface
}

func NewPushToServerClient(cc grpc.ClientConnInterface) PushToServerClient {
	return &pushToServerClient{cc}
}

func (c *pushToServerClient) PushTask(ctx context.Context, in *CommandItem, opts ...grpc.CallOption) (*CommandItem, error) {
	out := new(CommandItem)
	err := c.cc.Invoke(ctx, "/server_info.PushToServer/PushTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushToServerServer is the server API for PushToServer service.
// All implementations should embed UnimplementedPushToServerServer
// for forward compatibility
type PushToServerServer interface {
	PushTask(context.Context, *CommandItem) (*CommandItem, error)
}

// UnimplementedPushToServerServer should be embedded to have forward compatible implementations.
type UnimplementedPushToServerServer struct {
}

func (UnimplementedPushToServerServer) PushTask(context.Context, *CommandItem) (*CommandItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushTask not implemented")
}

// UnsafePushToServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushToServerServer will
// result in compilation errors.
type UnsafePushToServerServer interface {
	mustEmbedUnimplementedPushToServerServer()
}

func RegisterPushToServerServer(s grpc.ServiceRegistrar, srv PushToServerServer) {
	s.RegisterService(&PushToServer_ServiceDesc, srv)
}

func _PushToServer_PushTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushToServerServer).PushTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server_info.PushToServer/PushTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushToServerServer).PushTask(ctx, req.(*CommandItem))
	}
	return interceptor(ctx, in, info, handler)
}

// PushToServer_ServiceDesc is the grpc.ServiceDesc for PushToServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushToServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server_info.PushToServer",
	HandlerType: (*PushToServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushTask",
			Handler:    _PushToServer_PushTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server_info/server.proto",
}
