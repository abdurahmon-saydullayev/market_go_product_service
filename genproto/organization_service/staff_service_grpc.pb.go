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

// StaffServiceClient is the client API for StaffService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaffServiceClient interface {
	Create(ctx context.Context, in *CreateStaff, opts ...grpc.CallOption) (*Staff, error)
	GetByID(ctx context.Context, in *StaffPK, opts ...grpc.CallOption) (*Staff, error)
	GetList(ctx context.Context, in *GetListStaffRequest, opts ...grpc.CallOption) (*GetListStaffResponse, error)
	Update(ctx context.Context, in *UpdateStaff, opts ...grpc.CallOption) (*Staff, error)
	UpdatePatch(ctx context.Context, in *UpdatePatchStaff, opts ...grpc.CallOption) (*Staff, error)
	Delete(ctx context.Context, in *StaffPK, opts ...grpc.CallOption) (*empty.Empty, error)
}

type staffServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStaffServiceClient(cc grpc.ClientConnInterface) StaffServiceClient {
	return &staffServiceClient{cc}
}

func (c *staffServiceClient) Create(ctx context.Context, in *CreateStaff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/organization_service.StaffService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetByID(ctx context.Context, in *StaffPK, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/organization_service.StaffService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) GetList(ctx context.Context, in *GetListStaffRequest, opts ...grpc.CallOption) (*GetListStaffResponse, error) {
	out := new(GetListStaffResponse)
	err := c.cc.Invoke(ctx, "/organization_service.StaffService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Update(ctx context.Context, in *UpdateStaff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/organization_service.StaffService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) UpdatePatch(ctx context.Context, in *UpdatePatchStaff, opts ...grpc.CallOption) (*Staff, error) {
	out := new(Staff)
	err := c.cc.Invoke(ctx, "/organization_service.StaffService/UpdatePatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *staffServiceClient) Delete(ctx context.Context, in *StaffPK, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/organization_service.StaffService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaffServiceServer is the server API for StaffService service.
// All implementations must embed UnimplementedStaffServiceServer
// for forward compatibility
type StaffServiceServer interface {
	Create(context.Context, *CreateStaff) (*Staff, error)
	GetByID(context.Context, *StaffPK) (*Staff, error)
	GetList(context.Context, *GetListStaffRequest) (*GetListStaffResponse, error)
	Update(context.Context, *UpdateStaff) (*Staff, error)
	UpdatePatch(context.Context, *UpdatePatchStaff) (*Staff, error)
	Delete(context.Context, *StaffPK) (*empty.Empty, error)
	mustEmbedUnimplementedStaffServiceServer()
}

// UnimplementedStaffServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStaffServiceServer struct {
}

func (UnimplementedStaffServiceServer) Create(context.Context, *CreateStaff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStaffServiceServer) GetByID(context.Context, *StaffPK) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedStaffServiceServer) GetList(context.Context, *GetListStaffRequest) (*GetListStaffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStaffServiceServer) Update(context.Context, *UpdateStaff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStaffServiceServer) UpdatePatch(context.Context, *UpdatePatchStaff) (*Staff, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePatch not implemented")
}
func (UnimplementedStaffServiceServer) Delete(context.Context, *StaffPK) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaffServiceServer) mustEmbedUnimplementedStaffServiceServer() {}

// UnsafeStaffServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaffServiceServer will
// result in compilation errors.
type UnsafeStaffServiceServer interface {
	mustEmbedUnimplementedStaffServiceServer()
}

func RegisterStaffServiceServer(s grpc.ServiceRegistrar, srv StaffServiceServer) {
	s.RegisterService(&StaffService_ServiceDesc, srv)
}

func _StaffService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStaff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StaffService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Create(ctx, req.(*CreateStaff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StaffService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetByID(ctx, req.(*StaffPK))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStaffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StaffService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).GetList(ctx, req.(*GetListStaffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStaff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StaffService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Update(ctx, req.(*UpdateStaff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_UpdatePatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePatchStaff)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).UpdatePatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StaffService/UpdatePatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).UpdatePatch(ctx, req.(*UpdatePatchStaff))
	}
	return interceptor(ctx, in, info, handler)
}

func _StaffService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StaffPK)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaffServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/organization_service.StaffService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaffServiceServer).Delete(ctx, req.(*StaffPK))
	}
	return interceptor(ctx, in, info, handler)
}

// StaffService_ServiceDesc is the grpc.ServiceDesc for StaffService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StaffService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "organization_service.StaffService",
	HandlerType: (*StaffServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StaffService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _StaffService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _StaffService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StaffService_Update_Handler,
		},
		{
			MethodName: "UpdatePatch",
			Handler:    _StaffService_UpdatePatch_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StaffService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staff_service.proto",
}