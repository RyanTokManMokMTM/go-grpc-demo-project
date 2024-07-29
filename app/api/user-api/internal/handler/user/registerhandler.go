package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-rpc-todo-list_/app/api/user-api/internal/logic/user"
	"go-rpc-todo-list_/app/api/user-api/internal/svc"
	"go-rpc-todo-list_/app/api/user-api/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqigsterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
