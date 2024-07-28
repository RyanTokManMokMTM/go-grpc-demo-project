// Code generated by goctl. DO NOT EDIT.
// Source: authentication.proto

package authrpc

import (
	"context"

	"go-rpc-todo-list_/app/rpc/authentication/types/authentication"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AuthenticationReq  = authentication.AuthenticationReq
	AuthenticationResp = authentication.AuthenticationResp

	AuthRPC interface {
		Authentication(ctx context.Context, in *AuthenticationReq, opts ...grpc.CallOption) (*AuthenticationResp, error)
	}

	defaultAuthRPC struct {
		cli zrpc.Client
	}
)

func NewAuthRPC(cli zrpc.Client) AuthRPC {
	return &defaultAuthRPC{
		cli: cli,
	}
}

func (m *defaultAuthRPC) Authentication(ctx context.Context, in *AuthenticationReq, opts ...grpc.CallOption) (*AuthenticationResp, error) {
	client := authentication.NewAuthRPCClient(m.cli.Conn())
	return client.Authentication(ctx, in, opts...)
}