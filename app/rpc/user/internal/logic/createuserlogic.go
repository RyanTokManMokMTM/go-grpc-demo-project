package logic

import (
	"context"
	"go-rpc-todo-list_/app/common/cryptox"

	"go-rpc-todo-list_/app/rpc/user/internal/svc"
	"go-rpc-todo-list_/app/rpc/user/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	logx.Info(in)
	//TODO: Find a user with 'in' params
	if _, err := l.svcCtx.DAO.FindUserByUserName(in.Username, l.ctx); err != nil {
		logx.Error(err)
		return nil, err
	}

	//TODO : Create User
	encryptedPassword := cryptox.PasswordEncrypt(in.Password, l.svcCtx.Config.Salt)
	err := l.svcCtx.DAO.CreateUser(in.Username, encryptedPassword, l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &user.CreateUserResp{}, nil
}
