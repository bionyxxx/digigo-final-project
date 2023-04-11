package controllers

import (
	"Final_Project/configs"
	"Final_Project/helpers"
	"Final_Project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	db := configs.GetDB()
	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	if user.Email == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Email or Password is required",
		})
		return
	}
	passwordInput := user.Password
	err = db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email or Password",
		})
		return
	}

	isValid := helpers.CheckBcrypt([]byte(user.Password), []byte(passwordInput))

	if !isValid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid Email or Password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Username, user.Email)

	c.JSON(200, gin.H{
		"message": "Login Success",
		"token":   token,
	})
}

func Register(c *gin.Context) {
	db := configs.GetDB()
	user := models.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	err = db.Debug().Create(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "Bad Request",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
		"message": "Register Success",
	})
}
