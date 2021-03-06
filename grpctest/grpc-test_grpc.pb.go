// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpctest

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// GrpcTestClient is the client API for GrpcTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcTestClient interface {
	// Sends a greeting
	Msg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error)
}

type grpcTestClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcTestClient(cc grpc.ClientConnInterface) GrpcTestClient {
	return &grpcTestClient{cc}
}

func (c *grpcTestClient) Msg(ctx context.Context, in *MsgRequest, opts ...grpc.CallOption) (*MsgReply, error) {
	out := new(MsgReply)
	err := c.cc.Invoke(ctx, "/grpctest.GrpcTest/Msg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcTestServer is the server API for GrpcTest service.
// All implementations must embed UnimplementedGrpcTestServer
// for forward compatibility
type GrpcTestServer interface {
	// Sends a greeting
	Msg(context.Context, *MsgRequest) (*MsgReply, error)
	mustEmbedUnimplementedGrpcTestServer()
}

// UnimplementedGrpcTestServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcTestServer struct {
}

func (UnimplementedGrpcTestServer) Msg(context.Context, *MsgRequest) (*MsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Msg not implemented")
}
func (UnimplementedGrpcTestServer) mustEmbedUnimplementedGrpcTestServer() {}

// UnsafeGrpcTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcTestServer will
// result in compilation errors.
type UnsafeGrpcTestServer interface {
	mustEmbedUnimplementedGrpcTestServer()
}

func RegisterGrpcTestServer(s grpc.ServiceRegistrar, srv GrpcTestServer) {
	s.RegisterService(&_GrpcTest_serviceDesc, srv)
}

func _GrpcTest_Msg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcTestServer).Msg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.GrpcTest/Msg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcTestServer).Msg(ctx, req.(*MsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctest.GrpcTest",
	HandlerType: (*GrpcTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Msg",
			Handler:    _GrpcTest_Msg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpctest/grpc-test.proto",
}
