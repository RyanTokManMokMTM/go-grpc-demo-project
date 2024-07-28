package main

import (
	"flag"
	"fmt"

	"go-rpc-todo-list_/app/rpc/authentication/internal/config"
	"go-rpc-todo-list_/app/rpc/authentication/internal/server"
	"go-rpc-todo-list_/app/rpc/authentication/internal/svc"
	"go-rpc-todo-list_/app/rpc/authentication/types/authentication"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/authentication.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		authentication.RegisterAuthRPCServer(grpcServer, server.NewAuthRPCServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
