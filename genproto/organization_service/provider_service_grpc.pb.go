// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package organization_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProviderServiceClient is the client API for ProviderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServiceClient interface {
	Create(ctx context.Context, in *CreateProvider, opts ...grpc.CallOption) (*Provider, error)
	GetByID(ctx context.Context, in *ProviderPK, opts ...grpc.CallOption) (*Provider, error)
	GetList(ctx context.Context, in *GetListProviderRequest, opts ...grpc.CallOption) (*GetListProviderResponse, error)
	Update(ctx context.Context, in *UpdateProvider, opts ...grpc.CallOption) (*Provider, error)
	UpdatePatch(ctx context.Context, in *UpdatePatchProvider, opts ...grpc.CallOption) (*Provider, error)
	Delete(ctx context.Context, in *ProviderPK, opts ...grpc.CallOption) (*empty.Empty, error)
}

type providerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServiceClient(cc grpc.ClientConnInterface) ProviderServiceClient {
	return &providerServiceClient{cc}
}

func (c *providerServiceClient) Create(ctx context.Context, in *CreateProvider, opts ...grpc.CallOption) (*Provider, error) {
	out := new(Provider)
	err := c.cc.Invoke(ctx, "/organization_service.ProviderService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) GetByID(ctx context.Context, in *ProviderPK, opts ...grpc.CallOption) (*Provider, error) {
	out := new(Provider)
	err := c.cc.Invoke(ctx, "/organization_service.ProviderService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) GetList(ctx context.Context, in *GetListProviderRequest, opts ...grpc.CallOption) (*GetListProviderResponse, error) {
	out := new(GetListProviderResponse)
	err := c.cc.Invoke(ctx, "/organization_service.ProviderService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) Update(ctx context.Context, in *UpdateProvider, opts ...grpc.CallOption) (*Provider, error) {
	out := new(Provider)
	err := c.cc.Invoke(ctx, "/organization_service.ProviderService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) UpdatePatch(ctx context.Context, in *UpdatePatchProvider, opts ...grpc.CallOption) (*Provider, error) {
	out := new(Provider)
	err := c.cc.Invoke(ctx, "/organization_service.ProviderService/UpdatePatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceClient) Delete(ctx context.Context, in *ProviderPK, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/organization_service.ProviderService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServiceServer is the server API for ProviderService service.
// All implementations must embed UnimplementedProviderServiceServer
// for forward compatibility
type ProviderServiceServer interface {
	Create(context.Context, *CreateProvider) (*Provider, error)
	GetByID(context.Context, *ProviderPK) (*Provider, error)
	GetList(context.Context, *GetListProviderRequest) (*GetListProviderResponse, error)
	Update(context.Context, *UpdateProvider) (*Provider, error)
	UpdatePatch(context.Context, *UpdatePatchProvider) (*Provider, error)
	Delete(context.Context, *ProviderPK) (*empty.Empty, error)
	mustEmbedUnimplementedProviderServiceServer()
}

// UnimplementedProviderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServiceServer struct {
}

func (UnimplementedProviderServiceServer) Create(context.Context, *CreateProvider) (*Provider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProviderServiceServer) GetByID(context.Context, *ProviderPK) (*Provider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedProviderServiceServer) GetList(context.Context, *GetListProviderRequest) (*GetListProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedProviderServiceServer) Update(context.Context, *UpdateProvider) (*Provider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProviderServiceServer) UpdatePatch(context.Context, *UpdatePatchProvider) (*Provider, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatch not implemented")
}
func (UnimplementedProviderServiceServer) Delete(context.Context, *ProviderPK) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProviderServiceServer) mustEmbedUnimplementedProviderServiceServer() {}

// UnsafeProviderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServiceServer will
// result in compilation errors.
type UnsafeProviderServiceServer interface {
	mustEmbedUnimplementedProviderServiceServer()
}

func RegisterProviderServiceServer(s grpc.ServiceRegistrar, srv ProviderServiceServer) {
	s.RegisterService(&ProviderService_ServiceDesc, srv)
}

func _ProviderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.ProviderService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).Create(ctx, req.(*CreateProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.ProviderService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).GetByID(ctx, req.(*ProviderPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.ProviderService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).GetList(ctx, req.(*GetListProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.ProviderService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).Update(ctx, req.(*UpdateProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_UpdatePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatchProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).UpdatePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.ProviderService/UpdatePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).UpdatePatch(ctx, req.(*UpdatePatchProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.ProviderService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServer).Delete(ctx, req.(*ProviderPK))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderService_ServiceDesc is the grpc.ServiceDesc for ProviderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organization_service.ProviderService",
	HandlerType: (*ProviderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProviderService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _ProviderService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ProviderService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProviderService_Update_Handler,
		},
		{
			MethodName: "UpdatePatch",
			Handler:    _ProviderService_UpdatePatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProviderService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider_service.proto",
}
