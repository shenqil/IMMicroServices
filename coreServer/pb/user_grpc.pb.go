// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user.proto

package pb

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

// UserGreeterClient is the client API for UserGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGreeterClient interface {
	Query(ctx context.Context, in *UserQueryRequest, opts ...grpc.CallOption) (*UserQueryResult, error)
	Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*User, error)
	Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*IDResult, error)
	Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatusResult, error)
	Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*StatusResult, error)
	UpdateStatus(ctx context.Context, in *UserUpdateStatusRequest, opts ...grpc.CallOption) (*StatusResult, error)
	UpdatePassword(ctx context.Context, in *UserUpdatePasswordRequest, opts ...grpc.CallOption) (*StatusResult, error)
	Verify(ctx context.Context, in *UserVerifyRequest, opts ...grpc.CallOption) (*StatusResult, error)
}

type userGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGreeterClient(cc grpc.ClientConnInterface) UserGreeterClient {
	return &userGreeterClient{cc}
}

func (c *userGreeterClient) Query(ctx context.Context, in *UserQueryRequest, opts ...grpc.CallOption) (*UserQueryResult, error) {
	out := new(UserQueryResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) Get(ctx context.Context, in *UserGetRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) Create(ctx context.Context, in *User, opts ...grpc.CallOption) (*IDResult, error) {
	out := new(IDResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) Update(ctx context.Context, in *User, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) Delete(ctx context.Context, in *UserDeleteRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) UpdateStatus(ctx context.Context, in *UserUpdateStatusRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) UpdatePassword(ctx context.Context, in *UserUpdatePasswordRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGreeterClient) Verify(ctx context.Context, in *UserVerifyRequest, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/userPB.UserGreeter/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGreeterServer is the server API for UserGreeter service.
// All implementations must embed UnimplementedUserGreeterServer
// for forward compatibility
type UserGreeterServer interface {
	Query(context.Context, *UserQueryRequest) (*UserQueryResult, error)
	Get(context.Context, *UserGetRequest) (*User, error)
	Create(context.Context, *User) (*IDResult, error)
	Update(context.Context, *User) (*StatusResult, error)
	Delete(context.Context, *UserDeleteRequest) (*StatusResult, error)
	UpdateStatus(context.Context, *UserUpdateStatusRequest) (*StatusResult, error)
	UpdatePassword(context.Context, *UserUpdatePasswordRequest) (*StatusResult, error)
	Verify(context.Context, *UserVerifyRequest) (*StatusResult, error)
	mustEmbedUnimplementedUserGreeterServer()
}

// UnimplementedUserGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedUserGreeterServer struct {
}

func (UnimplementedUserGreeterServer) Query(context.Context, *UserQueryRequest) (*UserQueryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedUserGreeterServer) Get(context.Context, *UserGetRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserGreeterServer) Create(context.Context, *User) (*IDResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserGreeterServer) Update(context.Context, *User) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserGreeterServer) Delete(context.Context, *UserDeleteRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserGreeterServer) UpdateStatus(context.Context, *UserUpdateStatusRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedUserGreeterServer) UpdatePassword(context.Context, *UserUpdatePasswordRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserGreeterServer) Verify(context.Context, *UserVerifyRequest) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedUserGreeterServer) mustEmbedUnimplementedUserGreeterServer() {}

// UnsafeUserGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGreeterServer will
// result in compilation errors.
type UnsafeUserGreeterServer interface {
	mustEmbedUnimplementedUserGreeterServer()
}

func RegisterUserGreeterServer(s grpc.ServiceRegistrar, srv UserGreeterServer) {
	s.RegisterService(&UserGreeter_ServiceDesc, srv)
}

func _UserGreeter_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).Query(ctx, req.(*UserQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).Get(ctx, req.(*UserGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).Create(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).Update(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).Delete(ctx, req.(*UserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).UpdateStatus(ctx, req.(*UserUpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).UpdatePassword(ctx, req.(*UserUpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGreeter_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGreeterServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userPB.UserGreeter/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGreeterServer).Verify(ctx, req.(*UserVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGreeter_ServiceDesc is the grpc.ServiceDesc for UserGreeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userPB.UserGreeter",
	HandlerType: (*UserGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _UserGreeter_Query_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserGreeter_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserGreeter_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserGreeter_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserGreeter_Delete_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _UserGreeter_UpdateStatus_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserGreeter_UpdatePassword_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _UserGreeter_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}