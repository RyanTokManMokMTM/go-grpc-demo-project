package task

import (
	"context"
	"errors"
	"go-rpc-todo-list_/app/common/ctxtool"
	"go-rpc-todo-list_/app/rpc/task/taskrpc"

	"go-rpc-todo-list_/app/api/task-api/internal/svc"
	"go-rpc-todo-list_/app/api/task-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskInfoLogic {
	return &TaskInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskInfoLogic) TaskInfo(req *types.GetTaskInfoReq) (resp *types.GetTaskInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, errors.New("user id is missing")
	}

	t, err := l.svcCtx.TaskRPC.GetTaskInfo(l.ctx, &taskrpc.TaskInfoReq{
		TaskId: req.TaskId,
	})

	if err != nil {
		return nil, err
	}
	return &types.GetTaskInfoResp{
		TaskInfo: types.TaskInfo{
			TaskId:   t.TaskId,
			TaskName: t.TaskName,
		},
	}, nil
}
