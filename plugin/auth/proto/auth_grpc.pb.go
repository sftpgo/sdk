// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: auth/proto/auth.proto

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

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	CheckUserAndPass(ctx context.Context, in *CheckUserAndPassRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckUserAndTLSCert(ctx context.Context, in *CheckUserAndTLSCertRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckUserAndPublicKey(ctx context.Context, in *CheckUserAndPublicKeyRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	CheckUserAndKeyboardInteractive(ctx context.Context, in *CheckUserAndKeyboardInteractiveRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	SendKeyboardAuthRequest(ctx context.Context, in *KeyboardAuthRequest, opts ...grpc.CallOption) (*KeyboardAuthResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CheckUserAndPass(ctx context.Context, in *CheckUserAndPassRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckUserAndTLSCert(ctx context.Context, in *CheckUserAndTLSCertRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndTLSCert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckUserAndPublicKey(ctx context.Context, in *CheckUserAndPublicKeyRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckUserAndKeyboardInteractive(ctx context.Context, in *CheckUserAndKeyboardInteractiveRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/CheckUserAndKeyboardInteractive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SendKeyboardAuthRequest(ctx context.Context, in *KeyboardAuthRequest, opts ...grpc.CallOption) (*KeyboardAuthResponse, error) {
	out := new(KeyboardAuthResponse)
	err := c.cc.Invoke(ctx, "/proto.Auth/SendKeyboardAuthRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	CheckUserAndPass(context.Context, *CheckUserAndPassRequest) (*AuthResponse, error)
	CheckUserAndTLSCert(context.Context, *CheckUserAndTLSCertRequest) (*AuthResponse, error)
	CheckUserAndPublicKey(context.Context, *CheckUserAndPublicKeyRequest) (*AuthResponse, error)
	CheckUserAndKeyboardInteractive(context.Context, *CheckUserAndKeyboardInteractiveRequest) (*AuthResponse, error)
	SendKeyboardAuthRequest(context.Context, *KeyboardAuthRequest) (*KeyboardAuthResponse, error)
}

// UnimplementedAuthServer should be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) CheckUserAndPass(context.Context, *CheckUserAndPassRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndPass not implemented")
}
func (UnimplementedAuthServer) CheckUserAndTLSCert(context.Context, *CheckUserAndTLSCertRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndTLSCert not implemented")
}
func (UnimplementedAuthServer) CheckUserAndPublicKey(context.Context, *CheckUserAndPublicKeyRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndPublicKey not implemented")
}
func (UnimplementedAuthServer) CheckUserAndKeyboardInteractive(context.Context, *CheckUserAndKeyboardInteractiveRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAndKeyboardInteractive not implemented")
}
func (UnimplementedAuthServer) SendKeyboardAuthRequest(context.Context, *KeyboardAuthRequest) (*KeyboardAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKeyboardAuthRequest not implemented")
}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_CheckUserAndPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndPass(ctx, req.(*CheckUserAndPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckUserAndTLSCert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndTLSCertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndTLSCert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndTLSCert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndTLSCert(ctx, req.(*CheckUserAndTLSCertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckUserAndPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndPublicKey(ctx, req.(*CheckUserAndPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckUserAndKeyboardInteractive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserAndKeyboardInteractiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckUserAndKeyboardInteractive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/CheckUserAndKeyboardInteractive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckUserAndKeyboardInteractive(ctx, req.(*CheckUserAndKeyboardInteractiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SendKeyboardAuthRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyboardAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SendKeyboardAuthRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Auth/SendKeyboardAuthRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SendKeyboardAuthRequest(ctx, req.(*KeyboardAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUserAndPass",
			Handler:    _Auth_CheckUserAndPass_Handler,
		},
		{
			MethodName: "CheckUserAndTLSCert",
			Handler:    _Auth_CheckUserAndTLSCert_Handler,
		},
		{
			MethodName: "CheckUserAndPublicKey",
			Handler:    _Auth_CheckUserAndPublicKey_Handler,
		},
		{
			MethodName: "CheckUserAndKeyboardInteractive",
			Handler:    _Auth_CheckUserAndKeyboardInteractive_Handler,
		},
		{
			MethodName: "SendKeyboardAuthRequest",
			Handler:    _Auth_SendKeyboardAuthRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/proto/auth.proto",
}
