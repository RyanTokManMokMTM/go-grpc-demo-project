package main

import (
	"flag"
	"fmt"

	"go-rpc-todo-list_/app/rpc/task/internal/config"
	"go-rpc-todo-list_/app/rpc/task/internal/server"
	"go-rpc-todo-list_/app/rpc/task/internal/svc"
	"go-rpc-todo-list_/app/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/task.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		task.RegisterTaskRPCServer(grpcServer, server.NewTaskRPCServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
