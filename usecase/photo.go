package usecase

import (
	"finalproject/entity"
	"finalproject/infra"
	"finalproject/utils/helper"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UploadPhotos(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)

	Photo := entity.Photo{}
	userID := uint(userData["id"].(float64))

	if contenType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = userID
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.UserID,
		"created_at": Photo.CreatedAt,
	})
}

func GetPhoto(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)
	Photo := []entity.Photo{}
	_ = uint(userData["id"].(float64))

	if contenType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "username", "email")
	}).Find(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

func UpdatePhoto(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Photo := entity.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))
	userId := uint(userData["id"].(float64))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo.UserID = userId
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id=?", photoId).Updates(entity.Photo{
		Title:     Photo.Title,
		Caption:   Photo.Caption,
		Photo_url: Photo.Photo_url,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         Photo.ID,
		"title":      Photo.Title,
		"caption":    Photo.Caption,
		"photo_url":  Photo.Photo_url,
		"user_id":    Photo.UserID,
		"updated_at": Photo.CreatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	db := infra.GetDatabase()
	_ = c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Photo := entity.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	err := db.Debug().Where("id=?", photoId).Delete(&Photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been sucesfully deleted",
	})
}
