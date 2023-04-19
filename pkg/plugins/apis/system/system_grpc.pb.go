// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/pkg/plugins/apis/system/system.proto

package system

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	System_UseManagementAPI_FullMethodName     = "/system.System/UseManagementAPI"
	System_UseNodeManagerClient_FullMethodName = "/system.System/UseNodeManagerClient"
	System_UseKeyValueStore_FullMethodName     = "/system.System/UseKeyValueStore"
	System_UseAPIExtensions_FullMethodName     = "/system.System/UseAPIExtensions"
	System_UseCachingProvider_FullMethodName   = "/system.System/UseCachingProvider"
)

// SystemClient is the client API for System service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemClient interface {
	UseManagementAPI(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UseNodeManagerClient(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UseKeyValueStore(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UseAPIExtensions(ctx context.Context, in *DialAddress, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UseCachingProvider(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type systemClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemClient(cc grpc.ClientConnInterface) SystemClient {
	return &systemClient{cc}
}

func (c *systemClient) UseManagementAPI(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, System_UseManagementAPI_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UseNodeManagerClient(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, System_UseNodeManagerClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UseKeyValueStore(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, System_UseKeyValueStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UseAPIExtensions(ctx context.Context, in *DialAddress, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, System_UseAPIExtensions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemClient) UseCachingProvider(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, System_UseCachingProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServer is the server API for System service.
// All implementations must embed UnimplementedSystemServer
// for forward compatibility
type SystemServer interface {
	UseManagementAPI(context.Context, *BrokerID) (*emptypb.Empty, error)
	UseNodeManagerClient(context.Context, *BrokerID) (*emptypb.Empty, error)
	UseKeyValueStore(context.Context, *BrokerID) (*emptypb.Empty, error)
	UseAPIExtensions(context.Context, *DialAddress) (*emptypb.Empty, error)
	UseCachingProvider(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedSystemServer()
}

// UnimplementedSystemServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServer struct {
}

func (UnimplementedSystemServer) UseManagementAPI(context.Context, *BrokerID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseManagementAPI not implemented")
}
func (UnimplementedSystemServer) UseNodeManagerClient(context.Context, *BrokerID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseNodeManagerClient not implemented")
}
func (UnimplementedSystemServer) UseKeyValueStore(context.Context, *BrokerID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseKeyValueStore not implemented")
}
func (UnimplementedSystemServer) UseAPIExtensions(context.Context, *DialAddress) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseAPIExtensions not implemented")
}
func (UnimplementedSystemServer) UseCachingProvider(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseCachingProvider not implemented")
}
func (UnimplementedSystemServer) mustEmbedUnimplementedSystemServer() {}

// UnsafeSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServer will
// result in compilation errors.
type UnsafeSystemServer interface {
	mustEmbedUnimplementedSystemServer()
}

func RegisterSystemServer(s grpc.ServiceRegistrar, srv SystemServer) {
	s.RegisterService(&System_ServiceDesc, srv)
}

func _System_UseManagementAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UseManagementAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UseManagementAPI_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UseManagementAPI(ctx, req.(*BrokerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UseNodeManagerClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UseNodeManagerClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UseNodeManagerClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UseNodeManagerClient(ctx, req.(*BrokerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UseKeyValueStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UseKeyValueStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UseKeyValueStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UseKeyValueStore(ctx, req.(*BrokerID))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UseAPIExtensions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DialAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UseAPIExtensions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UseAPIExtensions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UseAPIExtensions(ctx, req.(*DialAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _System_UseCachingProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UseCachingProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: System_UseCachingProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UseCachingProvider(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// System_ServiceDesc is the grpc.ServiceDesc for System service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var System_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.System",
	HandlerType: (*SystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UseManagementAPI",
			Handler:    _System_UseManagementAPI_Handler,
		},
		{
			MethodName: "UseNodeManagerClient",
			Handler:    _System_UseNodeManagerClient_Handler,
		},
		{
			MethodName: "UseKeyValueStore",
			Handler:    _System_UseKeyValueStore_Handler,
		},
		{
			MethodName: "UseAPIExtensions",
			Handler:    _System_UseAPIExtensions_Handler,
		},
		{
			MethodName: "UseCachingProvider",
			Handler:    _System_UseCachingProvider_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/apis/system/system.proto",
}

const (
	KeyValueStore_Put_FullMethodName      = "/system.KeyValueStore/Put"
	KeyValueStore_Get_FullMethodName      = "/system.KeyValueStore/Get"
	KeyValueStore_Delete_FullMethodName   = "/system.KeyValueStore/Delete"
	KeyValueStore_ListKeys_FullMethodName = "/system.KeyValueStore/ListKeys"
)

// KeyValueStoreClient is the client API for KeyValueStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueStoreClient interface {
	Put(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListKeys(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyList, error)
}

type keyValueStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueStoreClient(cc grpc.ClientConnInterface) KeyValueStoreClient {
	return &keyValueStoreClient{cc}
}

func (c *keyValueStoreClient) Put(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KeyValueStore_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, KeyValueStore_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, KeyValueStore_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueStoreClient) ListKeys(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyList, error) {
	out := new(KeyList)
	err := c.cc.Invoke(ctx, KeyValueStore_ListKeys_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueStoreServer is the server API for KeyValueStore service.
// All implementations must embed UnimplementedKeyValueStoreServer
// for forward compatibility
type KeyValueStoreServer interface {
	Put(context.Context, *KeyValue) (*emptypb.Empty, error)
	Get(context.Context, *Key) (*Value, error)
	Delete(context.Context, *Key) (*emptypb.Empty, error)
	ListKeys(context.Context, *Key) (*KeyList, error)
	mustEmbedUnimplementedKeyValueStoreServer()
}

// UnimplementedKeyValueStoreServer must be embedded to have forward compatible implementations.
type UnimplementedKeyValueStoreServer struct {
}

func (UnimplementedKeyValueStoreServer) Put(context.Context, *KeyValue) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedKeyValueStoreServer) Get(context.Context, *Key) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKeyValueStoreServer) Delete(context.Context, *Key) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKeyValueStoreServer) ListKeys(context.Context, *Key) (*KeyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeys not implemented")
}
func (UnimplementedKeyValueStoreServer) mustEmbedUnimplementedKeyValueStoreServer() {}

// UnsafeKeyValueStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueStoreServer will
// result in compilation errors.
type UnsafeKeyValueStoreServer interface {
	mustEmbedUnimplementedKeyValueStoreServer()
}

func RegisterKeyValueStoreServer(s grpc.ServiceRegistrar, srv KeyValueStoreServer) {
	s.RegisterService(&KeyValueStore_ServiceDesc, srv)
}

func _KeyValueStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStore_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).Put(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStore_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueStore_ListKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueStoreServer).ListKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyValueStore_ListKeys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueStoreServer).ListKeys(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValueStore_ServiceDesc is the grpc.ServiceDesc for KeyValueStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValueStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.KeyValueStore",
	HandlerType: (*KeyValueStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _KeyValueStore_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KeyValueStore_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KeyValueStore_Delete_Handler,
		},
		{
			MethodName: "ListKeys",
			Handler:    _KeyValueStore_ListKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/pkg/plugins/apis/system/system.proto",
}
