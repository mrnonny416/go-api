package models

import (
	"gorm.io/gorm"
	"time"
)

type Channel struct {
	ID             *uint   `gorm:"primary_key; auto_increment; index;" json:"channelId"`
	Name           *string `json:"name"`
	Templatestatus *bool   `gorm:"default:false" json:"templatestatus"`
	TemplateId     *uint   ` json:"templateId"`
	Status         *bool   `gorm:"default:false" json:"status"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`

	Template *Template `gorm:"foreignkey:TemplateId" json:"template,omitempty"`
}
