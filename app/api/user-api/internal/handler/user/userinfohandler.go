package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-rpc-todo-list_/app/api/user-api/internal/logic/user"
	"go-rpc-todo-list_/app/api/user-api/internal/svc"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
