// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_reason_api

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReasonRpcClient is the client API for ReasonRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReasonRpcClient interface {
	CreateReason(ctx context.Context, in *CreateReasonRequest, opts ...grpc.CallOption) (*CreateReasonResponse, error)
	DescribeReason(ctx context.Context, in *DescribeReasonRequest, opts ...grpc.CallOption) (*DescribeReasonResponse, error)
	ListReasons(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListReasonsResponse, error)
	RemoveReason(ctx context.Context, in *RemoveReasonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type reasonRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewReasonRpcClient(cc grpc.ClientConnInterface) ReasonRpcClient {
	return &reasonRpcClient{cc}
}

func (c *reasonRpcClient) CreateReason(ctx context.Context, in *CreateReasonRequest, opts ...grpc.CallOption) (*CreateReasonResponse, error) {
	out := new(CreateReasonResponse)
	err := c.cc.Invoke(ctx, "/ova.reason.api.ReasonRpc/CreateReason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reasonRpcClient) DescribeReason(ctx context.Context, in *DescribeReasonRequest, opts ...grpc.CallOption) (*DescribeReasonResponse, error) {
	out := new(DescribeReasonResponse)
	err := c.cc.Invoke(ctx, "/ova.reason.api.ReasonRpc/DescribeReason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reasonRpcClient) ListReasons(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListReasonsResponse, error) {
	out := new(ListReasonsResponse)
	err := c.cc.Invoke(ctx, "/ova.reason.api.ReasonRpc/ListReasons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reasonRpcClient) RemoveReason(ctx context.Context, in *RemoveReasonRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.reason.api.ReasonRpc/RemoveReason", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReasonRpcServer is the server API for ReasonRpc service.
// All implementations must embed UnimplementedReasonRpcServer
// for forward compatibility
type ReasonRpcServer interface {
	CreateReason(context.Context, *CreateReasonRequest) (*CreateReasonResponse, error)
	DescribeReason(context.Context, *DescribeReasonRequest) (*DescribeReasonResponse, error)
	ListReasons(context.Context, *empty.Empty) (*ListReasonsResponse, error)
	RemoveReason(context.Context, *RemoveReasonRequest) (*empty.Empty, error)
	mustEmbedUnimplementedReasonRpcServer()
}

// UnimplementedReasonRpcServer must be embedded to have forward compatible implementations.
type UnimplementedReasonRpcServer struct {
}

func (UnimplementedReasonRpcServer) CreateReason(context.Context, *CreateReasonRequest) (*CreateReasonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReason not implemented")
}
func (UnimplementedReasonRpcServer) DescribeReason(context.Context, *DescribeReasonRequest) (*DescribeReasonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeReason not implemented")
}
func (UnimplementedReasonRpcServer) ListReasons(context.Context, *empty.Empty) (*ListReasonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReasons not implemented")
}
func (UnimplementedReasonRpcServer) RemoveReason(context.Context, *RemoveReasonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReason not implemented")
}
func (UnimplementedReasonRpcServer) mustEmbedUnimplementedReasonRpcServer() {}

// UnsafeReasonRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReasonRpcServer will
// result in compilation errors.
type UnsafeReasonRpcServer interface {
	mustEmbedUnimplementedReasonRpcServer()
}

func RegisterReasonRpcServer(s grpc.ServiceRegistrar, srv ReasonRpcServer) {
	s.RegisterService(&ReasonRpc_ServiceDesc, srv)
}

func _ReasonRpc_CreateReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReasonRpcServer).CreateReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.reason.api.ReasonRpc/CreateReason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReasonRpcServer).CreateReason(ctx, req.(*CreateReasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReasonRpc_DescribeReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeReasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReasonRpcServer).DescribeReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.reason.api.ReasonRpc/DescribeReason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReasonRpcServer).DescribeReason(ctx, req.(*DescribeReasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReasonRpc_ListReasons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReasonRpcServer).ListReasons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.reason.api.ReasonRpc/ListReasons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReasonRpcServer).ListReasons(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReasonRpc_RemoveReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReasonRpcServer).RemoveReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.reason.api.ReasonRpc/RemoveReason",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReasonRpcServer).RemoveReason(ctx, req.(*RemoveReasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReasonRpc_ServiceDesc is the grpc.ServiceDesc for ReasonRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReasonRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.reason.api.ReasonRpc",
	HandlerType: (*ReasonRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReason",
			Handler:    _ReasonRpc_CreateReason_Handler,
		},
		{
			MethodName: "DescribeReason",
			Handler:    _ReasonRpc_DescribeReason_Handler,
		},
		{
			MethodName: "ListReasons",
			Handler:    _ReasonRpc_ListReasons_Handler,
		},
		{
			MethodName: "RemoveReason",
			Handler:    _ReasonRpc_RemoveReason_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ova-reason-api.proto",
}
