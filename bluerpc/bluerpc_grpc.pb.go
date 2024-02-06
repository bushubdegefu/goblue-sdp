// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: bluerpc/bluerpc.proto

package bluesdp_proto

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
	BlueSDPHeartBeatService_SayAlive_FullMethodName = "/BlueSDPHeartBeatService/SayAlive"
	BlueSDPHeartBeatService_GetList_FullMethodName  = "/BlueSDPHeartBeatService/GetList"
)

// BlueSDPHeartBeatServiceClient is the client API for BlueSDPHeartBeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlueSDPHeartBeatServiceClient interface {
	SayAlive(ctx context.Context, in *BlueHeartBeat, opts ...grpc.CallOption) (*BlueHeartBeatResponse, error)
	GetList(ctx context.Context, in *GetRegistredServiceList, opts ...grpc.CallOption) (*RespondRegistredServiceList, error)
}

type blueSDPHeartBeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlueSDPHeartBeatServiceClient(cc grpc.ClientConnInterface) BlueSDPHeartBeatServiceClient {
	return &blueSDPHeartBeatServiceClient{cc}
}

func (c *blueSDPHeartBeatServiceClient) SayAlive(ctx context.Context, in *BlueHeartBeat, opts ...grpc.CallOption) (*BlueHeartBeatResponse, error) {
	out := new(BlueHeartBeatResponse)
	err := c.cc.Invoke(ctx, BlueSDPHeartBeatService_SayAlive_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blueSDPHeartBeatServiceClient) GetList(ctx context.Context, in *GetRegistredServiceList, opts ...grpc.CallOption) (*RespondRegistredServiceList, error) {
	out := new(RespondRegistredServiceList)
	err := c.cc.Invoke(ctx, BlueSDPHeartBeatService_GetList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlueSDPHeartBeatServiceServer is the server API for BlueSDPHeartBeatService service.
// All implementations must embed UnimplementedBlueSDPHeartBeatServiceServer
// for forward compatibility
type BlueSDPHeartBeatServiceServer interface {
	SayAlive(context.Context, *BlueHeartBeat) (*BlueHeartBeatResponse, error)
	GetList(context.Context, *GetRegistredServiceList) (*RespondRegistredServiceList, error)
	mustEmbedUnimplementedBlueSDPHeartBeatServiceServer()
}

// UnimplementedBlueSDPHeartBeatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlueSDPHeartBeatServiceServer struct {
}

func (UnimplementedBlueSDPHeartBeatServiceServer) SayAlive(context.Context, *BlueHeartBeat) (*BlueHeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayAlive not implemented")
}
func (UnimplementedBlueSDPHeartBeatServiceServer) GetList(context.Context, *GetRegistredServiceList) (*RespondRegistredServiceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedBlueSDPHeartBeatServiceServer) mustEmbedUnimplementedBlueSDPHeartBeatServiceServer() {
}

// UnsafeBlueSDPHeartBeatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlueSDPHeartBeatServiceServer will
// result in compilation errors.
type UnsafeBlueSDPHeartBeatServiceServer interface {
	mustEmbedUnimplementedBlueSDPHeartBeatServiceServer()
}

func RegisterBlueSDPHeartBeatServiceServer(s grpc.ServiceRegistrar, srv BlueSDPHeartBeatServiceServer) {
	s.RegisterService(&BlueSDPHeartBeatService_ServiceDesc, srv)
}

func _BlueSDPHeartBeatService_SayAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlueHeartBeat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSDPHeartBeatServiceServer).SayAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlueSDPHeartBeatService_SayAlive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSDPHeartBeatServiceServer).SayAlive(ctx, req.(*BlueHeartBeat))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlueSDPHeartBeatService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegistredServiceList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlueSDPHeartBeatServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlueSDPHeartBeatService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlueSDPHeartBeatServiceServer).GetList(ctx, req.(*GetRegistredServiceList))
	}
	return interceptor(ctx, in, info, handler)
}

// BlueSDPHeartBeatService_ServiceDesc is the grpc.ServiceDesc for BlueSDPHeartBeatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlueSDPHeartBeatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlueSDPHeartBeatService",
	HandlerType: (*BlueSDPHeartBeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayAlive",
			Handler:    _BlueSDPHeartBeatService_SayAlive_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _BlueSDPHeartBeatService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bluerpc/bluerpc.proto",
}
