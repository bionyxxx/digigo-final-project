package middlewares

import (
	"Final_Project/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": helpers.Ucfirst(err.Error()),
			})

			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
