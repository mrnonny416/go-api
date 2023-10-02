package models

import (
	"time"

	"gorm.io/gorm"
)

type NodeMCU struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Millis    string    `json:"millis"`
	Timestamp time.Time `json:"timestamp"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
