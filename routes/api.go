package routes

import (
	"Final_Project/controllers"
	"Final_Project/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiInit(router *gin.Engine, server controllers.HttpServer) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.Register)
		authRouter.POST("/login", controllers.Login)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", server.GetAllPhotos)
		photoRouter.POST("/", server.CreatePhoto)
		photoRouter.POST("/:photo_id/comment", server.CreateComment)
		photoRouter.GET("/:photo_id", server.GetPhoto)
		photoRouter.PUT("/:photo_id", middlewares.PhotoAuth(), server.UpdatePhoto)
		photoRouter.DELETE("/:photo_id", middlewares.PhotoAuth(), server.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", server.GetAllComments)
		commentRouter.GET("/:comment_id", server.GetComment)
		commentRouter.DELETE("/:comment_id", middlewares.CommentAuth(), server.DeleteComment)
		commentRouter.PUT("/:comment_id", middlewares.CommentAuth(), server.UpdateComment)
	}

	socialMediaRouter := router.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", server.GetAllSocialMedia)
		socialMediaRouter.POST("/", server.CreateSocialMedia)
		socialMediaRouter.GET("/:social_media_id", server.GetSocialMedia)
		socialMediaRouter.PUT("/", middlewares.SocialMediaAuth(), server.UpdateSocialMedia)
		socialMediaRouter.DELETE("/", middlewares.SocialMediaAuth(), server.DeleteSocialMedia)
	}
}
