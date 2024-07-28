package database

import (
	"go-rpc-todo-list_/app/rpc/internal/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(source string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err := initTables(db); err != nil {
		panic(err)
	}
	return db
}

func initTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Task{}); err != nil {
		return err
	}

	return nil
}
