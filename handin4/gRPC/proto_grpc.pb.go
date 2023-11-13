// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: gRPC/proto.proto

package gRPC

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

const (
	Mutex_Election_FullMethodName    = "/PhysicalTime.Mutex/Election"
	Mutex_Coordinator_FullMethodName = "/PhysicalTime.Mutex/Coordinator"
)

// MutexClient is the client API for Mutex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MutexClient interface {
	Election(ctx context.Context, in *ElectionMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
	Coordinator(ctx context.Context, in *CoordinatorMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type mutexClient struct {
	cc grpc.ClientConnInterface
}

func NewMutexClient(cc grpc.ClientConnInterface) MutexClient {
	return &mutexClient{cc}
}

func (c *mutexClient) Election(ctx context.Context, in *ElectionMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, Mutex_Election_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mutexClient) Coordinator(ctx context.Context, in *CoordinatorMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, Mutex_Coordinator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MutexServer is the server API for Mutex service.
// All implementations must embed UnimplementedMutexServer
// for forward compatibility
type MutexServer interface {
	Election(context.Context, *ElectionMessage) (*EmptyMessage, error)
	Coordinator(context.Context, *CoordinatorMessage) (*EmptyMessage, error)
	mustEmbedUnimplementedMutexServer()
}

// UnimplementedMutexServer must be embedded to have forward compatible implementations.
type UnimplementedMutexServer struct {
}

func (UnimplementedMutexServer) Election(context.Context, *ElectionMessage) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedMutexServer) Coordinator(context.Context, *CoordinatorMessage) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coordinator not implemented")
}
func (UnimplementedMutexServer) mustEmbedUnimplementedMutexServer() {}

// UnsafeMutexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MutexServer will
// result in compilation errors.
type UnsafeMutexServer interface {
	mustEmbedUnimplementedMutexServer()
}

func RegisterMutexServer(s grpc.ServiceRegistrar, srv MutexServer) {
	s.RegisterService(&Mutex_ServiceDesc, srv)
}

func _Mutex_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElectionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mutex_Election_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexServer).Election(ctx, req.(*ElectionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mutex_Coordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoordinatorMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MutexServer).Coordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mutex_Coordinator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MutexServer).Coordinator(ctx, req.(*CoordinatorMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Mutex_ServiceDesc is the grpc.ServiceDesc for Mutex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mutex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PhysicalTime.Mutex",
	HandlerType: (*MutexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Election",
			Handler:    _Mutex_Election_Handler,
		},
		{
			MethodName: "Coordinator",
			Handler:    _Mutex_Coordinator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPC/proto.proto",
}