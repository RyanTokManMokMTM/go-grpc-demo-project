package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-rpc-todo-list_/app/api/auth-api/internal/config"
	"go-rpc-todo-list_/app/rpc/authentication/authrpc"
)

type ServiceContext struct {
	Config  config.Config
	AuthRPC authrpc.AuthRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		AuthRPC: authrpc.NewAuthRPC(zrpc.MustNewClient(c.AuthRPC)),
	}
}
