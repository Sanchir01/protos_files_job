// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: size/size.proto

package sandjmav1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Sizes_AllSizes_FullMethodName = "/size.Sizes/AllSizes"
)

// SizesClient is the client API for Sizes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SizesClient interface {
	AllSizes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCategoryResponse, error)
}

type sizesClient struct {
	cc grpc.ClientConnInterface
}

func NewSizesClient(cc grpc.ClientConnInterface) SizesClient {
	return &sizesClient{cc}
}

func (c *sizesClient) AllSizes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllCategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCategoryResponse)
	err := c.cc.Invoke(ctx, Sizes_AllSizes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SizesServer is the server API for Sizes service.
// All implementations must embed UnimplementedSizesServer
// for forward compatibility.
type SizesServer interface {
	AllSizes(context.Context, *emptypb.Empty) (*GetAllCategoryResponse, error)
	mustEmbedUnimplementedSizesServer()
}

// UnimplementedSizesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSizesServer struct{}

func (UnimplementedSizesServer) AllSizes(context.Context, *emptypb.Empty) (*GetAllCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllSizes not implemented")
}
func (UnimplementedSizesServer) mustEmbedUnimplementedSizesServer() {}
func (UnimplementedSizesServer) testEmbeddedByValue()               {}

// UnsafeSizesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SizesServer will
// result in compilation errors.
type UnsafeSizesServer interface {
	mustEmbedUnimplementedSizesServer()
}

func RegisterSizesServer(s grpc.ServiceRegistrar, srv SizesServer) {
	// If the following call pancis, it indicates UnimplementedSizesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Sizes_ServiceDesc, srv)
}

func _Sizes_AllSizes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SizesServer).AllSizes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sizes_AllSizes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SizesServer).AllSizes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Sizes_ServiceDesc is the grpc.ServiceDesc for Sizes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sizes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "size.Sizes",
	HandlerType: (*SizesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllSizes",
			Handler:    _Sizes_AllSizes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "size/size.proto",
}
