// proto/operator.proto
// protoc -I proto proto/operator.proto --go_out=./internal/gen/ --go_opt=paths=source_relative --go-grpc_out=./internal/gen/ --go-grpc_opt=paths=source_relative
// Версия ProtoBuf

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: operator.proto

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
	OperatorService_List_FullMethodName = "/operator.OperatorService/List"
)

// OperatorServiceClient is the client API for OperatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorServiceClient interface {
	// Список операторов доставки
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type operatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorServiceClient(cc grpc.ClientConnInterface) OperatorServiceClient {
	return &operatorServiceClient{cc}
}

func (c *operatorServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, OperatorService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServiceServer is the server API for OperatorService service.
// All implementations must embed UnimplementedOperatorServiceServer
// for forward compatibility
type OperatorServiceServer interface {
	// Список операторов доставки
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedOperatorServiceServer()
}

// UnimplementedOperatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperatorServiceServer struct {
}

func (UnimplementedOperatorServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOperatorServiceServer) mustEmbedUnimplementedOperatorServiceServer() {}

// UnsafeOperatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServiceServer will
// result in compilation errors.
type UnsafeOperatorServiceServer interface {
	mustEmbedUnimplementedOperatorServiceServer()
}

func RegisterOperatorServiceServer(s grpc.ServiceRegistrar, srv OperatorServiceServer) {
	s.RegisterService(&OperatorService_ServiceDesc, srv)
}

func _OperatorService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperatorService_ServiceDesc is the grpc.ServiceDesc for OperatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "operator.OperatorService",
	HandlerType: (*OperatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _OperatorService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operator.proto",
}
