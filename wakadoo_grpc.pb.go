// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: wakadoo.proto

package wakadoo_grpc_contract

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

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	Save(ctx context.Context, in *FileSaveRequest, opts ...grpc.CallOption) (*FileSaveResponse, error)
	Get(ctx context.Context, in *FileGetRequest, opts ...grpc.CallOption) (*FileGetResponse, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) Save(ctx context.Context, in *FileSaveRequest, opts ...grpc.CallOption) (*FileSaveResponse, error) {
	out := new(FileSaveResponse)
	err := c.cc.Invoke(ctx, "/services.File/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Get(ctx context.Context, in *FileGetRequest, opts ...grpc.CallOption) (*FileGetResponse, error) {
	out := new(FileGetResponse)
	err := c.cc.Invoke(ctx, "/services.File/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility
type FileServer interface {
	Save(context.Context, *FileSaveRequest) (*FileSaveResponse, error)
	Get(context.Context, *FileGetRequest) (*FileGetResponse, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (UnimplementedFileServer) Save(context.Context, *FileSaveRequest) (*FileSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedFileServer) Get(context.Context, *FileGetRequest) (*FileGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.File/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Save(ctx, req.(*FileSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.File/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Get(ctx, req.(*FileGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _File_Save_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _File_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wakadoo.proto",
}

// SessionClient is the client API for Session service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionClient interface {
	Create(ctx context.Context, in *SessionCreateRequest, opts ...grpc.CallOption) (*SessionCreateResponse, error)
	GetEmail(ctx context.Context, in *SessionGetEmailRequest, opts ...grpc.CallOption) (*SessionGetEmailResponse, error)
	UpdateLastActivityTime(ctx context.Context, in *SessionUpdateActivityRequest, opts ...grpc.CallOption) (*SessionUpdateActivityResponse, error)
	GetActive(ctx context.Context, in *SessionGetActiveRequest, opts ...grpc.CallOption) (*SessionGetActiveResponse, error)
}

type sessionClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionClient(cc grpc.ClientConnInterface) SessionClient {
	return &sessionClient{cc}
}

func (c *sessionClient) Create(ctx context.Context, in *SessionCreateRequest, opts ...grpc.CallOption) (*SessionCreateResponse, error) {
	out := new(SessionCreateResponse)
	err := c.cc.Invoke(ctx, "/services.Session/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) GetEmail(ctx context.Context, in *SessionGetEmailRequest, opts ...grpc.CallOption) (*SessionGetEmailResponse, error) {
	out := new(SessionGetEmailResponse)
	err := c.cc.Invoke(ctx, "/services.Session/GetEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) UpdateLastActivityTime(ctx context.Context, in *SessionUpdateActivityRequest, opts ...grpc.CallOption) (*SessionUpdateActivityResponse, error) {
	out := new(SessionUpdateActivityResponse)
	err := c.cc.Invoke(ctx, "/services.Session/UpdateLastActivityTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionClient) GetActive(ctx context.Context, in *SessionGetActiveRequest, opts ...grpc.CallOption) (*SessionGetActiveResponse, error) {
	out := new(SessionGetActiveResponse)
	err := c.cc.Invoke(ctx, "/services.Session/GetActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServer is the server API for Session service.
// All implementations must embed UnimplementedSessionServer
// for forward compatibility
type SessionServer interface {
	Create(context.Context, *SessionCreateRequest) (*SessionCreateResponse, error)
	GetEmail(context.Context, *SessionGetEmailRequest) (*SessionGetEmailResponse, error)
	UpdateLastActivityTime(context.Context, *SessionUpdateActivityRequest) (*SessionUpdateActivityResponse, error)
	GetActive(context.Context, *SessionGetActiveRequest) (*SessionGetActiveResponse, error)
	mustEmbedUnimplementedSessionServer()
}

// UnimplementedSessionServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServer struct {
}

func (UnimplementedSessionServer) Create(context.Context, *SessionCreateRequest) (*SessionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSessionServer) GetEmail(context.Context, *SessionGetEmailRequest) (*SessionGetEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmail not implemented")
}
func (UnimplementedSessionServer) UpdateLastActivityTime(context.Context, *SessionUpdateActivityRequest) (*SessionUpdateActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateLastActivityTime not implemented")
}
func (UnimplementedSessionServer) GetActive(context.Context, *SessionGetActiveRequest) (*SessionGetActiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActive not implemented")
}
func (UnimplementedSessionServer) mustEmbedUnimplementedSessionServer() {}

// UnsafeSessionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServer will
// result in compilation errors.
type UnsafeSessionServer interface {
	mustEmbedUnimplementedSessionServer()
}

func RegisterSessionServer(s grpc.ServiceRegistrar, srv SessionServer) {
	s.RegisterService(&Session_ServiceDesc, srv)
}

func _Session_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Session/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).Create(ctx, req.(*SessionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_GetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionGetEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).GetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Session/GetEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).GetEmail(ctx, req.(*SessionGetEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_UpdateLastActivityTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionUpdateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).UpdateLastActivityTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Session/UpdateLastActivityTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).UpdateLastActivityTime(ctx, req.(*SessionUpdateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Session_GetActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionGetActiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServer).GetActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.Session/GetActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServer).GetActive(ctx, req.(*SessionGetActiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Session_ServiceDesc is the grpc.ServiceDesc for Session service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Session_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.Session",
	HandlerType: (*SessionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Session_Create_Handler,
		},
		{
			MethodName: "GetEmail",
			Handler:    _Session_GetEmail_Handler,
		},
		{
			MethodName: "UpdateLastActivityTime",
			Handler:    _Session_UpdateLastActivityTime_Handler,
		},
		{
			MethodName: "GetActive",
			Handler:    _Session_GetActive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wakadoo.proto",
}
