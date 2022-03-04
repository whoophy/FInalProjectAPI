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

func CreateComment(c *gin.Context) {
	var err error
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
	Comment, err = repo.CreateCommentRepo(Comment)
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
	var err error
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
	Comment, err = repo.GetCommentRepo(Comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func UpdateComment(c *gin.Context) {
	var err error
	userData := c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Comment := entity.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userId := uint(userData["id"].(float64))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment.UserID = userId
	Comment.ID = uint(commentId)
	Comment, err = repo.UpdateCommentRepo(Comment, uint(commentId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	var err error
	_ = c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	Comment := entity.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if ContentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	Comment, err = repo.DeleteCommentRepo(Comment, uint(commentId))
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
