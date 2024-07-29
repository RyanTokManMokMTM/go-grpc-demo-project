// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "go-rpc-todo-list_/app/api/user-api/internal/handler/user"
	"go-rpc-todo-list_/app/api/user-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userInfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}