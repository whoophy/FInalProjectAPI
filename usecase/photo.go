package usecase

import (
	"finalproject/entity"
	"finalproject/repo"
	"finalproject/utils/helper"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UploadPhotos(c *gin.Context) {
	var err error
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
	Photo, err = repo.CreatePhotoRepo(Photo)

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
	var err error
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)
	Photo := []entity.Photo{}
	_ = uint(userData["id"].(float64))

	if contenType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo, err = repo.GetPhotoRepo(Photo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

func UpdatePhoto(c *gin.Context) {
	var err error
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

	Photo, err = repo.UpdatePhotoRepo(Photo, uint(photoId))

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
		"updated_at": Photo.UpdatedAt,
	})
}

func DeletePhoto(c *gin.Context) {
	var err error
	_ = c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Photo := entity.Photo{}

	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}
	Photo, err = repo.DeletePhotoRepo(Photo, uint(photoId))
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
