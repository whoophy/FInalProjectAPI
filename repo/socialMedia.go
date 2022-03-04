package repo

import (
	"finalproject/entity"
	"finalproject/infra"

	"gorm.io/gorm"
)

func CreateSocialMediaRepo(SocialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	db := infra.GetDatabase()
	err := db.Debug().Create(&SocialMedia).Error
	return SocialMedia, err
}

func GetSocialMediaRepo(SocialMedia []entity.SocialMedia) ([]entity.SocialMedia, error) {
	db := infra.GetDatabase()
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "username", "email")
	}).Find(&SocialMedia).Error
	return SocialMedia, err
}

func UpdateSocialMediaRepo(SocialMedia entity.SocialMedia, socialmediaId uint) (entity.SocialMedia, error) {
	db := infra.GetDatabase()
	err := db.Model(&SocialMedia).Where("id=?", socialmediaId).Updates(entity.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
	}).Error
	err = db.Debug().Preload("User").First(&SocialMedia, socialmediaId).Error
	return SocialMedia, err
}

func DeleteSocialMediaRepo(SocialMedia entity.SocialMedia, socialmediaId uint) (entity.SocialMedia, error) {
	db := infra.GetDatabase()
	err := db.Debug().Where("id=?", socialmediaId).Delete(&SocialMedia).Error
	return SocialMedia, err
}
