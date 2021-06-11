// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_feedback_api

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

// OcpFeedbackApiClient is the client API for OcpFeedbackApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpFeedbackApiClient interface {
	CreateFeedbackV1(ctx context.Context, in *CreateFeedbackV1Request, opts ...grpc.CallOption) (*CreateFeedbackV1Response, error)
	CreateMultiFeedbackV1(ctx context.Context, in *CreateMultiFeedbackV1Request, opts ...grpc.CallOption) (*CreateMultiFeedbackV1Response, error)
	RemoveFeedbackV1(ctx context.Context, in *RemoveFeedbackV1Request, opts ...grpc.CallOption) (*RemoveFeedbackV1Response, error)
	DescribeFeedbackV1(ctx context.Context, in *DescribeFeedbackV1Request, opts ...grpc.CallOption) (*DescribeFeedbackV1Response, error)
	ListFeedbacksV1(ctx context.Context, in *ListFeedbacksV1Request, opts ...grpc.CallOption) (*ListFeedbacksV1Response, error)
	CreateProposalV1(ctx context.Context, in *CreateProposalV1Request, opts ...grpc.CallOption) (*CreateProposalV1Response, error)
	CreateMultiProposalV1(ctx context.Context, in *CreateMultiProposalV1Request, opts ...grpc.CallOption) (*CreateMultiProposalV1Response, error)
	RemoveProposalV1(ctx context.Context, in *RemoveProposalV1Request, opts ...grpc.CallOption) (*RemoveProposalV1Response, error)
	DescribeProposalV1(ctx context.Context, in *DescribeProposalV1Request, opts ...grpc.CallOption) (*DescribeProposalV1Response, error)
	ListProposalsV1(ctx context.Context, in *ListProposalsV1Request, opts ...grpc.CallOption) (*ListProposalsV1Response, error)
}

type ocpFeedbackApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpFeedbackApiClient(cc grpc.ClientConnInterface) OcpFeedbackApiClient {
	return &ocpFeedbackApiClient{cc}
}

func (c *ocpFeedbackApiClient) CreateFeedbackV1(ctx context.Context, in *CreateFeedbackV1Request, opts ...grpc.CallOption) (*CreateFeedbackV1Response, error) {
	out := new(CreateFeedbackV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/CreateFeedbackV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) CreateMultiFeedbackV1(ctx context.Context, in *CreateMultiFeedbackV1Request, opts ...grpc.CallOption) (*CreateMultiFeedbackV1Response, error) {
	out := new(CreateMultiFeedbackV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/CreateMultiFeedbackV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) RemoveFeedbackV1(ctx context.Context, in *RemoveFeedbackV1Request, opts ...grpc.CallOption) (*RemoveFeedbackV1Response, error) {
	out := new(RemoveFeedbackV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/RemoveFeedbackV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) DescribeFeedbackV1(ctx context.Context, in *DescribeFeedbackV1Request, opts ...grpc.CallOption) (*DescribeFeedbackV1Response, error) {
	out := new(DescribeFeedbackV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/DescribeFeedbackV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) ListFeedbacksV1(ctx context.Context, in *ListFeedbacksV1Request, opts ...grpc.CallOption) (*ListFeedbacksV1Response, error) {
	out := new(ListFeedbacksV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/ListFeedbacksV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) CreateProposalV1(ctx context.Context, in *CreateProposalV1Request, opts ...grpc.CallOption) (*CreateProposalV1Response, error) {
	out := new(CreateProposalV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/CreateProposalV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) CreateMultiProposalV1(ctx context.Context, in *CreateMultiProposalV1Request, opts ...grpc.CallOption) (*CreateMultiProposalV1Response, error) {
	out := new(CreateMultiProposalV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/CreateMultiProposalV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) RemoveProposalV1(ctx context.Context, in *RemoveProposalV1Request, opts ...grpc.CallOption) (*RemoveProposalV1Response, error) {
	out := new(RemoveProposalV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/RemoveProposalV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) DescribeProposalV1(ctx context.Context, in *DescribeProposalV1Request, opts ...grpc.CallOption) (*DescribeProposalV1Response, error) {
	out := new(DescribeProposalV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/DescribeProposalV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpFeedbackApiClient) ListProposalsV1(ctx context.Context, in *ListProposalsV1Request, opts ...grpc.CallOption) (*ListProposalsV1Response, error) {
	out := new(ListProposalsV1Response)
	err := c.cc.Invoke(ctx, "/ocp.feedback.api.OcpFeedbackApi/ListProposalsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpFeedbackApiServer is the server API for OcpFeedbackApi service.
// All implementations must embed UnimplementedOcpFeedbackApiServer
// for forward compatibility
type OcpFeedbackApiServer interface {
	CreateFeedbackV1(context.Context, *CreateFeedbackV1Request) (*CreateFeedbackV1Response, error)
	CreateMultiFeedbackV1(context.Context, *CreateMultiFeedbackV1Request) (*CreateMultiFeedbackV1Response, error)
	RemoveFeedbackV1(context.Context, *RemoveFeedbackV1Request) (*RemoveFeedbackV1Response, error)
	DescribeFeedbackV1(context.Context, *DescribeFeedbackV1Request) (*DescribeFeedbackV1Response, error)
	ListFeedbacksV1(context.Context, *ListFeedbacksV1Request) (*ListFeedbacksV1Response, error)
	CreateProposalV1(context.Context, *CreateProposalV1Request) (*CreateProposalV1Response, error)
	CreateMultiProposalV1(context.Context, *CreateMultiProposalV1Request) (*CreateMultiProposalV1Response, error)
	RemoveProposalV1(context.Context, *RemoveProposalV1Request) (*RemoveProposalV1Response, error)
	DescribeProposalV1(context.Context, *DescribeProposalV1Request) (*DescribeProposalV1Response, error)
	ListProposalsV1(context.Context, *ListProposalsV1Request) (*ListProposalsV1Response, error)
	mustEmbedUnimplementedOcpFeedbackApiServer()
}

// UnimplementedOcpFeedbackApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpFeedbackApiServer struct {
}

func (UnimplementedOcpFeedbackApiServer) CreateFeedbackV1(context.Context, *CreateFeedbackV1Request) (*CreateFeedbackV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFeedbackV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) CreateMultiFeedbackV1(context.Context, *CreateMultiFeedbackV1Request) (*CreateMultiFeedbackV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMultiFeedbackV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) RemoveFeedbackV1(context.Context, *RemoveFeedbackV1Request) (*RemoveFeedbackV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFeedbackV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) DescribeFeedbackV1(context.Context, *DescribeFeedbackV1Request) (*DescribeFeedbackV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeFeedbackV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) ListFeedbacksV1(context.Context, *ListFeedbacksV1Request) (*ListFeedbacksV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFeedbacksV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) CreateProposalV1(context.Context, *CreateProposalV1Request) (*CreateProposalV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProposalV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) CreateMultiProposalV1(context.Context, *CreateMultiProposalV1Request) (*CreateMultiProposalV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMultiProposalV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) RemoveProposalV1(context.Context, *RemoveProposalV1Request) (*RemoveProposalV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProposalV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) DescribeProposalV1(context.Context, *DescribeProposalV1Request) (*DescribeProposalV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProposalV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) ListProposalsV1(context.Context, *ListProposalsV1Request) (*ListProposalsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProposalsV1 not implemented")
}
func (UnimplementedOcpFeedbackApiServer) mustEmbedUnimplementedOcpFeedbackApiServer() {}

// UnsafeOcpFeedbackApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpFeedbackApiServer will
// result in compilation errors.
type UnsafeOcpFeedbackApiServer interface {
	mustEmbedUnimplementedOcpFeedbackApiServer()
}

func RegisterOcpFeedbackApiServer(s grpc.ServiceRegistrar, srv OcpFeedbackApiServer) {
	s.RegisterService(&OcpFeedbackApi_ServiceDesc, srv)
}

func _OcpFeedbackApi_CreateFeedbackV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFeedbackV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).CreateFeedbackV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/CreateFeedbackV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).CreateFeedbackV1(ctx, req.(*CreateFeedbackV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_CreateMultiFeedbackV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMultiFeedbackV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).CreateMultiFeedbackV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/CreateMultiFeedbackV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).CreateMultiFeedbackV1(ctx, req.(*CreateMultiFeedbackV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_RemoveFeedbackV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFeedbackV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).RemoveFeedbackV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/RemoveFeedbackV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).RemoveFeedbackV1(ctx, req.(*RemoveFeedbackV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_DescribeFeedbackV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeFeedbackV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).DescribeFeedbackV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/DescribeFeedbackV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).DescribeFeedbackV1(ctx, req.(*DescribeFeedbackV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_ListFeedbacksV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFeedbacksV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).ListFeedbacksV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/ListFeedbacksV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).ListFeedbacksV1(ctx, req.(*ListFeedbacksV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_CreateProposalV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProposalV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).CreateProposalV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/CreateProposalV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).CreateProposalV1(ctx, req.(*CreateProposalV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_CreateMultiProposalV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMultiProposalV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).CreateMultiProposalV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/CreateMultiProposalV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).CreateMultiProposalV1(ctx, req.(*CreateMultiProposalV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_RemoveProposalV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProposalV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).RemoveProposalV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/RemoveProposalV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).RemoveProposalV1(ctx, req.(*RemoveProposalV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_DescribeProposalV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProposalV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).DescribeProposalV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/DescribeProposalV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).DescribeProposalV1(ctx, req.(*DescribeProposalV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpFeedbackApi_ListProposalsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProposalsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpFeedbackApiServer).ListProposalsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.feedback.api.OcpFeedbackApi/ListProposalsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpFeedbackApiServer).ListProposalsV1(ctx, req.(*ListProposalsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpFeedbackApi_ServiceDesc is the grpc.ServiceDesc for OcpFeedbackApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpFeedbackApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.feedback.api.OcpFeedbackApi",
	HandlerType: (*OcpFeedbackApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFeedbackV1",
			Handler:    _OcpFeedbackApi_CreateFeedbackV1_Handler,
		},
		{
			MethodName: "CreateMultiFeedbackV1",
			Handler:    _OcpFeedbackApi_CreateMultiFeedbackV1_Handler,
		},
		{
			MethodName: "RemoveFeedbackV1",
			Handler:    _OcpFeedbackApi_RemoveFeedbackV1_Handler,
		},
		{
			MethodName: "DescribeFeedbackV1",
			Handler:    _OcpFeedbackApi_DescribeFeedbackV1_Handler,
		},
		{
			MethodName: "ListFeedbacksV1",
			Handler:    _OcpFeedbackApi_ListFeedbacksV1_Handler,
		},
		{
			MethodName: "CreateProposalV1",
			Handler:    _OcpFeedbackApi_CreateProposalV1_Handler,
		},
		{
			MethodName: "CreateMultiProposalV1",
			Handler:    _OcpFeedbackApi_CreateMultiProposalV1_Handler,
		},
		{
			MethodName: "RemoveProposalV1",
			Handler:    _OcpFeedbackApi_RemoveProposalV1_Handler,
		},
		{
			MethodName: "DescribeProposalV1",
			Handler:    _OcpFeedbackApi_DescribeProposalV1_Handler,
		},
		{
			MethodName: "ListProposalsV1",
			Handler:    _OcpFeedbackApi_ListProposalsV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedback-service.proto",
}
