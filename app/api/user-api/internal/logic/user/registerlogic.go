package user

import (
	"context"
	"go-rpc-todo-list_/app/rpc/user/types/user"

	"go-rpc-todo-list_/app/api/user-api/internal/svc"
	"go-rpc-todo-list_/app/api/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.ReqigsterReq) error {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.UserRPC.CreateUser(l.ctx, &user.CreateUserReq{
		Username: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logx.Error(err)
		return err
	}

	return nil
}
