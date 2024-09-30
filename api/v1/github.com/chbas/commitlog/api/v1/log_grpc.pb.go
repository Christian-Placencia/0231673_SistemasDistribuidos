// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: log.proto

package v1

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
	CommitLogService_AddEntry_FullMethodName = "/api_v1.CommitLogService/AddEntry"
	CommitLogService_GetEntry_FullMethodName = "/api_v1.CommitLogService/GetEntry"
)

// CommitLogServiceClient is the client API for CommitLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define a service for interacting with the CommitLog
type CommitLogServiceClient interface {
	// Add an entry to the log
	AddEntry(ctx context.Context, in *AddEntryRequest, opts ...grpc.CallOption) (*AddEntryResponse, error)
	// Get an entry from the log
	GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error)
}

type commitLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommitLogServiceClient(cc grpc.ClientConnInterface) CommitLogServiceClient {
	return &commitLogServiceClient{cc}
}

func (c *commitLogServiceClient) AddEntry(ctx context.Context, in *AddEntryRequest, opts ...grpc.CallOption) (*AddEntryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddEntryResponse)
	err := c.cc.Invoke(ctx, CommitLogService_AddEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commitLogServiceClient) GetEntry(ctx context.Context, in *GetEntryRequest, opts ...grpc.CallOption) (*GetEntryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEntryResponse)
	err := c.cc.Invoke(ctx, CommitLogService_GetEntry_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommitLogServiceServer is the server API for CommitLogService service.
// All implementations must embed UnimplementedCommitLogServiceServer
// for forward compatibility.
//
// Define a service for interacting with the CommitLog
type CommitLogServiceServer interface {
	// Add an entry to the log
	AddEntry(context.Context, *AddEntryRequest) (*AddEntryResponse, error)
	// Get an entry from the log
	GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error)
	mustEmbedUnimplementedCommitLogServiceServer()
}

// UnimplementedCommitLogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommitLogServiceServer struct{}

func (UnimplementedCommitLogServiceServer) AddEntry(context.Context, *AddEntryRequest) (*AddEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEntry not implemented")
}
func (UnimplementedCommitLogServiceServer) GetEntry(context.Context, *GetEntryRequest) (*GetEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}
func (UnimplementedCommitLogServiceServer) mustEmbedUnimplementedCommitLogServiceServer() {}
func (UnimplementedCommitLogServiceServer) testEmbeddedByValue()                          {}

// UnsafeCommitLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommitLogServiceServer will
// result in compilation errors.
type UnsafeCommitLogServiceServer interface {
	mustEmbedUnimplementedCommitLogServiceServer()
}

func RegisterCommitLogServiceServer(s grpc.ServiceRegistrar, srv CommitLogServiceServer) {
	// If the following call pancis, it indicates UnimplementedCommitLogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommitLogService_ServiceDesc, srv)
}

func _CommitLogService_AddEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitLogServiceServer).AddEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommitLogService_AddEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitLogServiceServer).AddEntry(ctx, req.(*AddEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommitLogService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommitLogServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommitLogService_GetEntry_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommitLogServiceServer).GetEntry(ctx, req.(*GetEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommitLogService_ServiceDesc is the grpc.ServiceDesc for CommitLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommitLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_v1.CommitLogService",
	HandlerType: (*CommitLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEntry",
			Handler:    _CommitLogService_AddEntry_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _CommitLogService_GetEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}

const (
	LogService_Produce_FullMethodName = "/api_v1.LogService/Produce"
	LogService_Consume_FullMethodName = "/api_v1.LogService/Consume"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for producing and consuming log records
type LogServiceClient interface {
	// Produce a log record
	Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error)
	// Consume a log record by offset
	Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) Produce(ctx context.Context, in *ProduceRequest, opts ...grpc.CallOption) (*ProduceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProduceResponse)
	err := c.cc.Invoke(ctx, LogService_Produce_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) Consume(ctx context.Context, in *ConsumeRequest, opts ...grpc.CallOption) (*ConsumeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConsumeResponse)
	err := c.cc.Invoke(ctx, LogService_Consume_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility.
//
// Service for producing and consuming log records
type LogServiceServer interface {
	// Produce a log record
	Produce(context.Context, *ProduceRequest) (*ProduceResponse, error)
	// Consume a log record by offset
	Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogServiceServer struct{}

func (UnimplementedLogServiceServer) Produce(context.Context, *ProduceRequest) (*ProduceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedLogServiceServer) Consume(context.Context, *ConsumeRequest) (*ConsumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}
func (UnimplementedLogServiceServer) testEmbeddedByValue()                    {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_Produce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProduceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Produce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_Produce_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Produce(ctx, req.(*ProduceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_Consume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).Consume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_Consume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).Consume(ctx, req.(*ConsumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_v1.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Produce",
			Handler:    _LogService_Produce_Handler,
		},
		{
			MethodName: "Consume",
			Handler:    _LogService_Consume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "log.proto",
}
