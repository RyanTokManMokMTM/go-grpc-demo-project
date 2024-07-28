package logic

import (
	"context"

	"go-rpc-todo-list_/app/rpc/task/internal/svc"
	"go-rpc-todo-list_/app/rpc/task/types/task"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskLogic) CreateTask(in *task.CreateTaskReq) (*task.CreateTaskResp, error) {
	// todo: add your logic here and delete this line
	t, err := l.svcCtx.DAO.CreateTask(in.TaskName, uint(in.UserId), l.ctx)
	if err != nil {
		return nil, err
	}
	return &task.CreateTaskResp{
		TaskId: t.Uuid,
	}, nil
}
