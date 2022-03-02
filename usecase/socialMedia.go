package usecase

import (
	"finalproject/entity"
	"finalproject/infra"
	"finalproject/utils/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ConnectSocialMedia(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)

	SocialMedia := entity.SocialMedia{}
	userID := uint(userData["id"].(float64))
	fmt.Println(userID)

	if contenType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	SocialMedia.UserID = userID
	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func GetSocialMedia(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contenType := helper.GetContentType(c)

	SocialMedia := []entity.SocialMedia{}
	userID := uint(userData["id"].(float64))
	_ = userID
	if contenType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	err := db.Debug().Preload("User").Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"social_media": SocialMedia,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	SocialMedia := entity.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))
	_ = uint(userData["id"].(float64))

	if ContentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Model(&SocialMedia).Where("id=?", socialmediaId).Updates(entity.SocialMedia{
		Name:           SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	if err := db.Debug().Preload("User").Find(&SocialMedia).Where("id=?", socialmediaId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := infra.GetDatabase()
	_ = c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	SocialMedia := entity.SocialMedia{}

	socialmediaId, _ := strconv.Atoi(c.Param("socialmediaId"))

	if ContentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}
	err := db.Debug().Where("id=?", socialmediaId).Delete(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been sucesfully deleted",
	})
}
