// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: vod/v1/vod.proto

package v1

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

// VODServiceClient is the client API for VODService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VODServiceClient interface {
	ListVODs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListVODsResponse, error)
}

type vODServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVODServiceClient(cc grpc.ClientConnInterface) VODServiceClient {
	return &vODServiceClient{cc}
}

func (c *vODServiceClient) ListVODs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListVODsResponse, error) {
	out := new(ListVODsResponse)
	err := c.cc.Invoke(ctx, "/vod.v1.VODService/ListVODs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VODServiceServer is the server API for VODService service.
// All implementations must embed UnimplementedVODServiceServer
// for forward compatibility
type VODServiceServer interface {
	ListVODs(context.Context, *emptypb.Empty) (*ListVODsResponse, error)
	mustEmbedUnimplementedVODServiceServer()
}

// UnimplementedVODServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVODServiceServer struct {
}

func (UnimplementedVODServiceServer) ListVODs(context.Context, *emptypb.Empty) (*ListVODsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVODs not implemented")
}
func (UnimplementedVODServiceServer) mustEmbedUnimplementedVODServiceServer() {}

// UnsafeVODServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VODServiceServer will
// result in compilation errors.
type UnsafeVODServiceServer interface {
	mustEmbedUnimplementedVODServiceServer()
}

func RegisterVODServiceServer(s grpc.ServiceRegistrar, srv VODServiceServer) {
	s.RegisterService(&VODService_ServiceDesc, srv)
}

func _VODService_ListVODs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VODServiceServer).ListVODs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vod.v1.VODService/ListVODs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VODServiceServer).ListVODs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VODService_ServiceDesc is the grpc.ServiceDesc for VODService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VODService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vod.v1.VODService",
	HandlerType: (*VODServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVODs",
			Handler:    _VODService_ListVODs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vod/v1/vod.proto",
}
