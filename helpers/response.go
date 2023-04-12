package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error":   err,
		"message": "Internal Server Error",
	})
}

func ResponseNotFound(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error":   err,
		"message": "Not Found",
	})
}

func ResponseBadRequest(c *gin.Context, err string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error":   err,
		"message": "Bad Request",
	})
}

func ResponseBadRequestWithMessage(c *gin.Context, err string, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error":   err,
		"message": message,
	})
}

func ResponseCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func ResponseOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
