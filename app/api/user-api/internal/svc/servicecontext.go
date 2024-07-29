package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-rpc-todo-list_/app/api/user-api/internal/config"
	"go-rpc-todo-list_/app/rpc/user/userrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC userrpc.UserRPC //injected UserRPC client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: userrpc.NewUserRPC(zrpc.MustNewClient(c.UserRPC)),
	}
}
