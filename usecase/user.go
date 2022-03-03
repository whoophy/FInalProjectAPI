package usecase

import (
	"finalproject/entity"
	"finalproject/infra"
	"finalproject/utils/helper"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := infra.GetDatabase()
	contentType := helper.GetContentType(c)
	_, _ = db, contentType
	User := entity.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Create(&User).Error
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
	db := infra.GetDatabase()
	contenType := helper.GetContentType(c)
	_, _ = db, contenType
	User := entity.User{}
	password := ""
	if contenType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}
	fmt.Println(User.Password)
	fmt.Println(password)
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
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	User := entity.User{}

	userId := uint(userData["id"].(float64))

	if ContentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	fmt.Println(userId)
	err := db.Model(&User).Where("id=?", userId).Updates(entity.User{
		Email:    User.Email,
		Username: User.Username,
	}).Error
	err = db.Debug().First(&User, userId).Error
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
	db := infra.GetDatabase()
	userData := c.MustGet("userData").(jwt.MapClaims)
	ContentType := helper.GetContentType(c)
	User := entity.User{}

	userId := uint(userData["id"].(float64))

	if ContentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}
	err := db.Debug().Where("id=?", userId).Delete(&User).Error
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
