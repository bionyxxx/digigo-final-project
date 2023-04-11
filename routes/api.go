package routes

import (
	"Final_Project/controllers"
	"Final_Project/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiInit() *gin.Engine {
	router := gin.Default()

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/register", controllers.Register)
		authRouter.POST("/login", controllers.Login)
	}

	photoRouter := router.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhotos)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.POST("/:photo_id/comment", controllers.CreateComment)
		photoRouter.GET("/:photo_id", controllers.GetPhoto)
		photoRouter.PUT("/:photo_id", middlewares.PhotoAuth(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photo_id", middlewares.PhotoAuth(), controllers.DeletePhoto)
	}

	commentRouter := router.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComments)
		commentRouter.GET("/:comment_id", controllers.GetComment)
		commentRouter.DELETE("/:comment_id", middlewares.CommentAuth(), controllers.DeleteComment)
		commentRouter.PUT("/:comment_id", middlewares.CommentAuth(), controllers.UpdateComment)
	}

	socialMediaRouter := router.Group("/social-media")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/:social_media_id", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/", middlewares.SocialMediaAuth(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/", middlewares.SocialMediaAuth(), controllers.DeleteSocialMedia)
	}
	return router
}
