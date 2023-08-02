// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/test/testdata/plugins/ext/ext.proto

package ext

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
	Ext_Foo_FullMethodName                 = "/ext.Ext/Foo"
	Ext_Bar_FullMethodName                 = "/ext.Ext/Bar"
	Ext_Baz_FullMethodName                 = "/ext.Ext/Baz"
	Ext_Set_FullMethodName                 = "/ext.Ext/Set"
	Ext_ServerStream_FullMethodName        = "/ext.Ext/ServerStream"
	Ext_ClientStream_FullMethodName        = "/ext.Ext/ClientStream"
	Ext_BidirectionalStream_FullMethodName = "/ext.Ext/BidirectionalStream"
)

// ExtClient is the client API for Ext service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExtClient interface {
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
	Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error)
	Baz(ctx context.Context, in *BazRequest, opts ...grpc.CallOption) (*BazRequest, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetRequest, error)
	ServerStream(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (Ext_ServerStreamClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Ext_ClientStreamClient, error)
	BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Ext_BidirectionalStreamClient, error)
}

type extClient struct {
	cc grpc.ClientConnInterface
}

func NewExtClient(cc grpc.ClientConnInterface) ExtClient {
	return &extClient{cc}
}

func (c *extClient) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, Ext_Foo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extClient) Bar(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	out := new(BarResponse)
	err := c.cc.Invoke(ctx, Ext_Bar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extClient) Baz(ctx context.Context, in *BazRequest, opts ...grpc.CallOption) (*BazRequest, error) {
	out := new(BazRequest)
	err := c.cc.Invoke(ctx, Ext_Baz_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetRequest, error) {
	out := new(SetRequest)
	err := c.cc.Invoke(ctx, Ext_Set_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *extClient) ServerStream(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (Ext_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ext_ServiceDesc.Streams[0], Ext_ServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &extServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Ext_ServerStreamClient interface {
	Recv() (*FooResponse, error)
	grpc.ClientStream
}

type extServerStreamClient struct {
	grpc.ClientStream
}

func (x *extServerStreamClient) Recv() (*FooResponse, error) {
	m := new(FooResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *extClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Ext_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ext_ServiceDesc.Streams[1], Ext_ClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &extClientStreamClient{stream}
	return x, nil
}

type Ext_ClientStreamClient interface {
	Send(*FooRequest) error
	CloseAndRecv() (*FooResponse, error)
	grpc.ClientStream
}

type extClientStreamClient struct {
	grpc.ClientStream
}

func (x *extClientStreamClient) Send(m *FooRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *extClientStreamClient) CloseAndRecv() (*FooResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FooResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *extClient) BidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Ext_BidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Ext_ServiceDesc.Streams[2], Ext_BidirectionalStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &extBidirectionalStreamClient{stream}
	return x, nil
}

type Ext_BidirectionalStreamClient interface {
	Send(*FooRequest) error
	Recv() (*FooResponse, error)
	grpc.ClientStream
}

type extBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *extBidirectionalStreamClient) Send(m *FooRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *extBidirectionalStreamClient) Recv() (*FooResponse, error) {
	m := new(FooResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExtServer is the server API for Ext service.
// All implementations must embed UnimplementedExtServer
// for forward compatibility
type ExtServer interface {
	Foo(context.Context, *FooRequest) (*FooResponse, error)
	Bar(context.Context, *BarRequest) (*BarResponse, error)
	Baz(context.Context, *BazRequest) (*BazRequest, error)
	Set(context.Context, *SetRequest) (*SetRequest, error)
	ServerStream(*FooRequest, Ext_ServerStreamServer) error
	ClientStream(Ext_ClientStreamServer) error
	BidirectionalStream(Ext_BidirectionalStreamServer) error
	mustEmbedUnimplementedExtServer()
}

// UnimplementedExtServer must be embedded to have forward compatible implementations.
type UnimplementedExtServer struct {
}

func (UnimplementedExtServer) Foo(context.Context, *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (UnimplementedExtServer) Bar(context.Context, *BarRequest) (*BarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bar not implemented")
}
func (UnimplementedExtServer) Baz(context.Context, *BazRequest) (*BazRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Baz not implemented")
}
func (UnimplementedExtServer) Set(context.Context, *SetRequest) (*SetRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedExtServer) ServerStream(*FooRequest, Ext_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedExtServer) ClientStream(Ext_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedExtServer) BidirectionalStream(Ext_BidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStream not implemented")
}
func (UnimplementedExtServer) mustEmbedUnimplementedExtServer() {}

// UnsafeExtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExtServer will
// result in compilation errors.
type UnsafeExtServer interface {
	mustEmbedUnimplementedExtServer()
}

func RegisterExtServer(s grpc.ServiceRegistrar, srv ExtServer) {
	s.RegisterService(&Ext_ServiceDesc, srv)
}

func _Ext_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ext_Foo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtServer).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ext_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ext_Bar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtServer).Bar(ctx, req.(*BarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ext_Baz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BazRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtServer).Baz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ext_Baz_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtServer).Baz(ctx, req.(*BazRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ext_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExtServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ext_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExtServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ext_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FooRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExtServer).ServerStream(m, &extServerStreamServer{stream})
}

type Ext_ServerStreamServer interface {
	Send(*FooResponse) error
	grpc.ServerStream
}

type extServerStreamServer struct {
	grpc.ServerStream
}

func (x *extServerStreamServer) Send(m *FooResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Ext_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExtServer).ClientStream(&extClientStreamServer{stream})
}

type Ext_ClientStreamServer interface {
	SendAndClose(*FooResponse) error
	Recv() (*FooRequest, error)
	grpc.ServerStream
}

type extClientStreamServer struct {
	grpc.ServerStream
}

func (x *extClientStreamServer) SendAndClose(m *FooResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *extClientStreamServer) Recv() (*FooRequest, error) {
	m := new(FooRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ext_BidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExtServer).BidirectionalStream(&extBidirectionalStreamServer{stream})
}

type Ext_BidirectionalStreamServer interface {
	Send(*FooResponse) error
	Recv() (*FooRequest, error)
	grpc.ServerStream
}

type extBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *extBidirectionalStreamServer) Send(m *FooResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *extBidirectionalStreamServer) Recv() (*FooRequest, error) {
	m := new(FooRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Ext_ServiceDesc is the grpc.ServiceDesc for Ext service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ext_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ext.Ext",
	HandlerType: (*ExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _Ext_Foo_Handler,
		},
		{
			MethodName: "Bar",
			Handler:    _Ext_Bar_Handler,
		},
		{
			MethodName: "Baz",
			Handler:    _Ext_Baz_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Ext_Set_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Ext_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _Ext_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStream",
			Handler:       _Ext_BidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/rancher/opni/pkg/test/testdata/plugins/ext/ext.proto",
}

const (
	Ext2_Foo_FullMethodName = "/ext.Ext2/Foo"
)

// Ext2Client is the client API for Ext2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Ext2Client interface {
	Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error)
}

type ext2Client struct {
	cc grpc.ClientConnInterface
}

func NewExt2Client(cc grpc.ClientConnInterface) Ext2Client {
	return &ext2Client{cc}
}

func (c *ext2Client) Foo(ctx context.Context, in *FooRequest, opts ...grpc.CallOption) (*FooResponse, error) {
	out := new(FooResponse)
	err := c.cc.Invoke(ctx, Ext2_Foo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Ext2Server is the server API for Ext2 service.
// All implementations must embed UnimplementedExt2Server
// for forward compatibility
type Ext2Server interface {
	Foo(context.Context, *FooRequest) (*FooResponse, error)
	mustEmbedUnimplementedExt2Server()
}

// UnimplementedExt2Server must be embedded to have forward compatible implementations.
type UnimplementedExt2Server struct {
}

func (UnimplementedExt2Server) Foo(context.Context, *FooRequest) (*FooResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Foo not implemented")
}
func (UnimplementedExt2Server) mustEmbedUnimplementedExt2Server() {}

// UnsafeExt2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Ext2Server will
// result in compilation errors.
type UnsafeExt2Server interface {
	mustEmbedUnimplementedExt2Server()
}

func RegisterExt2Server(s grpc.ServiceRegistrar, srv Ext2Server) {
	s.RegisterService(&Ext2_ServiceDesc, srv)
}

func _Ext2_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Ext2Server).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ext2_Foo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Ext2Server).Foo(ctx, req.(*FooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ext2_ServiceDesc is the grpc.ServiceDesc for Ext2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ext2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ext.Ext2",
	HandlerType: (*Ext2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _Ext2_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/test/testdata/plugins/ext/ext.proto",
}
