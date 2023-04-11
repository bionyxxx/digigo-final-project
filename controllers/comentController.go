package controllers

import (
	"Final_Project/configs"
	"Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAllComments(c *gin.Context) {
	db := configs.GetDB()
	var comments []models.Comment

	db.Find(&comments)

	c.JSON(http.StatusOK, gin.H{
		"message": "All photos successfully retrieved",
		"data":    comments,
	})
}

func UpdateComment(c *gin.Context) {
	db := configs.GetDB()
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}

	var comment models.Comment

	err = db.First(&comment, commentId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Comment not found",
		})
		return
	}

	err = c.ShouldBindJSON(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	err = db.Save(&comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to update comment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfully updated",
		"data":    comment,
	})
}

func GetComment(c *gin.Context) {
	db := configs.GetDB()
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}

	var comment models.Comment

	err = db.Preload("Photo").Preload("User").First(&comment, commentId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Comment not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfully retrieved",
		"data":    comment,
	})
}

func DeleteComment(c *gin.Context) {
	db := configs.GetDB()
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}

	comment := models.Comment{}

	userData := c.MustGet("userData").(jwt.MapClaims)

	if userData["role"] == "admin" {
		err = db.First(&comment, commentId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Comment not found",
			})
			return
		}
	} else {
		err = db.Where("user_id = ?", userData["id"]).First(&comment, commentId).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   err.Error(),
				"message": "Comment not found",
			})
			return
		}
	}

	err = db.Delete(&comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "Failed to delete comment",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfully deleted",
	})
}

func CreateComment(c *gin.Context) {
	db := configs.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, err := strconv.Atoi(c.Param("photo_id"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "ID must be a number",
		})
		return
	}
	photo := models.Photo{}
	err = db.Preload("User").First(&photo, photoId).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "Photo not found",
		})
		return
	}

	var comment models.Comment

	err = c.ShouldBindJSON(&comment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	comment.UserID = uint(userData["id"].(float64))
	comment.PhotoID = uint(photoId)

	err = db.Preload("Photo").Create(&comment).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment successfully created",
		"data":    comment,
	})
}
