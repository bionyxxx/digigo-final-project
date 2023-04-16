package controllers

import (
	"Final_Project/helpers"
	"Final_Project/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (h HttpServer) GetAllSocialMedia(c *gin.Context) {
	var socialMedia []models.SocialMedia

	res, err := h.app.GetAllSocialMedia(socialMedia)

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All social media successfully retrieved",
		"data":    res,
	})
}

func (h HttpServer) GetSocialMedia(c *gin.Context) {
	socialMediaId, err := strconv.Atoi(c.Param("social_media_id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), "ID must be a number")
		return
	}
	res, err := h.app.GetSocialMedia(uint(socialMediaId))

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
		"message": "Social media successfully retrieved",
		"data":    res,
	})
}

func (h HttpServer) DeleteSocialMedia(c *gin.Context) {

	var socialMedia models.SocialMedia
	socialMedia.UserID = uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	err := h.app.DeleteSocialMedia(socialMedia)

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Social media successfully deleted",
	})
}

func (h HttpServer) UpdateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	socialMedia.UserID = uint(c.MustGet("userData").(jwt.MapClaims)["id"].(float64))

	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	res, err := h.app.UpdateSocialMedia(socialMedia)

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Social media successfully updated",
		"data":    res,
	})
}

func (h HttpServer) CreateSocialMedia(c *gin.Context) {
	socialMedia := models.SocialMedia{}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socialMedia.UserID = userID

	err := c.ShouldBindJSON(&socialMedia)
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	res, err := h.app.CreateSocialMedia(socialMedia)

	if err != nil {
		if err.Error() == "already" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "Social media already exists",
			})
			return
		}
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseCreated(c, gin.H{
		"message": "Social media successfully created",
		"data":    res,
	})

}
