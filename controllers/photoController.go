package controllers

import (
	"Final_Project/configs"
	"Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeletePhoto(c *gin.Context) {
	db := configs.GetDB()
	photoId, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}

	photo := models.Photo{}

	err = db.First(&photo, photoId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Photo not found",
		})
		return
	}

	err = db.Delete(&photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete photo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo successfully deleted",
	})
}

func GetPhoto(c *gin.Context) {
	db := configs.GetDB()
	photoId, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}

	photo := models.Photo{}

	err = db.Preload("User").Preload("Comments").First(&photo, photoId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo successfully retrieved",
		"data":    photo,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := configs.GetDB()
	photoId, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}

	photo := models.Photo{}

	err = db.First(&photo, photoId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Photo not found",
		})
		return
	}

	err = c.ShouldBindJSON(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	err = db.Save(&photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo successfully updated",
		"data":    photo,
	})

}

func GetAllPhotos(c *gin.Context) {
	db := configs.GetDB()
	var photos []models.Photo

	db.Find(&photos)

	c.JSON(http.StatusOK, gin.H{
		"message": "All photos successfully retrieved",
		"data":    photos,
	})
}

func CreatePhoto(c *gin.Context) {
	db := configs.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	err := c.ShouldBindJSON(&photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	photo.UserID = userID

	err = db.Debug().Create(&photo).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Photo successfully created",
		"data":    photo,
	})
}
