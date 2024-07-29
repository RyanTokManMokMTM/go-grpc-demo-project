package authentication

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-rpc-todo-list_/app/api/auth-api/internal/logic/authentication"
	"go-rpc-todo-list_/app/api/auth-api/internal/svc"
	"go-rpc-todo-list_/app/api/auth-api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := authentication.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
