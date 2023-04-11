package controllers

import (
	"Final_Project/configs"
	"Final_Project/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetAllSocialMedia(c *gin.Context) {
	db := configs.GetDB()
	var socialMedia []models.SocialMedia

	db.Find(&socialMedia)

	c.JSON(http.StatusOK, gin.H{
		"message": "All social media successfully retrieved",
		"data":    socialMedia,
	})
}

func GetSocialMedia(c *gin.Context) {
	db := configs.GetDB()
	var socialMedia models.SocialMedia

	if err := db.Preload("User").First(&socialMedia).Error; err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Social media successfully retrieved",
		"data":    socialMedia,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := configs.GetDB()
	var socialMedia models.SocialMedia

	if err := db.First(&socialMedia).Error; err != nil {
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

	if err := db.Delete(&socialMedia).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete social media",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Social media successfully deleted",
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := configs.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
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

	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	err = db.Save(&socialMedia).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update social media",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Social media successfully updated",
		"data":    socialMedia,
	})
}

func CreateSocialMedia(c *gin.Context) {
	db := configs.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	socialMedia := models.SocialMedia{}

	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	socialMedia.UserID = uint(userData["id"].(float64))

	if err := db.First(&socialMedia, "user_id = ?", socialMedia.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Create(&socialMedia).Error
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":   err.Error(),
					"message": "Bad Request",
				})
				return
			}
			c.JSON(http.StatusCreated, gin.H{
				"message": "Social media successfully created",
				"data":    socialMedia,
			})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "Bad Request",
			})
			return
		}
	} else {
		c.AbortWithStatusJSON(http.StatusCreated, gin.H{
			"message": "Social media already exists",
			"data":    socialMedia,
		})
		return
	}

}
