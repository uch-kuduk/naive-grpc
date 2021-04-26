// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// SampleServerClient is the client API for SampleServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SampleServerClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	UnimplemetedMethod(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type sampleServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSampleServerClient(cc grpc.ClientConnInterface) SampleServerClient {
	return &sampleServerClient{cc}
}

func (c *sampleServerClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/api.SampleServer/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleServerClient) UnimplemetedMethod(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/api.SampleServer/UnimplemetedMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SampleServerServer is the server API for SampleServer service.
// All implementations must embed UnimplementedSampleServerServer
// for forward compatibility
type SampleServerServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	UnimplemetedMethod(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedSampleServerServer()
}

// UnimplementedSampleServerServer must be embedded to have forward compatible implementations.
type UnimplementedSampleServerServer struct {
}

func (UnimplementedSampleServerServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedSampleServerServer) UnimplemetedMethod(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnimplemetedMethod not implemented")
}
func (UnimplementedSampleServerServer) mustEmbedUnimplementedSampleServerServer() {}

// UnsafeSampleServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SampleServerServer will
// result in compilation errors.
type UnsafeSampleServerServer interface {
	mustEmbedUnimplementedSampleServerServer()
}

func RegisterSampleServerServer(s grpc.ServiceRegistrar, srv SampleServerServer) {
	s.RegisterService(&SampleServer_ServiceDesc, srv)
}

func _SampleServer_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServerServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SampleServer/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServerServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleServer_UnimplemetedMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServerServer).UnimplemetedMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.SampleServer/UnimplemetedMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServerServer).UnimplemetedMethod(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SampleServer_ServiceDesc is the grpc.ServiceDesc for SampleServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SampleServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SampleServer",
	HandlerType: (*SampleServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _SampleServer_Hello_Handler,
		},
		{
			MethodName: "UnimplemetedMethod",
			Handler:    _SampleServer_UnimplemetedMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s.proto",
}
