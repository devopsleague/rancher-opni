// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - ragu               v1.0.0
// source: github.com/rancher/opni/plugins/topology/apis/representation/representation.proto

package representation

import (
	context "context"
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TopologyRepresentation_GetGraph_FullMethodName    = "/representation.TopologyRepresentation/GetGraph"
	TopologyRepresentation_RenderGraph_FullMethodName = "/representation.TopologyRepresentation/RenderGraph"
)

// TopologyRepresentationClient is the client API for TopologyRepresentation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopologyRepresentationClient interface {
	// opni internal use
	GetGraph(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*TopologyGraph, error)
	// cluster id  --> kubernetes graph SVG
	RenderGraph(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*DOTRepresentation, error)
}

type topologyRepresentationClient struct {
	cc grpc.ClientConnInterface
}

func NewTopologyRepresentationClient(cc grpc.ClientConnInterface) TopologyRepresentationClient {
	return &topologyRepresentationClient{cc}
}

func (c *topologyRepresentationClient) GetGraph(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*TopologyGraph, error) {
	out := new(TopologyGraph)
	err := c.cc.Invoke(ctx, TopologyRepresentation_GetGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyRepresentationClient) RenderGraph(ctx context.Context, in *v1.Reference, opts ...grpc.CallOption) (*DOTRepresentation, error) {
	out := new(DOTRepresentation)
	err := c.cc.Invoke(ctx, TopologyRepresentation_RenderGraph_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopologyRepresentationServer is the server API for TopologyRepresentation service.
// All implementations must embed UnimplementedTopologyRepresentationServer
// for forward compatibility
type TopologyRepresentationServer interface {
	// opni internal use
	GetGraph(context.Context, *v1.Reference) (*TopologyGraph, error)
	// cluster id  --> kubernetes graph SVG
	RenderGraph(context.Context, *v1.Reference) (*DOTRepresentation, error)
	mustEmbedUnimplementedTopologyRepresentationServer()
}

// UnimplementedTopologyRepresentationServer must be embedded to have forward compatible implementations.
type UnimplementedTopologyRepresentationServer struct {
}

func (UnimplementedTopologyRepresentationServer) GetGraph(context.Context, *v1.Reference) (*TopologyGraph, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}
func (UnimplementedTopologyRepresentationServer) RenderGraph(context.Context, *v1.Reference) (*DOTRepresentation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenderGraph not implemented")
}
func (UnimplementedTopologyRepresentationServer) mustEmbedUnimplementedTopologyRepresentationServer() {
}

// UnsafeTopologyRepresentationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopologyRepresentationServer will
// result in compilation errors.
type UnsafeTopologyRepresentationServer interface {
	mustEmbedUnimplementedTopologyRepresentationServer()
}

func RegisterTopologyRepresentationServer(s grpc.ServiceRegistrar, srv TopologyRepresentationServer) {
	s.RegisterService(&TopologyRepresentation_ServiceDesc, srv)
}

func _TopologyRepresentation_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyRepresentationServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyRepresentation_GetGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyRepresentationServer).GetGraph(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyRepresentation_RenderGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyRepresentationServer).RenderGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyRepresentation_RenderGraph_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyRepresentationServer).RenderGraph(ctx, req.(*v1.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

// TopologyRepresentation_ServiceDesc is the grpc.ServiceDesc for TopologyRepresentation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopologyRepresentation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "representation.TopologyRepresentation",
	HandlerType: (*TopologyRepresentationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGraph",
			Handler:    _TopologyRepresentation_GetGraph_Handler,
		},
		{
			MethodName: "RenderGraph",
			Handler:    _TopologyRepresentation_RenderGraph_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/rancher/opni/plugins/topology/apis/representation/representation.proto",
}