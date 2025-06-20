// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	PokemonService_CreatePokemon_FullMethodName = "/grpc.PokemonService/CreatePokemon"
	PokemonService_GetPokemon_FullMethodName    = "/grpc.PokemonService/GetPokemon"
	PokemonService_ListPokemons_FullMethodName  = "/grpc.PokemonService/ListPokemons"
	PokemonService_UpdatePokemon_FullMethodName = "/grpc.PokemonService/UpdatePokemon"
	PokemonService_DeletePokemon_FullMethodName = "/grpc.PokemonService/DeletePokemon"
)

// PokemonServiceClient is the client API for PokemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PokemonServiceClient interface {
	CreatePokemon(ctx context.Context, in *CreatePokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error)
	GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error)
	ListPokemons(ctx context.Context, in *ListPokemonsRequest, opts ...grpc.CallOption) (*ListPokemonsResponse, error)
	UpdatePokemon(ctx context.Context, in *UpdatePokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error)
	DeletePokemon(ctx context.Context, in *DeletePokemonRequest, opts ...grpc.CallOption) (*DeletePokemonResponse, error)
}

type pokemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPokemonServiceClient(cc grpc.ClientConnInterface) PokemonServiceClient {
	return &pokemonServiceClient{cc}
}

func (c *pokemonServiceClient) CreatePokemon(ctx context.Context, in *CreatePokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PokemonResponse)
	err := c.cc.Invoke(ctx, PokemonService_CreatePokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) GetPokemon(ctx context.Context, in *GetPokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PokemonResponse)
	err := c.cc.Invoke(ctx, PokemonService_GetPokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) ListPokemons(ctx context.Context, in *ListPokemonsRequest, opts ...grpc.CallOption) (*ListPokemonsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPokemonsResponse)
	err := c.cc.Invoke(ctx, PokemonService_ListPokemons_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) UpdatePokemon(ctx context.Context, in *UpdatePokemonRequest, opts ...grpc.CallOption) (*PokemonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PokemonResponse)
	err := c.cc.Invoke(ctx, PokemonService_UpdatePokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokemonServiceClient) DeletePokemon(ctx context.Context, in *DeletePokemonRequest, opts ...grpc.CallOption) (*DeletePokemonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeletePokemonResponse)
	err := c.cc.Invoke(ctx, PokemonService_DeletePokemon_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokemonServiceServer is the server API for PokemonService service.
// All implementations must embed UnimplementedPokemonServiceServer
// for forward compatibility.
type PokemonServiceServer interface {
	CreatePokemon(context.Context, *CreatePokemonRequest) (*PokemonResponse, error)
	GetPokemon(context.Context, *GetPokemonRequest) (*PokemonResponse, error)
	ListPokemons(context.Context, *ListPokemonsRequest) (*ListPokemonsResponse, error)
	UpdatePokemon(context.Context, *UpdatePokemonRequest) (*PokemonResponse, error)
	DeletePokemon(context.Context, *DeletePokemonRequest) (*DeletePokemonResponse, error)
	mustEmbedUnimplementedPokemonServiceServer()
}

// UnimplementedPokemonServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPokemonServiceServer struct{}

func (UnimplementedPokemonServiceServer) CreatePokemon(context.Context, *CreatePokemonRequest) (*PokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) GetPokemon(context.Context, *GetPokemonRequest) (*PokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPokemon not implemented")
}
func (UnimplementedPokemonServiceServer) ListPokemons(context.Context, *ListPokemonsRequest) (*ListPokemonsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPokemons not implemented")
}
func (UnimplementedPokemonServiceServer) UpdatePokemon(context.Context, *UpdatePokemonRequest) (*PokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) DeletePokemon(context.Context, *DeletePokemonRequest) (*DeletePokemonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePokemon not implemented")
}
func (UnimplementedPokemonServiceServer) mustEmbedUnimplementedPokemonServiceServer() {}
func (UnimplementedPokemonServiceServer) testEmbeddedByValue()                        {}

// UnsafePokemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PokemonServiceServer will
// result in compilation errors.
type UnsafePokemonServiceServer interface {
	mustEmbedUnimplementedPokemonServiceServer()
}

func RegisterPokemonServiceServer(s grpc.ServiceRegistrar, srv PokemonServiceServer) {
	// If the following call pancis, it indicates UnimplementedPokemonServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PokemonService_ServiceDesc, srv)
}

func _PokemonService_CreatePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).CreatePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_CreatePokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).CreatePokemon(ctx, req.(*CreatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_GetPokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).GetPokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_GetPokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).GetPokemon(ctx, req.(*GetPokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_ListPokemons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPokemonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).ListPokemons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_ListPokemons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).ListPokemons(ctx, req.(*ListPokemonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_UpdatePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).UpdatePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_UpdatePokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).UpdatePokemon(ctx, req.(*UpdatePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PokemonService_DeletePokemon_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePokemonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokemonServiceServer).DeletePokemon(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PokemonService_DeletePokemon_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokemonServiceServer).DeletePokemon(ctx, req.(*DeletePokemonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PokemonService_ServiceDesc is the grpc.ServiceDesc for PokemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PokemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.PokemonService",
	HandlerType: (*PokemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePokemon",
			Handler:    _PokemonService_CreatePokemon_Handler,
		},
		{
			MethodName: "GetPokemon",
			Handler:    _PokemonService_GetPokemon_Handler,
		},
		{
			MethodName: "ListPokemons",
			Handler:    _PokemonService_ListPokemons_Handler,
		},
		{
			MethodName: "UpdatePokemon",
			Handler:    _PokemonService_UpdatePokemon_Handler,
		},
		{
			MethodName: "DeletePokemon",
			Handler:    _PokemonService_DeletePokemon_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}

const (
	UserService_CreateUser_FullMethodName        = "/grpc.UserService/CreateUser"
	UserService_GetUser_FullMethodName           = "/grpc.UserService/GetUser"
	UserService_GetUserByUsername_FullMethodName = "/grpc.UserService/GetUserByUsername"
	UserService_Login_FullMethodName             = "/grpc.UserService/Login"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*UserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*UserResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*GetUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
