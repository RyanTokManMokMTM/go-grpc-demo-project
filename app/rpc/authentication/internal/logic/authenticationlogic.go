package logic

import (
	"context"
	"errors"
	"go-rpc-todo-list_/app/common/cryptox"
	"go-rpc-todo-list_/app/common/ctxtool"
	"go-rpc-todo-list_/app/common/jwtx"
	"strings"
	"time"

	"go-rpc-todo-list_/app/rpc/authentication/internal/svc"
	"go-rpc-todo-list_/app/rpc/authentication/types/authentication"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthenticationLogic) Authentication(in *authentication.AuthenticationReq) (*authentication.AuthenticationResp, error) {
	// todo: add your logic here and delete this line
	logx.Info(in)
	u, err := l.svcCtx.DAO.FindUserByUserName(in.Username, l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	encryptedPassword := cryptox.PasswordEncrypt(in.Password, l.svcCtx.Config.Salt)
	if strings.Compare(encryptedPassword, u.Password) != 0 {
		return nil, errors.New("password incorrect")
	}

	now := time.Now().Unix()
	exp := l.svcCtx.Config.JWTAuth.AccessExpire

	payLoad := map[string]interface{}{
		ctxtool.CTXJWTUserID: u.ID,
	}

	token, err := jwtx.GenerateToken(now, exp, l.svcCtx.Config.JWTAuth.AccessSecret, payLoad)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &authentication.AuthenticationResp{
		Token: token,
	}, nil
}
