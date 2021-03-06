// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/user.proto

package api

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

// UserManagerClient is the client API for UserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagerClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error)
	GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersReply, error)
}

type userManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagerClient(cc grpc.ClientConnInterface) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, "/api.UserManager/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserReply, error) {
	out := new(GetUserReply)
	err := c.cc.Invoke(ctx, "/api.UserManager/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserReply, error) {
	out := new(UpdateUserReply)
	err := c.cc.Invoke(ctx, "/api.UserManager/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserReply, error) {
	out := new(DeleteUserReply)
	err := c.cc.Invoke(ctx, "/api.UserManager/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagerClient) GetAllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAllUsersReply, error) {
	out := new(GetAllUsersReply)
	err := c.cc.Invoke(ctx, "/api.UserManager/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagerServer is the server API for UserManager service.
// All implementations must embed UnimplementedUserManagerServer
// for forward compatibility
type UserManagerServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersReply, error)
	mustEmbedUnimplementedUserManagerServer()
}

// UnimplementedUserManagerServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagerServer struct {
}

func (UnimplementedUserManagerServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserManagerServer) GetUser(context.Context, *GetUserRequest) (*GetUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserManagerServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserManagerServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagerServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*GetAllUsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserManagerServer) mustEmbedUnimplementedUserManagerServer() {}

// UnsafeUserManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagerServer will
// result in compilation errors.
type UnsafeUserManagerServer interface {
	mustEmbedUnimplementedUserManagerServer()
}

func RegisterUserManagerServer(s grpc.ServiceRegistrar, srv UserManagerServer) {
	s.RegisterService(&UserManager_ServiceDesc, srv)
}

func _UserManager_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserManager/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserManager/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserManager/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserManager/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManager_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserManager/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManager_ServiceDesc is the grpc.ServiceDesc for UserManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserManager_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserManager_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManager_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManager_DeleteUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserManager_GetAllUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user.proto",
}
