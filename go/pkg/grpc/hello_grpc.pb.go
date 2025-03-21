// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: hello.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GreetingService_Hello_FullMethodName             = "/myapp.GreetingService/Hello"
	GreetingService_HelloServerStream_FullMethodName = "/myapp.GreetingService/HelloServerStream"
	GreetingService_HelloClientStream_FullMethodName = "/myapp.GreetingService/HelloClientStream"
	GreetingService_HelloBiStreams_FullMethodName    = "/myapp.GreetingService/HelloBiStreams"
)

// GreetingServiceClient is the client API for GreetingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error)
	HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloResponse], error)
	HelloBiStreams(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error)
}

type greetingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingServiceClient(cc grpc.ClientConnInterface) GreetingServiceClient {
	return &greetingServiceClient{cc}
}

func (c *greetingServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, GreetingService_Hello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetingServiceClient) HelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[0], GreetingService_HelloServerStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetingService_HelloServerStreamClient = grpc.ServerStreamingClient[HelloResponse]

func (c *greetingServiceClient) HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[1], GreetingService_HelloClientStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetingService_HelloClientStreamClient = grpc.ClientStreamingClient[HelloRequest, HelloResponse]

func (c *greetingServiceClient) HelloBiStreams(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[HelloRequest, HelloResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GreetingService_ServiceDesc.Streams[2], GreetingService_HelloBiStreams_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[HelloRequest, HelloResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetingService_HelloBiStreamsClient = grpc.BidiStreamingClient[HelloRequest, HelloResponse]

// GreetingServiceServer is the server API for GreetingService service.
// All implementations must embed UnimplementedGreetingServiceServer
// for forward compatibility.
type GreetingServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloServerStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error
	HelloClientStream(grpc.ClientStreamingServer[HelloRequest, HelloResponse]) error
	HelloBiStreams(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error
	mustEmbedUnimplementedGreetingServiceServer()
}

// UnimplementedGreetingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGreetingServiceServer struct{}

func (UnimplementedGreetingServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedGreetingServiceServer) HelloServerStream(*HelloRequest, grpc.ServerStreamingServer[HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedGreetingServiceServer) HelloClientStream(grpc.ClientStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloClientStream not implemented")
}
func (UnimplementedGreetingServiceServer) HelloBiStreams(grpc.BidiStreamingServer[HelloRequest, HelloResponse]) error {
	return status.Errorf(codes.Unimplemented, "method HelloBiStreams not implemented")
}
func (UnimplementedGreetingServiceServer) mustEmbedUnimplementedGreetingServiceServer() {}
func (UnimplementedGreetingServiceServer) testEmbeddedByValue()                         {}

// UnsafeGreetingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingServiceServer will
// result in compilation errors.
type UnsafeGreetingServiceServer interface {
	mustEmbedUnimplementedGreetingServiceServer()
}

func RegisterGreetingServiceServer(s grpc.ServiceRegistrar, srv GreetingServiceServer) {
	// If the following call pancis, it indicates UnimplementedGreetingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GreetingService_ServiceDesc, srv)
}

func _GreetingService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetingServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetingService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetingServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetingService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetingServiceServer).HelloServerStream(m, &grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetingService_HelloServerStreamServer = grpc.ServerStreamingServer[HelloResponse]

func _GreetingService_HelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).HelloClientStream(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetingService_HelloClientStreamServer = grpc.ClientStreamingServer[HelloRequest, HelloResponse]

func _GreetingService_HelloBiStreams_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingServiceServer).HelloBiStreams(&grpc.GenericServerStream[HelloRequest, HelloResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GreetingService_HelloBiStreamsServer = grpc.BidiStreamingServer[HelloRequest, HelloResponse]

// GreetingService_ServiceDesc is the grpc.ServiceDesc for GreetingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myapp.GreetingService",
	HandlerType: (*GreetingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreetingService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServerStream",
			Handler:       _GreetingService_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloClientStream",
			Handler:       _GreetingService_HelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloBiStreams",
			Handler:       _GreetingService_HelloBiStreams_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}
