// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: auth_service.proto

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
	// auth
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginRes, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenRes, error)
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

func (c *authServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyTokenRes, error) {
	out := new(VerifyTokenRes)
	err := c.cc.Invoke(ctx, "/authentication.authService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations should embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	// auth
	Login(context.Context, *LoginReq) (*LoginRes, error)
	Register(context.Context, *RegisterReq) (*RegisterRes, error)
	VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error)
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
func (UnimplementedAuthServiceServer) VerifyToken(context.Context, *VerifyTokenReq) (*VerifyTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
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
			MethodName: "VerifyToken",
			Handler:    _AuthService_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	// account
	GetAccountById(ctx context.Context, in *GetAccountByIdReq, opts ...grpc.CallOption) (*GetAccountByIdRes, error)
	GetAccountByUsername(ctx context.Context, in *GetAccountByUsernameReq, opts ...grpc.CallOption) (*GetAccountByUsernameRes, error)
	GetAccountByEmail(ctx context.Context, in *GetAccountByEmailReq, opts ...grpc.CallOption) (*GetAccountByEmailRes, error)
	ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordRes, error)
	SearchAccount(ctx context.Context, in *SearchAccountsReq, opts ...grpc.CallOption) (*SearchAccountsRes, error)
	DeleteAccountById(ctx context.Context, in *DeleteAccountByIdReq, opts ...grpc.CallOption) (*DeleteAccountByIdRes, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) GetAccountById(ctx context.Context, in *GetAccountByIdReq, opts ...grpc.CallOption) (*GetAccountByIdRes, error) {
	out := new(GetAccountByIdRes)
	err := c.cc.Invoke(ctx, "/authentication.accountService/GetAccountById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByUsername(ctx context.Context, in *GetAccountByUsernameReq, opts ...grpc.CallOption) (*GetAccountByUsernameRes, error) {
	out := new(GetAccountByUsernameRes)
	err := c.cc.Invoke(ctx, "/authentication.accountService/GetAccountByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByEmail(ctx context.Context, in *GetAccountByEmailReq, opts ...grpc.CallOption) (*GetAccountByEmailRes, error) {
	out := new(GetAccountByEmailRes)
	err := c.cc.Invoke(ctx, "/authentication.accountService/GetAccountByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordReq, opts ...grpc.CallOption) (*ChangePasswordRes, error) {
	out := new(ChangePasswordRes)
	err := c.cc.Invoke(ctx, "/authentication.accountService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SearchAccount(ctx context.Context, in *SearchAccountsReq, opts ...grpc.CallOption) (*SearchAccountsRes, error) {
	out := new(SearchAccountsRes)
	err := c.cc.Invoke(ctx, "/authentication.accountService/SearchAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccountById(ctx context.Context, in *DeleteAccountByIdReq, opts ...grpc.CallOption) (*DeleteAccountByIdRes, error) {
	out := new(DeleteAccountByIdRes)
	err := c.cc.Invoke(ctx, "/authentication.accountService/DeleteAccountByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	// account
	GetAccountById(context.Context, *GetAccountByIdReq) (*GetAccountByIdRes, error)
	GetAccountByUsername(context.Context, *GetAccountByUsernameReq) (*GetAccountByUsernameRes, error)
	GetAccountByEmail(context.Context, *GetAccountByEmailReq) (*GetAccountByEmailRes, error)
	ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordRes, error)
	SearchAccount(context.Context, *SearchAccountsReq) (*SearchAccountsRes, error)
	DeleteAccountByID(context.Context, *DeleteAccountByIdReq) (*DeleteAccountByIdRes, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) GetAccountById(context.Context, *GetAccountByIdReq) (*GetAccountByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByUsername(context.Context, *GetAccountByUsernameReq) (*GetAccountByUsernameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByUsername not implemented")
}
func (UnimplementedAccountServiceServer) GetAccountByEmail(context.Context, *GetAccountByEmailReq) (*GetAccountByEmailRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByEmail not implemented")
}
func (UnimplementedAccountServiceServer) ChangePassword(context.Context, *ChangePasswordReq) (*ChangePasswordRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAccountServiceServer) SearchAccount(context.Context, *SearchAccountsReq) (*SearchAccountsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAccount not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccountByID(context.Context, *DeleteAccountByIdReq) (*DeleteAccountByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccountByID not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.accountService/GetAccountById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountById(ctx, req.(*GetAccountByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByUsernameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.accountService/GetAccountByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByUsername(ctx, req.(*GetAccountByUsernameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountByEmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.accountService/GetAccountByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByEmail(ctx, req.(*GetAccountByEmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.accountService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ChangePassword(ctx, req.(*ChangePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SearchAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAccountsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SearchAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.accountService/SearchAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SearchAccount(ctx, req.(*SearchAccountsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccountByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.accountService/DeleteAccountByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccountByID(ctx, req.(*DeleteAccountByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.accountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountById",
			Handler:    _AccountService_GetAccountById_Handler,
		},
		{
			MethodName: "GetAccountByUsername",
			Handler:    _AccountService_GetAccountByUsername_Handler,
		},
		{
			MethodName: "GetAccountByEmail",
			Handler:    _AccountService_GetAccountByEmail_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AccountService_ChangePassword_Handler,
		},
		{
			MethodName: "SearchAccount",
			Handler:    _AccountService_SearchAccount_Handler,
		},
		{
			MethodName: "DeleteAccountByID",
			Handler:    _AccountService_DeleteAccountById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}

// RoleServiceClient is the client API for RoleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleServiceClient interface {
	// role
	CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleRes, error)
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleRes, error)
	SearchRole(ctx context.Context, in *SearchRolesReq, opts ...grpc.CallOption) (*SearchRolesRes, error)
	DeleteRoleById(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleRes, error)
}

type roleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleServiceClient(cc grpc.ClientConnInterface) RoleServiceClient {
	return &roleServiceClient{cc}
}

func (c *roleServiceClient) CreateRole(ctx context.Context, in *CreateRoleReq, opts ...grpc.CallOption) (*CreateRoleRes, error) {
	out := new(CreateRoleRes)
	err := c.cc.Invoke(ctx, "/authentication.roleService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*UpdateRoleRes, error) {
	out := new(UpdateRoleRes)
	err := c.cc.Invoke(ctx, "/authentication.roleService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) SearchRole(ctx context.Context, in *SearchRolesReq, opts ...grpc.CallOption) (*SearchRolesRes, error) {
	out := new(SearchRolesRes)
	err := c.cc.Invoke(ctx, "/authentication.roleService/SearchRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleServiceClient) DeleteRoleById(ctx context.Context, in *DeleteRoleReq, opts ...grpc.CallOption) (*DeleteRoleRes, error) {
	out := new(DeleteRoleRes)
	err := c.cc.Invoke(ctx, "/authentication.roleService/DeleteRoleById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServiceServer is the server API for RoleService service.
// All implementations should embed UnimplementedRoleServiceServer
// for forward compatibility
type RoleServiceServer interface {
	// role
	CreateRole(context.Context, *CreateRoleReq) (*CreateRoleRes, error)
	UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleRes, error)
	SearchRole(context.Context, *SearchRolesReq) (*SearchRolesRes, error)
	DeleteRoleById(context.Context, *DeleteRoleReq) (*DeleteRoleRes, error)
}

// UnimplementedRoleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRoleServiceServer struct {
}

func (UnimplementedRoleServiceServer) CreateRole(context.Context, *CreateRoleReq) (*CreateRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedRoleServiceServer) UpdateRole(context.Context, *UpdateRoleReq) (*UpdateRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServiceServer) SearchRole(context.Context, *SearchRolesReq) (*SearchRolesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRole not implemented")
}
func (UnimplementedRoleServiceServer) DeleteRoleById(context.Context, *DeleteRoleReq) (*DeleteRoleRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleById not implemented")
}

// UnsafeRoleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServiceServer will
// result in compilation errors.
type UnsafeRoleServiceServer interface {
	mustEmbedUnimplementedRoleServiceServer()
}

func RegisterRoleServiceServer(s grpc.ServiceRegistrar, srv RoleServiceServer) {
	s.RegisterService(&RoleService_ServiceDesc, srv)
}

func _RoleService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.roleService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).CreateRole(ctx, req.(*CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.roleService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_SearchRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRolesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).SearchRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.roleService/SearchRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).SearchRole(ctx, req.(*SearchRolesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoleService_DeleteRoleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServiceServer).DeleteRoleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.roleService/DeleteRoleById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServiceServer).DeleteRoleById(ctx, req.(*DeleteRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RoleService_ServiceDesc is the grpc.ServiceDesc for RoleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.roleService",
	HandlerType: (*RoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRole",
			Handler:    _RoleService_CreateRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _RoleService_UpdateRole_Handler,
		},
		{
			MethodName: "SearchRole",
			Handler:    _RoleService_SearchRole_Handler,
		},
		{
			MethodName: "DeleteRoleById",
			Handler:    _RoleService_DeleteRoleById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
