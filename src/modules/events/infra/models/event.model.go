package models

import "time"

type Event struct {
	Id          string     `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name" gorm:"not null"`
	Description string     `json:"description" gorm:"not null"`
	Location    string     `json:"location" gorm:"not null"`
	DateTime    time.Time  `json:"date_time" gorm:"not null"`
	UserID      string     `json:"user_id" gorm:"not null"`
	CreatedAt   *time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Event) TableName() string {
	return "events"
}
