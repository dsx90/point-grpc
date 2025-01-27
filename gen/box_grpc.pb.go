// proto/box.proto
// protoc -I proto proto/box.proto --go_out=./internal/gen/ --go_opt=paths=source_relative --go-grpc_out=./internal/gen/ --go-grpc_opt=paths=source_relative
// Версия ProtoBuf

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: box.proto

package boxv1

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
	BoxService_List_FullMethodName  = "/box.BoxService/List"
	BoxService_Open_FullMethodName  = "/box.BoxService/Open"
	BoxService_Login_FullMethodName = "/box.BoxService/Login"
)

// BoxServiceClient is the client API for BoxService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoxServiceClient interface {
	// Список свободных ящиков
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Команда открытия ящика
	Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error)
	// Команда закрытия ящика
	Login(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
}

type boxServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoxServiceClient(cc grpc.ClientConnInterface) BoxServiceClient {
	return &boxServiceClient{cc}
}

func (c *boxServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, BoxService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) Open(ctx context.Context, in *OpenRequest, opts ...grpc.CallOption) (*OpenResponse, error) {
	out := new(OpenResponse)
	err := c.cc.Invoke(ctx, BoxService_Open_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boxServiceClient) Login(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, BoxService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoxServiceServer is the server API for BoxService service.
// All implementations must embed UnimplementedBoxServiceServer
// for forward compatibility
type BoxServiceServer interface {
	// Список свободных ящиков
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Команда открытия ящика
	Open(context.Context, *OpenRequest) (*OpenResponse, error)
	// Команда закрытия ящика
	Login(context.Context, *CloseRequest) (*CloseResponse, error)
	mustEmbedUnimplementedBoxServiceServer()
}

// UnimplementedBoxServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBoxServiceServer struct {
}

func (UnimplementedBoxServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBoxServiceServer) Open(context.Context, *OpenRequest) (*OpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedBoxServiceServer) Login(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBoxServiceServer) mustEmbedUnimplementedBoxServiceServer() {}

// UnsafeBoxServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoxServiceServer will
// result in compilation errors.
type UnsafeBoxServiceServer interface {
	mustEmbedUnimplementedBoxServiceServer()
}

func RegisterBoxServiceServer(s grpc.ServiceRegistrar, srv BoxServiceServer) {
	s.RegisterService(&BoxService_ServiceDesc, srv)
}

func _BoxService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_Open_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).Open(ctx, req.(*OpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoxService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoxService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxServiceServer).Login(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoxService_ServiceDesc is the grpc.ServiceDesc for BoxService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoxService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "box.BoxService",
	HandlerType: (*BoxServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BoxService_List_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _BoxService_Open_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _BoxService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "box.proto",
}
