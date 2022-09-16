// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: auth_account.proto

package authService

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error)
	GetAccountById(ctx context.Context, in *GetAccountByIdReq, opts ...grpc.CallOption) (*Account, error)
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenRes, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*UpdateAccountRes, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error) {
	out := new(LoginRes)
	err := c.cc.Invoke(ctx, "/authentication.authService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, "/authentication.authService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountReq, opts ...grpc.CallOption) (*UpdateAccountRes, error) {
	out := new(UpdateAccountRes)
	err := c.cc.Invoke(ctx, "/authentication.authService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetAccountById(ctx context.Context, in *GetAccountByIdReq, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/authentication.authService/GetAccountById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenRes, error) {
	out := new(VerifyTokenRes)
	err := c.cc.Invoke(ctx, "/authentication.authService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*UpdateAccountRes, error) {
	out := new(UpdateAccountRes)
	err := c.cc.Invoke(ctx, "/authentication.authService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRes, error)
	GetAccountById(context.Context, *GetAccountByIdReq) (*Account, error)
	VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*UpdateAccountRes, error)
}

// UnimplementedAuthServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Login(context.Context, *LoginReq) (*LoginRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) UpdateAccount(context.Context, *UpdateAccountReq) (*UpdateAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (UnimplementedAuthServiceServer) GetAccountById(context.Context, *GetAccountByIdReq) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}
func (UnimplementedAuthServiceServer) VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthServiceServer) ChangePassword(context.Context, *ChangePasswordReq) (*UpdateAccountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.authService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.authService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.authService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateAccount(ctx, req.(*UpdateAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.authService/GetAccountById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetAccountById(ctx, req.(*GetAccountByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.authService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyToken(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.authService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.authService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AuthService_UpdateAccount_Handler,
		},
		{
			MethodName: "GetAccountById",
			Handler:    _AuthService_GetAccountById_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _AuthService_VerifyToken_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthService_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_account.proto",
}
