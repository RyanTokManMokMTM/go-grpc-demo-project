package user

import (
	"context"
	"errors"
	"go-rpc-todo-list_/app/api/user-api/internal/svc"
	"go-rpc-todo-list_/app/api/user-api/internal/types"
	"go-rpc-todo-list_/app/common/ctxtool"
	"go-rpc-todo-list_/app/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	if userID == 0 {
		return nil, errors.New("user id is missing")
	}

	u, err := l.svcCtx.UserRPC.GetUserInfo(l.ctx, &user.UserInfoReq{
		UserId: uint32(userID),
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResp{
		Id:       uint(u.Id),
		UserName: u.UserName,
	}, nil
}
