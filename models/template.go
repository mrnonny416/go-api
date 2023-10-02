package models

import (
	"gorm.io/gorm"
	"time"
)

type Template struct {
	ID       *uint   `gorm:"primary_key; auto_increment; index; " json:"templateId,omitempty" `
	Name     *string `json:"name"`
	Type     *uint   `json:"type"`
	Daystart *uint   `json:"daystart"`
	Dayend   *uint   `json:"dayend"`
	Datestart *string `json:"datestart"`
	DateEnd   *string `json:"dateend"`
	Timestart *string `json:"timestart"`
	TimeEnd   *string `json:"timeend"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
