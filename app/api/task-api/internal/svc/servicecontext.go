package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-rpc-todo-list_/app/api/task-api/internal/config"
	"go-rpc-todo-list_/app/rpc/task/taskrpc"
)

type ServiceContext struct {
	Config  config.Config
	TaskRPC taskrpc.TaskRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TaskRPC: taskrpc.NewTaskRPC(zrpc.MustNewClient(c.TaskRPC)),
	}
}
