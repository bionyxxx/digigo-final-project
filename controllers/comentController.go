package controllers

import (
	"Final_Project/configs"
	"Final_Project/helpers"
	"Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (h HttpServer) GetAllComments(c *gin.Context) {
	var comments []models.Comment

	res, err := h.app.GetAllComments(comments)

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "All photos successfully retrieved",
		"data":    res,
	})
}

func (h HttpServer) UpdateComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("comment_id"))

	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	var comment models.Comment
	comment.ID = uint(commentId)
	comment.UserID = uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	err = c.ShouldBindJSON(&comment)
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateComment(comment)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ResponseNotFound(c, err.Error())
			return
		} else {
			helpers.ResponseError(c, err.Error())
			return
		}
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Comment successfully updated",
		"data":    res,
	})
}

func (h HttpServer) GetComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	res, err := h.app.GetComment(uint(commentId))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ResponseNotFound(c, err.Error())
			return
		} else {
			helpers.ResponseError(c, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfully retrieved",
		"data":    res,
	})
}

func (h HttpServer) DeleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	err = h.app.DeleteComment(uint(commentId))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ResponseNotFound(c, err.Error())
			return
		} else {
			helpers.ResponseError(c, err.Error())
			return
		}
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Comment successfully deleted",
	})
}

func (h HttpServer) CreateComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoId, err := strconv.Atoi(c.Param("photo_id"))

	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	var comment models.Comment
	comment.UserID = uint(userData["id"].(float64))
	comment.PhotoID = uint(photoId)

	err = c.ShouldBindJSON(&comment)
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	var photo = models.Photo{}
	err = configs.DB.First(&photo, photoId).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	res, err := h.app.CreateComment(comment, uint(photoId))

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseCreated(c, gin.H{
		"message": "Comment successfully created",
		"data":    res,
	})
}
