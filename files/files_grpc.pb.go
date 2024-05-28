// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: files.proto

package files

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
	Files_CreateFile_FullMethodName   = "/Files/CreateFile"
	Files_CheckFile_FullMethodName    = "/Files/CheckFile"
	Files_WroteToFile_FullMethodName  = "/Files/WroteToFile"
	Files_ReadFromFile_FullMethodName = "/Files/ReadFromFile"
)

// FilesClient is the client API for Files service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesClient interface {
	CreateFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	CheckFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	WroteToFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	ReadFromFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type filesClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesClient(cc grpc.ClientConnInterface) FilesClient {
	return &filesClient{cc}
}

func (c *filesClient) CreateFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Files_CreateFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) CheckFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Files_CheckFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) WroteToFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Files_WroteToFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesClient) ReadFromFile(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, Files_ReadFromFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilesServer is the server API for Files service.
// All implementations must embed UnimplementedFilesServer
// for forward compatibility
type FilesServer interface {
	CreateFile(context.Context, *CreateRequest) (*CreateResponse, error)
	CheckFile(context.Context, *CreateRequest) (*CreateResponse, error)
	WroteToFile(context.Context, *CreateRequest) (*CreateResponse, error)
	ReadFromFile(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedFilesServer()
}

// UnimplementedFilesServer must be embedded to have forward compatible implementations.
type UnimplementedFilesServer struct {
}

func (UnimplementedFilesServer) CreateFile(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedFilesServer) CheckFile(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckFile not implemented")
}
func (UnimplementedFilesServer) WroteToFile(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WroteToFile not implemented")
}
func (UnimplementedFilesServer) ReadFromFile(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFromFile not implemented")
}
func (UnimplementedFilesServer) mustEmbedUnimplementedFilesServer() {}

// UnsafeFilesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesServer will
// result in compilation errors.
type UnsafeFilesServer interface {
	mustEmbedUnimplementedFilesServer()
}

func RegisterFilesServer(s grpc.ServiceRegistrar, srv FilesServer) {
	s.RegisterService(&Files_ServiceDesc, srv)
}

func _Files_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Files_CreateFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).CreateFile(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_CheckFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).CheckFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Files_CheckFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).CheckFile(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_WroteToFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).WroteToFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Files_WroteToFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).WroteToFile(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Files_ReadFromFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServer).ReadFromFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Files_ReadFromFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServer).ReadFromFile(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Files_ServiceDesc is the grpc.ServiceDesc for Files service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Files_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Files",
	HandlerType: (*FilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFile",
			Handler:    _Files_CreateFile_Handler,
		},
		{
			MethodName: "CheckFile",
			Handler:    _Files_CheckFile_Handler,
		},
		{
			MethodName: "WroteToFile",
			Handler:    _Files_WroteToFile_Handler,
		},
		{
			MethodName: "ReadFromFile",
			Handler:    _Files_ReadFromFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "files.proto",
}