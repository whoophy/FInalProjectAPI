package repo

import (
	"finalproject/entity"
	"finalproject/infra"

	"gorm.io/gorm"
)

func CreatePhotoRepo(Photo entity.Photo) (entity.Photo, error) {
	db := infra.GetDatabase()
	err := db.Debug().Create(&Photo).Error
	return Photo, err
}

func GetPhotoRepo(Photo []entity.Photo) ([]entity.Photo, error) {
	db := infra.GetDatabase()
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "username", "email")
	}).Find(&Photo).Error
	return Photo, err
}

func UpdatePhotoRepo(Photo entity.Photo, photoId uint) (entity.Photo, error) {
	db := infra.GetDatabase()
	err := db.Model(&Photo).Where("id=?", photoId).Updates(entity.Photo{
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		Photo_url: Photo.Photo_url,
	}).Error
	err = db.Debug().First(&Photo, photoId).Error
	return Photo, err
}

func DeletePhotoRepo(Photo entity.Photo, photoId uint) (entity.Photo, error) {
	db := infra.GetDatabase()
	err := db.Debug().Where("id=?", photoId).Delete(&Photo).Error
	return Photo, err
}
