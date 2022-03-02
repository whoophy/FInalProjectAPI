package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption   string    `json:"caption"`
	Photo_url string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~photo is required,url~invalid url format"`
	UserID    uint      `gorm:"not null" `
	User      *User     `json:"user,omitempty"`
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"comment,omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
