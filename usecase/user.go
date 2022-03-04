package usecase

import (
	"finalproject/entity"
	"finalproject/repo"
	"finalproject/utils/helper"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	contentType := helper.GetContentType(c)
	_ = contentType
	User := entity.User{}
	var err error
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User, err = repo.CreateUserRepo(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

func UserLogin(c *gin.Context) {
	contentType := helper.GetContentType(c)
	_ = contentType
	User := entity.User{}
	var err error
	password := ""
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password = User.Password
	User, err = repo.UserLoginRepo(User)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	comparePass := helper.ComparePass([]byte(User.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "unauthorized",
			"message": "Invalid email/password 2",
		})
		return
	}
	token := helper.GenerateToken(User.ID, User.Email)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	_ = contentType
	User := entity.User{}
	var err error
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User, err = repo.UpdateUserRepo(User, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":        User.ID,
		"email":     User.Email,
		"username":  User.Username,
		"age":       User.Age,
		"update_at": User.UpdatedAt,
	})
}

func DeleteUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)
	_ = contentType
	User := entity.User{}
	var err error
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	User, err = repo.DeleteUserRepo(User, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been sucessfully deleted",
	})
}
