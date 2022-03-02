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

func CreateComment(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)

	Comment := entity.Comment{}
	userID := uint(userData["id"].(float64))

	if contenType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserID = userID
	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":         Comment.ID,
		"message":    Comment.Message,
		"photo_id":   Comment.PhotoID,
		"user_id":    Comment.UserID,
		"created_at": Comment.CreatedAt,
	})
}

func GetComment(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)

	Comment := []entity.Comment{}
	userID := uint(userData["id"].(float64))
	_ = userID
	if contenType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "username", "email")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("ID", "title", "caption", "photo_url", "user_id")
	}).Find(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Comment)
}

func UpdateComment(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Comment := entity.Comment{}
	Photo := entity.Photo{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserID = userId
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id=?", commentId).Updates(entity.Comment{
		Message: Comment.Message,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if err := db.Debug().Find(&Photo).Where("id=?", Comment.PhotoID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

func DeleteComment(c *gin.Context) {
	db := infra.GetDatabase()
	_ = c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Comment := entity.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	err := db.Debug().Where("id=?", commentId).Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been sucesfully deleted",
	})
}
