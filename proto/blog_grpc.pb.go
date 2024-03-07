// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/blog.proto

package blogging_platform

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

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogClient interface {
	GetBlog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*BlogData, error)
	GetBlogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Blog_GetBlogsClient, error)
	CreateBlog(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*BlogData, error)
	UpdateBlog(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*UpdateStatus, error)
	DeleteBlog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) GetBlog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*BlogData, error) {
	out := new(BlogData)
	err := c.cc.Invoke(ctx, "/proto.Blog/GetBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) GetBlogs(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Blog_GetBlogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Blog_ServiceDesc.Streams[0], "/proto.Blog/GetBlogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogGetBlogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Blog_GetBlogsClient interface {
	Recv() (*BlogData, error)
	grpc.ClientStream
}

type blogGetBlogsClient struct {
	grpc.ClientStream
}

func (x *blogGetBlogsClient) Recv() (*BlogData, error) {
	m := new(BlogData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *blogClient) CreateBlog(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*BlogData, error) {
	out := new(BlogData)
	err := c.cc.Invoke(ctx, "/proto.Blog/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) UpdateBlog(ctx context.Context, in *BlogData, opts ...grpc.CallOption) (*UpdateStatus, error) {
	out := new(UpdateStatus)
	err := c.cc.Invoke(ctx, "/proto.Blog/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogClient) DeleteBlog(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.Blog/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServer is the server API for Blog service.
// All implementations must embed UnimplementedBlogServer
// for forward compatibility
type BlogServer interface {
	GetBlog(context.Context, *ID) (*BlogData, error)
	GetBlogs(*Empty, Blog_GetBlogsServer) error
	CreateBlog(context.Context, *BlogData) (*BlogData, error)
	UpdateBlog(context.Context, *BlogData) (*UpdateStatus, error)
	DeleteBlog(context.Context, *ID) (*Status, error)
	mustEmbedUnimplementedBlogServer()
}

// UnimplementedBlogServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (UnimplementedBlogServer) GetBlog(context.Context, *ID) (*BlogData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlog not implemented")
}
func (UnimplementedBlogServer) GetBlogs(*Empty, Blog_GetBlogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlogs not implemented")
}
func (UnimplementedBlogServer) CreateBlog(context.Context, *BlogData) (*BlogData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (UnimplementedBlogServer) UpdateBlog(context.Context, *BlogData) (*UpdateStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (UnimplementedBlogServer) DeleteBlog(context.Context, *ID) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (UnimplementedBlogServer) mustEmbedUnimplementedBlogServer() {}

// UnsafeBlogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServer will
// result in compilation errors.
type UnsafeBlogServer interface {
	mustEmbedUnimplementedBlogServer()
}

func RegisterBlogServer(s grpc.ServiceRegistrar, srv BlogServer) {
	s.RegisterService(&Blog_ServiceDesc, srv)
}

func _Blog_GetBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).GetBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blog/GetBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).GetBlog(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_GetBlogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServer).GetBlogs(m, &blogGetBlogsServer{stream})
}

type Blog_GetBlogsServer interface {
	Send(*BlogData) error
	grpc.ServerStream
}

type blogGetBlogsServer struct {
	grpc.ServerStream
}

func (x *blogGetBlogsServer) Send(m *BlogData) error {
	return x.ServerStream.SendMsg(m)
}

func _Blog_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blog/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).CreateBlog(ctx, req.(*BlogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlogData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blog/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).UpdateBlog(ctx, req.(*BlogData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blog_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Blog/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServer).DeleteBlog(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

// Blog_ServiceDesc is the grpc.ServiceDesc for Blog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlog",
			Handler:    _Blog_GetBlog_Handler,
		},
		{
			MethodName: "CreateBlog",
			Handler:    _Blog_CreateBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _Blog_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _Blog_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlogs",
			Handler:       _Blog_GetBlogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/blog.proto",
}
