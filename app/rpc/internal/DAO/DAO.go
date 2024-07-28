package DAO

import (
	"context"
	"go-rpc-todo-list_/app/rpc/internal/database/models"
	"gorm.io/gorm"
)

var _ IDAO = (*DAO)(nil)

type DAO struct {
	db *gorm.DB
}

func NewDAO(db *gorm.DB) IDAO {
	return &DAO{
		db: db,
	}
}

func (D *DAO) CreateUser(userName, password string, ctx context.Context) error {
	//TODO implement me
	u := &models.User{
		UserName: userName,
		Password: password,
	}
	return u.CreateUser(D.db, ctx)
}

func (D *DAO) FindUser(userId string, ctx context.Context) (*models.User, error) {
	//TODO implement me
	u := &models.User{
		Uuid: userId,
	}
	if err := u.FindOne(D.db, ctx); err != nil {
		return nil, err
	}
	return u, nil
}

func (D *DAO) FindUserByUserName(userName string, ctx context.Context) (*models.User, error) {
	//TODO implement me
	u := &models.User{
		UserName: userName,
	}

	if err := u.FindOne(D.db, ctx); err != nil {
		return nil, err
	}
	return u, nil
}

func (D *DAO) CreateTask(name string, uid uint, ctx context.Context) (*models.Task, error) {
	//TODO implement me
	t := &models.Task{
		TaskName: name,
		UserID:   uid,
	}
	if err := t.CreateTask(D.db, ctx); err != nil {
		return nil, err
	}
	return t, nil
}

func (D *DAO) FindTask(taskId string, ctx context.Context) (*models.Task, error) {
	//TODO implement me
	t := &models.Task{
		Uuid: taskId,
	}
	if err := t.FindOne(D.db, ctx); err != nil {
		return nil, err
	}
	return t, nil
}
