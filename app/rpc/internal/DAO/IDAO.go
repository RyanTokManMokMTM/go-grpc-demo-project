package DAO

import (
	"context"
	"go-rpc-todo-list_/app/rpc/internal/database/models"
)

type IDAO interface {
	CreateUser(userName, password string, ctx context.Context) error
	FindUser(userId string, ctx context.Context) (*models.User, error)
	FindUserByUserName(userName string, ctx context.Context) (*models.User, error)

	CreateTask(name string, uid uint, ctx context.Context) (*models.Task, error)
	FindTask(taskId string, ctx context.Context) (*models.Task, error)
}
