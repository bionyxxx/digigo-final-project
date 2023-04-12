package middlewares

import (
	"Final_Project/configs"
	"Final_Project/helpers"
	"Final_Project/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func PhotoAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := configs.GetDB()
		photoId, err := strconv.Atoi(c.Param("photo_id"))
		if err != nil {
			helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		product := models.Photo{}

		err = db.Select("user_id").First(&product, uint(photoId)).Error
		if err != nil {
			helpers.ResponseNotFound(c, err.Error())
			return
		}
		println(product.UserID, userID)
		if product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Not allowed to access this photo",
			})

			return
		}

		c.Next()
	}
}

func CommentAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := configs.GetDB()
		commentId, err := strconv.Atoi(c.Param("comment_id"))
		if err != nil {
			helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		comment := models.Comment{}

		err = db.Select("user_id").First(&comment, uint(commentId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Comment not found",
			})

			return
		}

		if comment.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Not allowed to access this comment",
			})

			return
		}

		c.Next()
	}
}

func SocialMediaAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := configs.GetDB()

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		socialMedia := models.SocialMedia{}

		if err := db.First(&socialMedia, "user_id = ?", uint(userData["id"].(float64))).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error":   err.Error(),
					"message": "Social media not found",
				})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "Bad Request",
				})
				return
			}
		}

		if socialMedia.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Not allowed to access this social media",
			})

			return
		}

		c.Next()
	}
}
