// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/master.proto

package proto

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

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterServiceClient interface {
	CreateMaster(ctx context.Context, in *CreateMasterRequest, opts ...grpc.CallOption) (*CreateMasterResponse, error)
	GetMaster(ctx context.Context, in *GetMasterRequest, opts ...grpc.CallOption) (*GetMasterResponse, error)
	UpdateMaster(ctx context.Context, in *UpdateMasterRequest, opts ...grpc.CallOption) (*UpdateMasterResponse, error)
	DeleteMaster(ctx context.Context, in *DeleteMasterRequest, opts ...grpc.CallOption) (*DeleteMasterResponse, error)
	ListMasters(ctx context.Context, in *ListMastersRequest, opts ...grpc.CallOption) (*ListMastersResponse, error)
	LoginMaster(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LogoutMaster(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) CreateMaster(ctx context.Context, in *CreateMasterRequest, opts ...grpc.CallOption) (*CreateMasterResponse, error) {
	out := new(CreateMasterResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/CreateMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetMaster(ctx context.Context, in *GetMasterRequest, opts ...grpc.CallOption) (*GetMasterResponse, error) {
	out := new(GetMasterResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/GetMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) UpdateMaster(ctx context.Context, in *UpdateMasterRequest, opts ...grpc.CallOption) (*UpdateMasterResponse, error) {
	out := new(UpdateMasterResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/UpdateMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) DeleteMaster(ctx context.Context, in *DeleteMasterRequest, opts ...grpc.CallOption) (*DeleteMasterResponse, error) {
	out := new(DeleteMasterResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/DeleteMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ListMasters(ctx context.Context, in *ListMastersRequest, opts ...grpc.CallOption) (*ListMastersResponse, error) {
	out := new(ListMastersResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/ListMasters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) LoginMaster(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/LoginMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) LogoutMaster(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/proto.MasterService/LogoutMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility
type MasterServiceServer interface {
	CreateMaster(context.Context, *CreateMasterRequest) (*CreateMasterResponse, error)
	GetMaster(context.Context, *GetMasterRequest) (*GetMasterResponse, error)
	UpdateMaster(context.Context, *UpdateMasterRequest) (*UpdateMasterResponse, error)
	DeleteMaster(context.Context, *DeleteMasterRequest) (*DeleteMasterResponse, error)
	ListMasters(context.Context, *ListMastersRequest) (*ListMastersResponse, error)
	LoginMaster(context.Context, *LoginRequest) (*LoginResponse, error)
	LogoutMaster(context.Context, *LogoutRequest) (*LogoutResponse, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServiceServer struct {
}

func (UnimplementedMasterServiceServer) CreateMaster(context.Context, *CreateMasterRequest) (*CreateMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMaster not implemented")
}
func (UnimplementedMasterServiceServer) GetMaster(context.Context, *GetMasterRequest) (*GetMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaster not implemented")
}
func (UnimplementedMasterServiceServer) UpdateMaster(context.Context, *UpdateMasterRequest) (*UpdateMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMaster not implemented")
}
func (UnimplementedMasterServiceServer) DeleteMaster(context.Context, *DeleteMasterRequest) (*DeleteMasterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMaster not implemented")
}
func (UnimplementedMasterServiceServer) ListMasters(context.Context, *ListMastersRequest) (*ListMastersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMasters not implemented")
}
func (UnimplementedMasterServiceServer) LoginMaster(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginMaster not implemented")
}
func (UnimplementedMasterServiceServer) LogoutMaster(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogoutMaster not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_CreateMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).CreateMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/CreateMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).CreateMaster(ctx, req.(*CreateMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/GetMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetMaster(ctx, req.(*GetMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_UpdateMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).UpdateMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/UpdateMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).UpdateMaster(ctx, req.(*UpdateMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_DeleteMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).DeleteMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/DeleteMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).DeleteMaster(ctx, req.(*DeleteMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ListMasters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMastersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ListMasters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/ListMasters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ListMasters(ctx, req.(*ListMastersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_LoginMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).LoginMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/LoginMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).LoginMaster(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_LogoutMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).LogoutMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MasterService/LogoutMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).LogoutMaster(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMaster",
			Handler:    _MasterService_CreateMaster_Handler,
		},
		{
			MethodName: "GetMaster",
			Handler:    _MasterService_GetMaster_Handler,
		},
		{
			MethodName: "UpdateMaster",
			Handler:    _MasterService_UpdateMaster_Handler,
		},
		{
			MethodName: "DeleteMaster",
			Handler:    _MasterService_DeleteMaster_Handler,
		},
		{
			MethodName: "ListMasters",
			Handler:    _MasterService_ListMasters_Handler,
		},
		{
			MethodName: "LoginMaster",
			Handler:    _MasterService_LoginMaster_Handler,
		},
		{
			MethodName: "LogoutMaster",
			Handler:    _MasterService_LogoutMaster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/master.proto",
}