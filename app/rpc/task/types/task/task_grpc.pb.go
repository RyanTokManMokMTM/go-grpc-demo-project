// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: task.proto

package task

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
	TaskRPC_CreateTask_FullMethodName  = "/task.TaskRPC/CreateTask"
	TaskRPC_GetTaskInfo_FullMethodName = "/task.TaskRPC/GetTaskInfo"
)

// TaskRPCClient is the client API for TaskRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskRPCClient interface {
	CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskResp, error)
	GetTaskInfo(ctx context.Context, in *TaskInfoReq, opts ...grpc.CallOption) (*TaskInfoResp, error)
}

type taskRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskRPCClient(cc grpc.ClientConnInterface) TaskRPCClient {
	return &taskRPCClient{cc}
}

func (c *taskRPCClient) CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskResp, error) {
	out := new(CreateTaskResp)
	err := c.cc.Invoke(ctx, TaskRPC_CreateTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskRPCClient) GetTaskInfo(ctx context.Context, in *TaskInfoReq, opts ...grpc.CallOption) (*TaskInfoResp, error) {
	out := new(TaskInfoResp)
	err := c.cc.Invoke(ctx, TaskRPC_GetTaskInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskRPCServer is the server API for TaskRPC service.
// All implementations must embed UnimplementedTaskRPCServer
// for forward compatibility
type TaskRPCServer interface {
	CreateTask(context.Context, *CreateTaskReq) (*CreateTaskResp, error)
	GetTaskInfo(context.Context, *TaskInfoReq) (*TaskInfoResp, error)
	mustEmbedUnimplementedTaskRPCServer()
}

// UnimplementedTaskRPCServer must be embedded to have forward compatible implementations.
type UnimplementedTaskRPCServer struct {
}

func (UnimplementedTaskRPCServer) CreateTask(context.Context, *CreateTaskReq) (*CreateTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskRPCServer) GetTaskInfo(context.Context, *TaskInfoReq) (*TaskInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskInfo not implemented")
}
func (UnimplementedTaskRPCServer) mustEmbedUnimplementedTaskRPCServer() {}

// UnsafeTaskRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskRPCServer will
// result in compilation errors.
type UnsafeTaskRPCServer interface {
	mustEmbedUnimplementedTaskRPCServer()
}

func RegisterTaskRPCServer(s grpc.ServiceRegistrar, srv TaskRPCServer) {
	s.RegisterService(&TaskRPC_ServiceDesc, srv)
}

func _TaskRPC_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskRPCServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskRPC_CreateTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskRPCServer).CreateTask(ctx, req.(*CreateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskRPC_GetTaskInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskRPCServer).GetTaskInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskRPC_GetTaskInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskRPCServer).GetTaskInfo(ctx, req.(*TaskInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskRPC_ServiceDesc is the grpc.ServiceDesc for TaskRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskRPC",
	HandlerType: (*TaskRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskRPC_CreateTask_Handler,
		},
		{
			MethodName: "GetTaskInfo",
			Handler:    _TaskRPC_GetTaskInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "task.proto",
}
