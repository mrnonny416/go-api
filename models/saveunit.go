package models

import (
	"time"

	"gorm.io/gorm"
)

type Saveunit struct {
	ID     *uint      `gorm:"primary_key; auto_increment; index;" json:"save_uintId"`
	BfUnit *float32   `json:"bfunit"`
	AtUnit *float32   `json:"atunit"`
	Unit   *float32   `json:"unit"`
	Date   *time.Time `json:"date"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Requnit struct {
	Date *time.Time `json:"date"`
	Unit *float32   `json:"unit"`
}
