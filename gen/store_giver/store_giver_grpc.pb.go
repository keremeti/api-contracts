// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: store_giver.proto

package store_giver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	StoreGiver_GetByPurchaser_FullMethodName = "/store_giver.StoreGiver/GetByPurchaser"
)

// StoreGiverClient is the client API for StoreGiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreGiverClient interface {
	GetByPurchaser(ctx context.Context, in *PurchaserRequest, opts ...grpc.CallOption) (*BriefStoreResponse, error)
}

type storeGiverClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreGiverClient(cc grpc.ClientConnInterface) StoreGiverClient {
	return &storeGiverClient{cc}
}

func (c *storeGiverClient) GetByPurchaser(ctx context.Context, in *PurchaserRequest, opts ...grpc.CallOption) (*BriefStoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BriefStoreResponse)
	err := c.cc.Invoke(ctx, StoreGiver_GetByPurchaser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreGiverServer is the server API for StoreGiver service.
// All implementations must embed UnimplementedStoreGiverServer
// for forward compatibility
type StoreGiverServer interface {
	GetByPurchaser(context.Context, *PurchaserRequest) (*BriefStoreResponse, error)
	mustEmbedUnimplementedStoreGiverServer()
}

// UnimplementedStoreGiverServer must be embedded to have forward compatible implementations.
type UnimplementedStoreGiverServer struct {
}

func (UnimplementedStoreGiverServer) GetByPurchaser(context.Context, *PurchaserRequest) (*BriefStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPurchaser not implemented")
}
func (UnimplementedStoreGiverServer) mustEmbedUnimplementedStoreGiverServer() {}

// UnsafeStoreGiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreGiverServer will
// result in compilation errors.
type UnsafeStoreGiverServer interface {
	mustEmbedUnimplementedStoreGiverServer()
}

func RegisterStoreGiverServer(s grpc.ServiceRegistrar, srv StoreGiverServer) {
	s.RegisterService(&StoreGiver_ServiceDesc, srv)
}

func _StoreGiver_GetByPurchaser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreGiverServer).GetByPurchaser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreGiver_GetByPurchaser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreGiverServer).GetByPurchaser(ctx, req.(*PurchaserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreGiver_ServiceDesc is the grpc.ServiceDesc for StoreGiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreGiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store_giver.StoreGiver",
	HandlerType: (*StoreGiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByPurchaser",
			Handler:    _StoreGiver_GetByPurchaser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store_giver.proto",
}
