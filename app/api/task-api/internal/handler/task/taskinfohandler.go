package task

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-rpc-todo-list_/app/api/task-api/internal/logic/task"
	"go-rpc-todo-list_/app/api/task-api/internal/svc"
	"go-rpc-todo-list_/app/api/task-api/internal/types"
)

func TaskInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTaskInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := task.NewTaskInfoLogic(r.Context(), svcCtx)
		resp, err := l.TaskInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
