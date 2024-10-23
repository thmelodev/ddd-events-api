package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	Id        string     `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email" gorm:"not null"`
	Password  string     `json:"password" gorm:"not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (user *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Id == "" {
		user.Id = uuid.New().String()
	}
	return
}

func (*UserModel) TableName() string {
	return "users"
}
