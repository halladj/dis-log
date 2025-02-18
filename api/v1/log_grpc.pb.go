// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: api/v1/log.proto

package log_1v

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
	Log_Produce_FullMethodName       = "/log.v1.Log/Produce"
	Log_Consume_FullMethodName       = "/log.v1.Log/Consume"
	Log_ConsumeStream_FullMethodName = "/log.v1.Log/ConsumeStream"
	Log_ProduceStream_FullMethodName = "/log.v1.Log/ProduceStream"
	Log_GetServers_FullMethodName    = "/log.v1.Log/GetServers"
)

// LogClient is the client API for Log service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogClient interface {
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
	ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ConsumeResponse], error)
	ProduceStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ProduceRequest, ProduceResponse], error)
	GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error)
}

type logClient struct {
	cc grpc.ClientConnInterface
}

func NewLogClient(cc grpc.ClientConnInterface) LogClient {
	return &logClient{cc}
}

func (c *logClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, Log_Produce_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, Log_Consume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logClient) ConsumeStream(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ConsumeResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Log_ServiceDesc.Streams[0], Log_ConsumeStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ConsumeRequest, ConsumeResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Log_ConsumeStreamClient = grpc.ServerStreamingClient[ConsumeResponse]

func (c *logClient) ProduceStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ProduceRequest, ProduceResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Log_ServiceDesc.Streams[1], Log_ProduceStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ProduceRequest, ProduceResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Log_ProduceStreamClient = grpc.BidiStreamingClient[ProduceRequest, ProduceResponse]

func (c *logClient) GetServers(ctx context.Context, in *GetServersRequest, opts ...grpc.CallOption) (*GetServersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetServersResponse)
	err := c.cc.Invoke(ctx, Log_GetServers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServer is the server API for Log service.
// All implementations must embed UnimplementedLogServer
// for forward compatibility.
type LogServer interface {
	Produce(context.Context, *ProduceRequest) (*ProduceResponse, error)
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	ConsumeStream(*ConsumeRequest, grpc.ServerStreamingServer[ConsumeResponse]) error
	ProduceStream(grpc.BidiStreamingServer[ProduceRequest, ProduceResponse]) error
	GetServers(context.Context, *GetServersRequest) (*GetServersResponse, error)
	mustEmbedUnimplementedLogServer()
}

// UnimplementedLogServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServer struct{}

func (UnimplementedLogServer) Produce(context.Context, *ProduceRequest) (*ProduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedLogServer) Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedLogServer) ConsumeStream(*ConsumeRequest, grpc.ServerStreamingServer[ConsumeResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeStream not implemented")
}
func (UnimplementedLogServer) ProduceStream(grpc.BidiStreamingServer[ProduceRequest, ProduceResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ProduceStream not implemented")
}
func (UnimplementedLogServer) GetServers(context.Context, *GetServersRequest) (*GetServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (UnimplementedLogServer) mustEmbedUnimplementedLogServer() {}
func (UnimplementedLogServer) testEmbeddedByValue()             {}

// UnsafeLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServer will
// result in compilation errors.
type UnsafeLogServer interface {
	mustEmbedUnimplementedLogServer()
}

func RegisterLogServer(s grpc.ServiceRegistrar, srv LogServer) {
	// If the following call pancis, it indicates UnimplementedLogServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Log_ServiceDesc, srv)
}

func _Log_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Produce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_Consume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Log_ConsumeStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogServer).ConsumeStream(m, &grpc.GenericServerStream[ConsumeRequest, ConsumeResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Log_ConsumeStreamServer = grpc.ServerStreamingServer[ConsumeResponse]

func _Log_ProduceStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LogServer).ProduceStream(&grpc.GenericServerStream[ProduceRequest, ProduceResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Log_ProduceStreamServer = grpc.BidiStreamingServer[ProduceRequest, ProduceResponse]

func _Log_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Log_GetServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).GetServers(ctx, req.(*GetServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Log_ServiceDesc is the grpc.ServiceDesc for Log service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Log_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "log.v1.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _Log_Produce_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _Log_Consume_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _Log_GetServers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeStream",
			Handler:       _Log_ConsumeStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProduceStream",
			Handler:       _Log_ProduceStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/v1/log.proto",
}
