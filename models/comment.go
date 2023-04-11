package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null;type:text" json:"message" valid:"required~Message is required"`
	PhotoID uint   `json:"photo_id"`
	Photo   *Photo `json:"photo,omitempty"`
	UserID  uint   `json:"user_id"`
	User    *User  `json:"user,omitempty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return
}
