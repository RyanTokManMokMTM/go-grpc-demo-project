// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-rpc-todo-list_/app/rpc/user/internal/logic"
	"go-rpc-todo-list_/app/rpc/user/internal/svc"
	"go-rpc-todo-list_/app/rpc/user/types/user"
)

type UserRPCServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserRPCServer
}

func NewUserRPCServer(svcCtx *svc.ServiceContext) *UserRPCServer {
	return &UserRPCServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRPCServer) CreateUser(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserResp, error) {
	l := logic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserRPCServer) GetUserInfo(ctx context.Context, in *user.UserInfoReq) (*user.UserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}