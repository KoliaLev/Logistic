package api

import (
	"logistics/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TransporUnits(group *gin.RouterGroup) {
	handler := group.Group("transport/units")

	handler.POST("buy", middleware.AuthMiddleware(), createUnit())
}

func createUnit() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("call - createUnits")
		id, _ := c.Get("userId")
		logrus.Info("userId ", id)
		c.JSON(200, gin.H{
			"message": "unit succesfuly created",
		})
	}
}
