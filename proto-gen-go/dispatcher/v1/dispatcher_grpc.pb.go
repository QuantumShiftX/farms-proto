// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.3
// source: dispatcher/v1/dispatcher.proto

package v1

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
	DispatcherAsync_TestAsync_FullMethodName = "/dispatcher.v1.DispatcherAsync/TestAsync"
)

// DispatcherAsyncClient is the client API for DispatcherAsync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherAsyncClient interface {
	TestAsync(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
}

type dispatcherAsyncClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherAsyncClient(cc grpc.ClientConnInterface) DispatcherAsyncClient {
	return &dispatcherAsyncClient{cc}
}

func (c *dispatcherAsyncClient) TestAsync(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	out := new(DispatcherReply)
	err := c.cc.Invoke(ctx, DispatcherAsync_TestAsync_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherAsyncServer is the server API for DispatcherAsync service.
// All implementations must embed UnimplementedDispatcherAsyncServer
// for forward compatibility
type DispatcherAsyncServer interface {
	TestAsync(context.Context, *DispatcherReq) (*DispatcherReply, error)
	mustEmbedUnimplementedDispatcherAsyncServer()
}

// UnimplementedDispatcherAsyncServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherAsyncServer struct {
}

func (UnimplementedDispatcherAsyncServer) TestAsync(context.Context, *DispatcherReq) (*DispatcherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestAsync not implemented")
}
func (UnimplementedDispatcherAsyncServer) mustEmbedUnimplementedDispatcherAsyncServer() {}

// UnsafeDispatcherAsyncServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherAsyncServer will
// result in compilation errors.
type UnsafeDispatcherAsyncServer interface {
	mustEmbedUnimplementedDispatcherAsyncServer()
}

func RegisterDispatcherAsyncServer(s grpc.ServiceRegistrar, srv DispatcherAsyncServer) {
	s.RegisterService(&DispatcherAsync_ServiceDesc, srv)
}

func _DispatcherAsync_TestAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherAsyncServer).TestAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherAsync_TestAsync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherAsyncServer).TestAsync(ctx, req.(*DispatcherReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherAsync_ServiceDesc is the grpc.ServiceDesc for DispatcherAsync service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherAsync_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dispatcher.v1.DispatcherAsync",
	HandlerType: (*DispatcherAsyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestAsync",
			Handler:    _DispatcherAsync_TestAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatcher/v1/dispatcher.proto",
}

const (
	DispatcherTimer_CycleFertileTreeStatusCheck_FullMethodName = "/dispatcher.v1.DispatcherTimer/CycleFertileTreeStatusCheck"
	DispatcherTimer_CycleCropStatusCheck_FullMethodName        = "/dispatcher.v1.DispatcherTimer/CycleCropStatusCheck"
	DispatcherTimer_CycleCropStageUpdate_FullMethodName        = "/dispatcher.v1.DispatcherTimer/CycleCropStageUpdate"
	DispatcherTimer_CycleBulletinMsgCheck_FullMethodName       = "/dispatcher.v1.DispatcherTimer/CycleBulletinMsgCheck"
)

// DispatcherTimerClient is the client API for DispatcherTimer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherTimerClient interface {
	// 发财树状态检查
	CycleFertileTreeStatusCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
	// 农场作物状态检查
	CycleCropStatusCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
	// 农场作物状态更新
	CycleCropStageUpdate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
	// 公告消息检查
	CycleBulletinMsgCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error)
}

type dispatcherTimerClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherTimerClient(cc grpc.ClientConnInterface) DispatcherTimerClient {
	return &dispatcherTimerClient{cc}
}

func (c *dispatcherTimerClient) CycleFertileTreeStatusCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	out := new(DispatcherReply)
	err := c.cc.Invoke(ctx, DispatcherTimer_CycleFertileTreeStatusCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherTimerClient) CycleCropStatusCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	out := new(DispatcherReply)
	err := c.cc.Invoke(ctx, DispatcherTimer_CycleCropStatusCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherTimerClient) CycleCropStageUpdate(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	out := new(DispatcherReply)
	err := c.cc.Invoke(ctx, DispatcherTimer_CycleCropStageUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherTimerClient) CycleBulletinMsgCheck(ctx context.Context, in *DispatcherReq, opts ...grpc.CallOption) (*DispatcherReply, error) {
	out := new(DispatcherReply)
	err := c.cc.Invoke(ctx, DispatcherTimer_CycleBulletinMsgCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherTimerServer is the server API for DispatcherTimer service.
// All implementations must embed UnimplementedDispatcherTimerServer
// for forward compatibility
type DispatcherTimerServer interface {
	// 发财树状态检查
	CycleFertileTreeStatusCheck(context.Context, *DispatcherReq) (*DispatcherReply, error)
	// 农场作物状态检查
	CycleCropStatusCheck(context.Context, *DispatcherReq) (*DispatcherReply, error)
	// 农场作物状态更新
	CycleCropStageUpdate(context.Context, *DispatcherReq) (*DispatcherReply, error)
	// 公告消息检查
	CycleBulletinMsgCheck(context.Context, *DispatcherReq) (*DispatcherReply, error)
	mustEmbedUnimplementedDispatcherTimerServer()
}

// UnimplementedDispatcherTimerServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherTimerServer struct {
}

func (UnimplementedDispatcherTimerServer) CycleFertileTreeStatusCheck(context.Context, *DispatcherReq) (*DispatcherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleFertileTreeStatusCheck not implemented")
}
func (UnimplementedDispatcherTimerServer) CycleCropStatusCheck(context.Context, *DispatcherReq) (*DispatcherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleCropStatusCheck not implemented")
}
func (UnimplementedDispatcherTimerServer) CycleCropStageUpdate(context.Context, *DispatcherReq) (*DispatcherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleCropStageUpdate not implemented")
}
func (UnimplementedDispatcherTimerServer) CycleBulletinMsgCheck(context.Context, *DispatcherReq) (*DispatcherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CycleBulletinMsgCheck not implemented")
}
func (UnimplementedDispatcherTimerServer) mustEmbedUnimplementedDispatcherTimerServer() {}

// UnsafeDispatcherTimerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherTimerServer will
// result in compilation errors.
type UnsafeDispatcherTimerServer interface {
	mustEmbedUnimplementedDispatcherTimerServer()
}

func RegisterDispatcherTimerServer(s grpc.ServiceRegistrar, srv DispatcherTimerServer) {
	s.RegisterService(&DispatcherTimer_ServiceDesc, srv)
}

func _DispatcherTimer_CycleFertileTreeStatusCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherTimerServer).CycleFertileTreeStatusCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherTimer_CycleFertileTreeStatusCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherTimerServer).CycleFertileTreeStatusCheck(ctx, req.(*DispatcherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherTimer_CycleCropStatusCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherTimerServer).CycleCropStatusCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherTimer_CycleCropStatusCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherTimerServer).CycleCropStatusCheck(ctx, req.(*DispatcherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherTimer_CycleCropStageUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherTimerServer).CycleCropStageUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherTimer_CycleCropStageUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherTimerServer).CycleCropStageUpdate(ctx, req.(*DispatcherReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DispatcherTimer_CycleBulletinMsgCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatcherReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherTimerServer).CycleBulletinMsgCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DispatcherTimer_CycleBulletinMsgCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherTimerServer).CycleBulletinMsgCheck(ctx, req.(*DispatcherReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DispatcherTimer_ServiceDesc is the grpc.ServiceDesc for DispatcherTimer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DispatcherTimer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dispatcher.v1.DispatcherTimer",
	HandlerType: (*DispatcherTimerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CycleFertileTreeStatusCheck",
			Handler:    _DispatcherTimer_CycleFertileTreeStatusCheck_Handler,
		},
		{
			MethodName: "CycleCropStatusCheck",
			Handler:    _DispatcherTimer_CycleCropStatusCheck_Handler,
		},
		{
			MethodName: "CycleCropStageUpdate",
			Handler:    _DispatcherTimer_CycleCropStageUpdate_Handler,
		},
		{
			MethodName: "CycleBulletinMsgCheck",
			Handler:    _DispatcherTimer_CycleBulletinMsgCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dispatcher/v1/dispatcher.proto",
}
