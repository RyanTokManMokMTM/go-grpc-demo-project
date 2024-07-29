package authentication

import (
	"context"
	"go-rpc-todo-list_/app/rpc/authentication/authrpc"

	"go-rpc-todo-list_/app/api/auth-api/internal/svc"
	"go-rpc-todo-list_/app/api/auth-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	token, err := l.svcCtx.AuthRPC.Authentication(l.ctx, &authrpc.AuthenticationReq{
		Username: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Token: token.Token,
	}, nil
}
