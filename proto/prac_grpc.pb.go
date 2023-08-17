// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: prac.proto

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

const (
	AuthSampleService_LoginMethod_FullMethodName           = "/proto.AuthSampleService/LoginMethod"
	AuthSampleService_AdminMethod_FullMethodName           = "/proto.AuthSampleService/AdminMethod"
	AuthSampleService_RequiredAuthMethod_FullMethodName    = "/proto.AuthSampleService/RequiredAuthMethod"
	AuthSampleService_NotRequiredAuthMethod_FullMethodName = "/proto.AuthSampleService/NotRequiredAuthMethod"
)

// AuthSampleServiceClient is the client API for AuthSampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthSampleServiceClient interface {
	LoginMethod(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	AdminMethod(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResponse, error)
	RequiredAuthMethod(ctx context.Context, in *RequiredAuthRequest, opts ...grpc.CallOption) (*RequiredAuthResponse, error)
	NotRequiredAuthMethod(ctx context.Context, in *NotRequiredAuthRequest, opts ...grpc.CallOption) (*NotRequiredAuthResponse, error)
}

type authSampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthSampleServiceClient(cc grpc.ClientConnInterface) AuthSampleServiceClient {
	return &authSampleServiceClient{cc}
}

func (c *authSampleServiceClient) LoginMethod(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthSampleService_LoginMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSampleServiceClient) AdminMethod(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AuthSampleService_AdminMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSampleServiceClient) RequiredAuthMethod(ctx context.Context, in *RequiredAuthRequest, opts ...grpc.CallOption) (*RequiredAuthResponse, error) {
	out := new(RequiredAuthResponse)
	err := c.cc.Invoke(ctx, AuthSampleService_RequiredAuthMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authSampleServiceClient) NotRequiredAuthMethod(ctx context.Context, in *NotRequiredAuthRequest, opts ...grpc.CallOption) (*NotRequiredAuthResponse, error) {
	out := new(NotRequiredAuthResponse)
	err := c.cc.Invoke(ctx, AuthSampleService_NotRequiredAuthMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthSampleServiceServer is the server API for AuthSampleService service.
// All implementations must embed UnimplementedAuthSampleServiceServer
// for forward compatibility
type AuthSampleServiceServer interface {
	LoginMethod(context.Context, *LoginRequest) (*LoginResponse, error)
	AdminMethod(context.Context, *AdminRequest) (*AdminResponse, error)
	RequiredAuthMethod(context.Context, *RequiredAuthRequest) (*RequiredAuthResponse, error)
	NotRequiredAuthMethod(context.Context, *NotRequiredAuthRequest) (*NotRequiredAuthResponse, error)
	mustEmbedUnimplementedAuthSampleServiceServer()
}

// UnimplementedAuthSampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthSampleServiceServer struct {
}

func (UnimplementedAuthSampleServiceServer) LoginMethod(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginMethod not implemented")
}
func (UnimplementedAuthSampleServiceServer) AdminMethod(context.Context, *AdminRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminMethod not implemented")
}
func (UnimplementedAuthSampleServiceServer) RequiredAuthMethod(context.Context, *RequiredAuthRequest) (*RequiredAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequiredAuthMethod not implemented")
}
func (UnimplementedAuthSampleServiceServer) NotRequiredAuthMethod(context.Context, *NotRequiredAuthRequest) (*NotRequiredAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotRequiredAuthMethod not implemented")
}
func (UnimplementedAuthSampleServiceServer) mustEmbedUnimplementedAuthSampleServiceServer() {}

// UnsafeAuthSampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthSampleServiceServer will
// result in compilation errors.
type UnsafeAuthSampleServiceServer interface {
	mustEmbedUnimplementedAuthSampleServiceServer()
}

func RegisterAuthSampleServiceServer(s grpc.ServiceRegistrar, srv AuthSampleServiceServer) {
	s.RegisterService(&AuthSampleService_ServiceDesc, srv)
}

func _AuthSampleService_LoginMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSampleServiceServer).LoginMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSampleService_LoginMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSampleServiceServer).LoginMethod(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSampleService_AdminMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSampleServiceServer).AdminMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSampleService_AdminMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSampleServiceServer).AdminMethod(ctx, req.(*AdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSampleService_RequiredAuthMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequiredAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSampleServiceServer).RequiredAuthMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSampleService_RequiredAuthMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSampleServiceServer).RequiredAuthMethod(ctx, req.(*RequiredAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthSampleService_NotRequiredAuthMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotRequiredAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthSampleServiceServer).NotRequiredAuthMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthSampleService_NotRequiredAuthMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthSampleServiceServer).NotRequiredAuthMethod(ctx, req.(*NotRequiredAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthSampleService_ServiceDesc is the grpc.ServiceDesc for AuthSampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthSampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthSampleService",
	HandlerType: (*AuthSampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginMethod",
			Handler:    _AuthSampleService_LoginMethod_Handler,
		},
		{
			MethodName: "AdminMethod",
			Handler:    _AuthSampleService_AdminMethod_Handler,
		},
		{
			MethodName: "RequiredAuthMethod",
			Handler:    _AuthSampleService_RequiredAuthMethod_Handler,
		},
		{
			MethodName: "NotRequiredAuthMethod",
			Handler:    _AuthSampleService_NotRequiredAuthMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prac.proto",
}
