package controllers

import (
	"Final_Project/helpers"
	"Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (h HttpServer) DeletePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	err = h.app.DeletePhoto(uint(photoId))

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
		"message": "Photo successfully deleted",
	})
}

func (h HttpServer) GetPhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	res, err := h.app.GetPhoto(uint(photoId))

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
		"message": "Photo successfully retrieved",
		"data":    res,
	})
}

func (h HttpServer) UpdatePhoto(c *gin.Context) {
	photoId, err := strconv.Atoi(c.Param("photo_id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}

	photo := models.Photo{}
	photo.ID = uint(photoId)
	photo.UserID = uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	err = c.ShouldBindJSON(&photo)
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdatePhoto(photo)

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
		"message": "Photo successfully updated",
		"data":    res,
	})
}

func (h HttpServer) GetAllPhotos(c *gin.Context) {
	var photos []models.Photo
	res, err := h.app.GetAllPhotos(photos)

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "All photos successfully retrieved",
		"data":    res,
	})
}

func (h HttpServer) CreatePhoto(c *gin.Context) {
	photo := models.Photo{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photo.UserID = userID

	err := c.ShouldBindJSON(&photo)
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}
	res, err := h.app.CreatePhoto(photo)

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseCreated(c, gin.H{
		"message": "Photo successfully created",
		"data":    res,
	})
}
