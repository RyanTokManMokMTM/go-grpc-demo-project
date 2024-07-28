package svc

import (
	"go-rpc-todo-list_/app/rpc/internal/DAO"
	"go-rpc-todo-list_/app/rpc/internal/database"
	"go-rpc-todo-list_/app/rpc/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DAO    DAO.IDAO
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := database.InitDatabase(c.MySql.Datasource)
	dao := DAO.NewDAO(db)
	return &ServiceContext{
		Config: c,
		DAO:    dao,
	}
}
