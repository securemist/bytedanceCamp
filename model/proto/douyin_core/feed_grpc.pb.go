// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: feed.proto

package douyin_core

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

// FeedClient is the client API for Feed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedClient interface {
	GetFeed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error)
	PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error)
	PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error)
}

type feedClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedClient(cc grpc.ClientConnInterface) FeedClient {
	return &feedClient{cc}
}

func (c *feedClient) GetFeed(ctx context.Context, in *FeedRequest, opts ...grpc.CallOption) (*FeedResponse, error) {
	out := new(FeedResponse)
	err := c.cc.Invoke(ctx, "/douyin.core.Feed/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error) {
	out := new(PublishVideoResponse)
	err := c.cc.Invoke(ctx, "/douyin.core.Feed/PublishVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedClient) PublishList(ctx context.Context, in *PublishListRequest, opts ...grpc.CallOption) (*PublishListResponse, error) {
	out := new(PublishListResponse)
	err := c.cc.Invoke(ctx, "/douyin.core.Feed/PublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServer is the server API for Feed service.
// All implementations must embed UnimplementedFeedServer
// for forward compatibility
type FeedServer interface {
	GetFeed(context.Context, *FeedRequest) (*FeedResponse, error)
	PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error)
	PublishList(context.Context, *PublishListRequest) (*PublishListResponse, error)
	mustEmbedUnimplementedFeedServer()
}

// UnimplementedFeedServer must be embedded to have forward compatible implementations.
type UnimplementedFeedServer struct {
}

func (UnimplementedFeedServer) GetFeed(context.Context, *FeedRequest) (*FeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (UnimplementedFeedServer) PublishVideo(context.Context, *PublishVideoRequest) (*PublishVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedFeedServer) PublishList(context.Context, *PublishListRequest) (*PublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedFeedServer) mustEmbedUnimplementedFeedServer() {}

// UnsafeFeedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedServer will
// result in compilation errors.
type UnsafeFeedServer interface {
	mustEmbedUnimplementedFeedServer()
}

func RegisterFeedServer(s grpc.ServiceRegistrar, srv FeedServer) {
	s.RegisterService(&Feed_ServiceDesc, srv)
}

func _Feed_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.core.Feed/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).GetFeed(ctx, req.(*FeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.core.Feed/PublishVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).PublishVideo(ctx, req.(*PublishVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Feed_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.core.Feed/PublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServer).PublishList(ctx, req.(*PublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Feed_ServiceDesc is the grpc.ServiceDesc for Feed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Feed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.core.Feed",
	HandlerType: (*FeedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _Feed_GetFeed_Handler,
		},
		{
			MethodName: "PublishVideo",
			Handler:    _Feed_PublishVideo_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _Feed_PublishList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feed.proto",
}
