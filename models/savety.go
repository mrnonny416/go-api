package models

import (
	"gorm.io/gorm"
	"time"
)

type Savety struct {
	ID     *uint    `gorm:"primary_key; auto_increment; index;" json:"savetyId"`
	Volt   *float32 `json:"volt"`
	Status *bool    `gorm:"default:false" json:"status"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
