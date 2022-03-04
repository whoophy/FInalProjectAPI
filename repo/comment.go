package repo

import (
	"finalproject/entity"
	"finalproject/infra"

	"gorm.io/gorm"
)

func CreateCommentRepo(Comment entity.Comment) (entity.Comment, error) {
	db := infra.GetDatabase()
	err := db.Debug().Create(&Comment).Error
	return Comment, err
}

func GetCommentRepo(Comment []entity.Comment) ([]entity.Comment, error) {
	db := infra.GetDatabase()
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "username", "email")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "title", "caption", "photo_url", "user_id")
	}).Find(&Comment).Error
	return Comment, err
}

func UpdateCommentRepo(Comment entity.Comment, commentId uint) (entity.Comment, error) {
	db := infra.GetDatabase()
	err := db.Model(&Comment).Where("id=?", commentId).Updates(entity.Comment{
		Message: Comment.Message,
	}).Error
	err = db.Debug().First(&Comment, commentId).Error
	return Comment, err
}

func DeleteCommentRepo(Comment entity.Comment, commentId uint) (entity.Comment, error) {
	db := infra.GetDatabase()
	err := db.Debug().Where("id=?", commentId).Delete(&Comment).Error
	return Comment, err
}
