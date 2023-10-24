// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: octopus.proto

package octopus

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

// OctopusServiceClient is the client API for OctopusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OctopusServiceClient interface {
	GetTopology(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*TopologyResponse, error)
}

type octopusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOctopusServiceClient(cc grpc.ClientConnInterface) OctopusServiceClient {
	return &octopusServiceClient{cc}
}

func (c *octopusServiceClient) GetTopology(ctx context.Context, in *TopologyRequest, opts ...grpc.CallOption) (*TopologyResponse, error) {
	out := new(TopologyResponse)
	err := c.cc.Invoke(ctx, "/cloudflare.net.octopus.OctopusService/GetTopology", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OctopusServiceServer is the server API for OctopusService service.
// All implementations should embed UnimplementedOctopusServiceServer
// for forward compatibility
type OctopusServiceServer interface {
	GetTopology(context.Context, *TopologyRequest) (*TopologyResponse, error)
}

// UnimplementedOctopusServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOctopusServiceServer struct {
}

func (UnimplementedOctopusServiceServer) GetTopology(context.Context, *TopologyRequest) (*TopologyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopology not implemented")
}

// UnsafeOctopusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OctopusServiceServer will
// result in compilation errors.
type UnsafeOctopusServiceServer interface {
	mustEmbedUnimplementedOctopusServiceServer()
}

func RegisterOctopusServiceServer(s grpc.ServiceRegistrar, srv OctopusServiceServer) {
	s.RegisterService(&OctopusService_ServiceDesc, srv)
}

func _OctopusService_GetTopology_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopologyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OctopusServiceServer).GetTopology(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudflare.net.octopus.OctopusService/GetTopology",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OctopusServiceServer).GetTopology(ctx, req.(*TopologyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OctopusService_ServiceDesc is the grpc.ServiceDesc for OctopusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OctopusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cloudflare.net.octopus.OctopusService",
	HandlerType: (*OctopusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopology",
			Handler:    _OctopusService_GetTopology_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "octopus.proto",
}
