package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Uuid     string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	UserName string `gorm:"types:varchar(64);not null;unique_index:idx_un"`
	Password string `gorm:"types:varchar(32);not null"`
	gorm.Model
}

func (u *User) TableName() string {
	return "User"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Uuid = uuid.New().String()
	return nil
}

func (u *User) CreateUser(tx *gorm.DB, ctx context.Context) error {
	return tx.Debug().WithContext(ctx).Create(&u).Error
}

func (u *User) FindOne(tx *gorm.DB, ctx context.Context) error {
	return tx.Debug().WithContext(ctx).Where(u).First(&u).Error
}
