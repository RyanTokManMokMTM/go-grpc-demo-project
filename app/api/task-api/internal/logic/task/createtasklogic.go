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

type CreateTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskLogic {
	return &CreateTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTaskLogic) CreateTask(req *types.CreateTaskReq) (resp *types.CreateTaskResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	if userId == 0 {
		return nil, errors.New("user id is missing")
	}

	t, err := l.svcCtx.TaskRPC.CreateTask(l.ctx, &taskrpc.CreateTaskReq{
		TaskName: req.TaskName,
		UserId:   int32(userId),
	})

	if err != nil {
		return nil, err
	}
	return &types.CreateTaskResp{
		TaskId: t.TaskId,
	}, nil
}
