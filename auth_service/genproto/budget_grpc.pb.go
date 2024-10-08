// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: budget.proto

package genproto

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
	BudgetService_CreateBudget_FullMethodName = "/submodule.BudgetService/CreateBudget"
	BudgetService_UpdateBudget_FullMethodName = "/submodule.BudgetService/UpdateBudget"
	BudgetService_DeleteBudget_FullMethodName = "/submodule.BudgetService/DeleteBudget"
	BudgetService_GetBudget_FullMethodName    = "/submodule.BudgetService/GetBudget"
	BudgetService_ListBudgets_FullMethodName  = "/submodule.BudgetService/ListBudgets"
)

// BudgetServiceClient is the client API for BudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BudgetServiceClient interface {
	CreateBudget(ctx context.Context, in *BudgetCreate, opts ...grpc.CallOption) (*Void, error)
	UpdateBudget(ctx context.Context, in *BudgetUpdate, opts ...grpc.CallOption) (*Void, error)
	DeleteBudget(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	GetBudget(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BudgetGet, error)
	ListBudgets(ctx context.Context, in *BudgetFilter, opts ...grpc.CallOption) (*BudgetList, error)
}

type budgetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBudgetServiceClient(cc grpc.ClientConnInterface) BudgetServiceClient {
	return &budgetServiceClient{cc}
}

func (c *budgetServiceClient) CreateBudget(ctx context.Context, in *BudgetCreate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, BudgetService_CreateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) UpdateBudget(ctx context.Context, in *BudgetUpdate, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, BudgetService_UpdateBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) DeleteBudget(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Void)
	err := c.cc.Invoke(ctx, BudgetService_DeleteBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) GetBudget(ctx context.Context, in *ById, opts ...grpc.CallOption) (*BudgetGet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BudgetGet)
	err := c.cc.Invoke(ctx, BudgetService_GetBudget_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *budgetServiceClient) ListBudgets(ctx context.Context, in *BudgetFilter, opts ...grpc.CallOption) (*BudgetList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BudgetList)
	err := c.cc.Invoke(ctx, BudgetService_ListBudgets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BudgetServiceServer is the server API for BudgetService service.
// All implementations must embed UnimplementedBudgetServiceServer
// for forward compatibility
type BudgetServiceServer interface {
	CreateBudget(context.Context, *BudgetCreate) (*Void, error)
	UpdateBudget(context.Context, *BudgetUpdate) (*Void, error)
	DeleteBudget(context.Context, *ById) (*Void, error)
	GetBudget(context.Context, *ById) (*BudgetGet, error)
	ListBudgets(context.Context, *BudgetFilter) (*BudgetList, error)
	mustEmbedUnimplementedBudgetServiceServer()
}

// UnimplementedBudgetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBudgetServiceServer struct {
}

func (UnimplementedBudgetServiceServer) CreateBudget(context.Context, *BudgetCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBudget not implemented")
}
func (UnimplementedBudgetServiceServer) UpdateBudget(context.Context, *BudgetUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBudget not implemented")
}
func (UnimplementedBudgetServiceServer) DeleteBudget(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBudget not implemented")
}
func (UnimplementedBudgetServiceServer) GetBudget(context.Context, *ById) (*BudgetGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudget not implemented")
}
func (UnimplementedBudgetServiceServer) ListBudgets(context.Context, *BudgetFilter) (*BudgetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBudgets not implemented")
}
func (UnimplementedBudgetServiceServer) mustEmbedUnimplementedBudgetServiceServer() {}

// UnsafeBudgetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BudgetServiceServer will
// result in compilation errors.
type UnsafeBudgetServiceServer interface {
	mustEmbedUnimplementedBudgetServiceServer()
}

func RegisterBudgetServiceServer(s grpc.ServiceRegistrar, srv BudgetServiceServer) {
	s.RegisterService(&BudgetService_ServiceDesc, srv)
}

func _BudgetService_CreateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).CreateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_CreateBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).CreateBudget(ctx, req.(*BudgetCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_UpdateBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).UpdateBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_UpdateBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).UpdateBudget(ctx, req.(*BudgetUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_DeleteBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).DeleteBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_DeleteBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).DeleteBudget(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_GetBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).GetBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_GetBudget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).GetBudget(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _BudgetService_ListBudgets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BudgetFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BudgetServiceServer).ListBudgets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BudgetService_ListBudgets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BudgetServiceServer).ListBudgets(ctx, req.(*BudgetFilter))
	}
	return interceptor(ctx, in, info, handler)
}

// BudgetService_ServiceDesc is the grpc.ServiceDesc for BudgetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BudgetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "submodule.BudgetService",
	HandlerType: (*BudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBudget",
			Handler:    _BudgetService_CreateBudget_Handler,
		},
		{
			MethodName: "UpdateBudget",
			Handler:    _BudgetService_UpdateBudget_Handler,
		},
		{
			MethodName: "DeleteBudget",
			Handler:    _BudgetService_DeleteBudget_Handler,
		},
		{
			MethodName: "GetBudget",
			Handler:    _BudgetService_GetBudget_Handler,
		},
		{
			MethodName: "ListBudgets",
			Handler:    _BudgetService_ListBudgets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "budget.proto",
}
