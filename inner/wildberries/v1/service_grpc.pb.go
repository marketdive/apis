// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wildberriespb

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

// WildberriesClient is the client API for Wildberries service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WildberriesClient interface {
	GetCategoryList(ctx context.Context, in *GetCategoryListRequest, opts ...grpc.CallOption) (*GetCategoryListResponse, error)
	GetProductsList(ctx context.Context, in *GetProductsListRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error)
	GetSupplier(ctx context.Context, in *GetSupplierInfoRequest, opts ...grpc.CallOption) (*GetSupplierInfoResponse, error)
	GetBrand(ctx context.Context, in *GetBrandInfoRequest, opts ...grpc.CallOption) (*GetBrandInfoResponse, error)
	GetProduct(ctx context.Context, in *GetProductInfoRequest, opts ...grpc.CallOption) (*GetProductInfoResponse, error)
}

type wildberriesClient struct {
	cc grpc.ClientConnInterface
}

func NewWildberriesClient(cc grpc.ClientConnInterface) WildberriesClient {
	return &wildberriesClient{cc}
}

func (c *wildberriesClient) GetCategoryList(ctx context.Context, in *GetCategoryListRequest, opts ...grpc.CallOption) (*GetCategoryListResponse, error) {
	out := new(GetCategoryListResponse)
	err := c.cc.Invoke(ctx, "/marketdive.inner.wildberries.v1.Wildberries/GetCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wildberriesClient) GetProductsList(ctx context.Context, in *GetProductsListRequest, opts ...grpc.CallOption) (*GetProductsListResponse, error) {
	out := new(GetProductsListResponse)
	err := c.cc.Invoke(ctx, "/marketdive.inner.wildberries.v1.Wildberries/GetProductsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wildberriesClient) GetSupplier(ctx context.Context, in *GetSupplierInfoRequest, opts ...grpc.CallOption) (*GetSupplierInfoResponse, error) {
	out := new(GetSupplierInfoResponse)
	err := c.cc.Invoke(ctx, "/marketdive.inner.wildberries.v1.Wildberries/GetSupplier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wildberriesClient) GetBrand(ctx context.Context, in *GetBrandInfoRequest, opts ...grpc.CallOption) (*GetBrandInfoResponse, error) {
	out := new(GetBrandInfoResponse)
	err := c.cc.Invoke(ctx, "/marketdive.inner.wildberries.v1.Wildberries/GetBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wildberriesClient) GetProduct(ctx context.Context, in *GetProductInfoRequest, opts ...grpc.CallOption) (*GetProductInfoResponse, error) {
	out := new(GetProductInfoResponse)
	err := c.cc.Invoke(ctx, "/marketdive.inner.wildberries.v1.Wildberries/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WildberriesServer is the server API for Wildberries service.
// All implementations must embed UnimplementedWildberriesServer
// for forward compatibility
type WildberriesServer interface {
	GetCategoryList(context.Context, *GetCategoryListRequest) (*GetCategoryListResponse, error)
	GetProductsList(context.Context, *GetProductsListRequest) (*GetProductsListResponse, error)
	GetSupplier(context.Context, *GetSupplierInfoRequest) (*GetSupplierInfoResponse, error)
	GetBrand(context.Context, *GetBrandInfoRequest) (*GetBrandInfoResponse, error)
	GetProduct(context.Context, *GetProductInfoRequest) (*GetProductInfoResponse, error)
	mustEmbedUnimplementedWildberriesServer()
}

// UnimplementedWildberriesServer must be embedded to have forward compatible implementations.
type UnimplementedWildberriesServer struct {
}

func (UnimplementedWildberriesServer) GetCategoryList(context.Context, *GetCategoryListRequest) (*GetCategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryList not implemented")
}
func (UnimplementedWildberriesServer) GetProductsList(context.Context, *GetProductsListRequest) (*GetProductsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsList not implemented")
}
func (UnimplementedWildberriesServer) GetSupplier(context.Context, *GetSupplierInfoRequest) (*GetSupplierInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupplier not implemented")
}
func (UnimplementedWildberriesServer) GetBrand(context.Context, *GetBrandInfoRequest) (*GetBrandInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrand not implemented")
}
func (UnimplementedWildberriesServer) GetProduct(context.Context, *GetProductInfoRequest) (*GetProductInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedWildberriesServer) mustEmbedUnimplementedWildberriesServer() {}

// UnsafeWildberriesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WildberriesServer will
// result in compilation errors.
type UnsafeWildberriesServer interface {
	mustEmbedUnimplementedWildberriesServer()
}

func RegisterWildberriesServer(s grpc.ServiceRegistrar, srv WildberriesServer) {
	s.RegisterService(&Wildberries_ServiceDesc, srv)
}

func _Wildberries_GetCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WildberriesServer).GetCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdive.inner.wildberries.v1.Wildberries/GetCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WildberriesServer).GetCategoryList(ctx, req.(*GetCategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wildberries_GetProductsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WildberriesServer).GetProductsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdive.inner.wildberries.v1.Wildberries/GetProductsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WildberriesServer).GetProductsList(ctx, req.(*GetProductsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wildberries_GetSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupplierInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WildberriesServer).GetSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdive.inner.wildberries.v1.Wildberries/GetSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WildberriesServer).GetSupplier(ctx, req.(*GetSupplierInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wildberries_GetBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBrandInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WildberriesServer).GetBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdive.inner.wildberries.v1.Wildberries/GetBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WildberriesServer).GetBrand(ctx, req.(*GetBrandInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wildberries_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WildberriesServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/marketdive.inner.wildberries.v1.Wildberries/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WildberriesServer).GetProduct(ctx, req.(*GetProductInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Wildberries_ServiceDesc is the grpc.ServiceDesc for Wildberries service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wildberries_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "marketdive.inner.wildberries.v1.Wildberries",
	HandlerType: (*WildberriesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCategoryList",
			Handler:    _Wildberries_GetCategoryList_Handler,
		},
		{
			MethodName: "GetProductsList",
			Handler:    _Wildberries_GetProductsList_Handler,
		},
		{
			MethodName: "GetSupplier",
			Handler:    _Wildberries_GetSupplier_Handler,
		},
		{
			MethodName: "GetBrand",
			Handler:    _Wildberries_GetBrand_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Wildberries_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inner/wildberries/v1/service.proto",
}
