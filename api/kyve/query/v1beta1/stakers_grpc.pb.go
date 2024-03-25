// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package queryv1beta1

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

// QueryStakersClient is the client API for QueryStakers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryStakersClient interface {
	// Stakers queries for all stakers.
	Stakers(ctx context.Context, in *QueryStakersRequest, opts ...grpc.CallOption) (*QueryStakersResponse, error)
	// Staker queries for all stakers.
	Staker(ctx context.Context, in *QueryStakerRequest, opts ...grpc.CallOption) (*QueryStakerResponse, error)
	// StakersByPool queries for all stakers that are currently participating in the given pool
	StakersByPool(ctx context.Context, in *QueryStakersByPoolRequest, opts ...grpc.CallOption) (*QueryStakersByPoolResponse, error)
	// StakersByPool queries for all stakers and sorted them first by number of pools participating and
	// then by delegation
	StakersByPoolCount(ctx context.Context, in *QueryStakersByPoolCountRequest, opts ...grpc.CallOption) (*QueryStakersByPoolCountResponse, error)
}

type queryStakersClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryStakersClient(cc grpc.ClientConnInterface) QueryStakersClient {
	return &queryStakersClient{cc}
}

func (c *queryStakersClient) Stakers(ctx context.Context, in *QueryStakersRequest, opts ...grpc.CallOption) (*QueryStakersResponse, error) {
	out := new(QueryStakersResponse)
	err := c.cc.Invoke(ctx, "/kyve.query.v1beta1.QueryStakers/Stakers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryStakersClient) Staker(ctx context.Context, in *QueryStakerRequest, opts ...grpc.CallOption) (*QueryStakerResponse, error) {
	out := new(QueryStakerResponse)
	err := c.cc.Invoke(ctx, "/kyve.query.v1beta1.QueryStakers/Staker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryStakersClient) StakersByPool(ctx context.Context, in *QueryStakersByPoolRequest, opts ...grpc.CallOption) (*QueryStakersByPoolResponse, error) {
	out := new(QueryStakersByPoolResponse)
	err := c.cc.Invoke(ctx, "/kyve.query.v1beta1.QueryStakers/StakersByPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryStakersClient) StakersByPoolCount(ctx context.Context, in *QueryStakersByPoolCountRequest, opts ...grpc.CallOption) (*QueryStakersByPoolCountResponse, error) {
	out := new(QueryStakersByPoolCountResponse)
	err := c.cc.Invoke(ctx, "/kyve.query.v1beta1.QueryStakers/StakersByPoolCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryStakersServer is the server API for QueryStakers service.
// All implementations must embed UnimplementedQueryStakersServer
// for forward compatibility
type QueryStakersServer interface {
	// Stakers queries for all stakers.
	Stakers(context.Context, *QueryStakersRequest) (*QueryStakersResponse, error)
	// Staker queries for all stakers.
	Staker(context.Context, *QueryStakerRequest) (*QueryStakerResponse, error)
	// StakersByPool queries for all stakers that are currently participating in the given pool
	StakersByPool(context.Context, *QueryStakersByPoolRequest) (*QueryStakersByPoolResponse, error)
	// StakersByPool queries for all stakers and sorted them first by number of pools participating and
	// then by delegation
	StakersByPoolCount(context.Context, *QueryStakersByPoolCountRequest) (*QueryStakersByPoolCountResponse, error)
	mustEmbedUnimplementedQueryStakersServer()
}

// UnimplementedQueryStakersServer must be embedded to have forward compatible implementations.
type UnimplementedQueryStakersServer struct {
}

func (UnimplementedQueryStakersServer) Stakers(context.Context, *QueryStakersRequest) (*QueryStakersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stakers not implemented")
}
func (UnimplementedQueryStakersServer) Staker(context.Context, *QueryStakerRequest) (*QueryStakerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Staker not implemented")
}
func (UnimplementedQueryStakersServer) StakersByPool(context.Context, *QueryStakersByPoolRequest) (*QueryStakersByPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakersByPool not implemented")
}
func (UnimplementedQueryStakersServer) StakersByPoolCount(context.Context, *QueryStakersByPoolCountRequest) (*QueryStakersByPoolCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakersByPoolCount not implemented")
}
func (UnimplementedQueryStakersServer) mustEmbedUnimplementedQueryStakersServer() {}

// UnsafeQueryStakersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryStakersServer will
// result in compilation errors.
type UnsafeQueryStakersServer interface {
	mustEmbedUnimplementedQueryStakersServer()
}

func RegisterQueryStakersServer(s grpc.ServiceRegistrar, srv QueryStakersServer) {
	s.RegisterService(&QueryStakers_ServiceDesc, srv)
}

func _QueryStakers_Stakers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryStakersServer).Stakers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.query.v1beta1.QueryStakers/Stakers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryStakersServer).Stakers(ctx, req.(*QueryStakersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryStakers_Staker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryStakersServer).Staker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.query.v1beta1.QueryStakers/Staker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryStakersServer).Staker(ctx, req.(*QueryStakerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryStakers_StakersByPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakersByPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryStakersServer).StakersByPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.query.v1beta1.QueryStakers/StakersByPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryStakersServer).StakersByPool(ctx, req.(*QueryStakersByPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryStakers_StakersByPoolCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStakersByPoolCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryStakersServer).StakersByPoolCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.query.v1beta1.QueryStakers/StakersByPoolCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryStakersServer).StakersByPoolCount(ctx, req.(*QueryStakersByPoolCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueryStakers_ServiceDesc is the grpc.ServiceDesc for QueryStakers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryStakers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kyve.query.v1beta1.QueryStakers",
	HandlerType: (*QueryStakersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Stakers",
			Handler:    _QueryStakers_Stakers_Handler,
		},
		{
			MethodName: "Staker",
			Handler:    _QueryStakers_Staker_Handler,
		},
		{
			MethodName: "StakersByPool",
			Handler:    _QueryStakers_StakersByPool_Handler,
		},
		{
			MethodName: "StakersByPoolCount",
			Handler:    _QueryStakers_StakersByPoolCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kyve/query/v1beta1/stakers.proto",
}
