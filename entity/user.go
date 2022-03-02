package entity

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	//"github.com/jinzhu/gorm"
	"finalproject/utils/helper"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~invalid email format"`
	Password    string        `gorm:"not null" json:"password,omitempty" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age         uint          `gorm:"not null" json:"age,omitempty" form:"age" valid:"range(8|100)~You should over 8 year to register"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"photo,omitempty"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"comment,omitempty"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"socialmedia,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println(u.Password)
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helper.HashPass(u.Password)
	fmt.Println(u.Password)
	err = nil
	return
}
