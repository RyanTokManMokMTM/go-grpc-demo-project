package logic

import (
	"context"

	"go-rpc-todo-list_/app/rpc/user/internal/svc"
	"go-rpc-todo-list_/app/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	//TODO: Get User Info with in param
	u, err := l.svcCtx.DAO.FindUser(in.UserId, l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// Get User Information from database
	// Check error
	// Return error
	return &user.UserInfoResp{
		UserName: u.UserName,
	}, nil
}
