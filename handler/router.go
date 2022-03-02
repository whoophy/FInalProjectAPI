package handler

import (
	"finalproject/usecase"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", usecase.UserRegister)
		userRouter.POST("/login", usecase.UserLogin)
		userRouter.PUT("/", Authentication(), UserAuthorization(), usecase.UpdateUser)
		userRouter.DELETE("/", Authentication(), UserAuthorization(), usecase.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(Authentication())
		photoRouter.POST("/", usecase.UploadPhotos)
		photoRouter.PUT("/:photoId", PhotoAuthorization(), usecase.UpdatePhoto)
		photoRouter.GET("/", usecase.GetPhoto)
		photoRouter.DELETE("/:photoId", PhotoAuthorization(), usecase.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(Authentication())
		commentRouter.POST("/", usecase.CreateComment)
		commentRouter.GET("/", usecase.GetComment)
		commentRouter.PUT("/:commentId", CommentAuthorization(), usecase.UpdateComment)
		commentRouter.DELETE("/:commentId", CommentAuthorization(), usecase.DeleteComment)
	}
	socialmediaRouter := r.Group("/socialmedias")
	{
		socialmediaRouter.Use(Authentication())
		socialmediaRouter.POST("/", usecase.ConnectSocialMedia)
		socialmediaRouter.GET("/", usecase.GetSocialMedia)
		socialmediaRouter.PUT("/:socialmediaId", SocialMediaAuthorization(), usecase.UpdateSocialMedia)
	}
	return r
}
