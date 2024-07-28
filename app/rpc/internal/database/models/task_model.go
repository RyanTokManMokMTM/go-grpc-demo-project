package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	Uuid     string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	TaskName string `gorm:"types:varchar(64)"`
	UserID   uint   `gorm:"not null;index'"`
	UserInfo User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
}

func (t *Task) TableName() string {
	return "Task"
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	t.Uuid = uuid.New().String()
	return nil
}

func (t *Task) CreateTask(tx *gorm.DB, ctx context.Context) error {
	return tx.Debug().WithContext(ctx).Create(&t).Error
}

func (t *Task) FindOne(tx *gorm.DB, ctx context.Context) error {
	return tx.Debug().WithContext(ctx).Find(&t).Error
}
