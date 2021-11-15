package middleware

import (
	"logistics/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authLine, ok := c.Request.Header["Authorization"]
		logrus.Info("authLine", authLine)
		if !ok {
			c.AbortWithStatus(403)
		}
		token := strings.Split(authLine[0], " ")[1]
		if !services.TokenIsValid(token) {
			c.AbortWithStatus(403)
		}
		c.Set("userId", 1) // TODO (needs to get userID from token)
		c.Next()
	}

}
