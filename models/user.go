package models

import (
	"Final_Project/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string       `gorm:"not null;uniqueIndex" json:"username" valid:"required~Username is required"`
	Email       string       `gorm:"not null;uniqueIndex" json:"email" valid:"required~Email is required,email"`
	Password    string       `gorm:"not null" json:"password" valid:"required~Password is required,minstringlength(6)"`
	Age         int          `gorm:"not null;type:int" json:"age" valid:"required~Age is required,range(8|100)"`
	Photos      []Photo      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos,omitempty"`
	Comments    []Comment    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments,omitempty"`
	SocialMedia *SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}
	u.Password = helpers.BcryptHash(u.Password)

	return
}
