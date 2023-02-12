// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: metadata/proto/metadata.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetadataClient is the client API for Metadata service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataClient interface {
	SetModificationTime(ctx context.Context, in *SetModificationTimeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetModificationTime(ctx context.Context, in *GetModificationTimeRequest, opts ...grpc.CallOption) (*GetModificationTimeResponse, error)
	GetModificationTimes(ctx context.Context, in *GetModificationTimesRequest, opts ...grpc.CallOption) (*GetModificationTimesResponse, error)
	RemoveMetadata(ctx context.Context, in *RemoveMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFolders(ctx context.Context, in *GetFoldersRequest, opts ...grpc.CallOption) (*GetFoldersResponse, error)
}

type metadataClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataClient(cc grpc.ClientConnInterface) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) SetModificationTime(ctx context.Context, in *SetModificationTimeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Metadata/SetModificationTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) GetModificationTime(ctx context.Context, in *GetModificationTimeRequest, opts ...grpc.CallOption) (*GetModificationTimeResponse, error) {
	out := new(GetModificationTimeResponse)
	err := c.cc.Invoke(ctx, "/proto.Metadata/GetModificationTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) GetModificationTimes(ctx context.Context, in *GetModificationTimesRequest, opts ...grpc.CallOption) (*GetModificationTimesResponse, error) {
	out := new(GetModificationTimesResponse)
	err := c.cc.Invoke(ctx, "/proto.Metadata/GetModificationTimes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) RemoveMetadata(ctx context.Context, in *RemoveMetadataRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.Metadata/RemoveMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) GetFolders(ctx context.Context, in *GetFoldersRequest, opts ...grpc.CallOption) (*GetFoldersResponse, error) {
	out := new(GetFoldersResponse)
	err := c.cc.Invoke(ctx, "/proto.Metadata/GetFolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataServer is the server API for Metadata service.
// All implementations should embed UnimplementedMetadataServer
// for forward compatibility
type MetadataServer interface {
	SetModificationTime(context.Context, *SetModificationTimeRequest) (*emptypb.Empty, error)
	GetModificationTime(context.Context, *GetModificationTimeRequest) (*GetModificationTimeResponse, error)
	GetModificationTimes(context.Context, *GetModificationTimesRequest) (*GetModificationTimesResponse, error)
	RemoveMetadata(context.Context, *RemoveMetadataRequest) (*emptypb.Empty, error)
	GetFolders(context.Context, *GetFoldersRequest) (*GetFoldersResponse, error)
}

// UnimplementedMetadataServer should be embedded to have forward compatible implementations.
type UnimplementedMetadataServer struct {
}

func (UnimplementedMetadataServer) SetModificationTime(context.Context, *SetModificationTimeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetModificationTime not implemented")
}
func (UnimplementedMetadataServer) GetModificationTime(context.Context, *GetModificationTimeRequest) (*GetModificationTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModificationTime not implemented")
}
func (UnimplementedMetadataServer) GetModificationTimes(context.Context, *GetModificationTimesRequest) (*GetModificationTimesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModificationTimes not implemented")
}
func (UnimplementedMetadataServer) RemoveMetadata(context.Context, *RemoveMetadataRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMetadata not implemented")
}
func (UnimplementedMetadataServer) GetFolders(context.Context, *GetFoldersRequest) (*GetFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFolders not implemented")
}

// UnsafeMetadataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataServer will
// result in compilation errors.
type UnsafeMetadataServer interface {
	mustEmbedUnimplementedMetadataServer()
}

func RegisterMetadataServer(s grpc.ServiceRegistrar, srv MetadataServer) {
	s.RegisterService(&Metadata_ServiceDesc, srv)
}

func _Metadata_SetModificationTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetModificationTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).SetModificationTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metadata/SetModificationTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).SetModificationTime(ctx, req.(*SetModificationTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_GetModificationTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModificationTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).GetModificationTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metadata/GetModificationTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).GetModificationTime(ctx, req.(*GetModificationTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_GetModificationTimes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModificationTimesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).GetModificationTimes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metadata/GetModificationTimes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).GetModificationTimes(ctx, req.(*GetModificationTimesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_RemoveMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).RemoveMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metadata/RemoveMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).RemoveMetadata(ctx, req.(*RemoveMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_GetFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).GetFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metadata/GetFolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).GetFolders(ctx, req.(*GetFoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metadata_ServiceDesc is the grpc.ServiceDesc for Metadata service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metadata_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetModificationTime",
			Handler:    _Metadata_SetModificationTime_Handler,
		},
		{
			MethodName: "GetModificationTime",
			Handler:    _Metadata_GetModificationTime_Handler,
		},
		{
			MethodName: "GetModificationTimes",
			Handler:    _Metadata_GetModificationTimes_Handler,
		},
		{
			MethodName: "RemoveMetadata",
			Handler:    _Metadata_RemoveMetadata_Handler,
		},
		{
			MethodName: "GetFolders",
			Handler:    _Metadata_GetFolders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata/proto/metadata.proto",
}
