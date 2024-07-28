package logic

import (
	"context"

	"go-rpc-todo-list_/app/rpc/task/internal/svc"
	"go-rpc-todo-list_/app/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskInfoLogic {
	return &GetTaskInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskInfoLogic) GetTaskInfo(in *task.TaskInfoReq) (*task.TaskInfoResp, error) {
	// todo: add your logic here and delete this line
	t, err := l.svcCtx.DAO.FindTask(in.TaskId, l.ctx)
	if err != nil {
		return nil, err
	}
	return &task.TaskInfoResp{
		TaskId:   t.Uuid,
		TaskName: t.TaskName,
	}, nil
}
