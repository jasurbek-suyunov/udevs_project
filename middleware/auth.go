package middleware

import (
	"github.com/jasurbek-suyunov/udevs_project/helper"

	"github.com/gin-gonic/gin"
)
// Auth is a middleware to check if user is authenticated
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "No token found"})
			return
		}

		param, err := helper.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}
		c.Set("user_id", param.UserId)
		c.Next()
	}
}
