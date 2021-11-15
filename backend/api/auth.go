package api

import (
	"fmt"
	"logistics/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Auth(group *gin.RouterGroup, service services.AuthService) {
	handler := group.Group("auth")

	handler.POST("login", login(service))
	handler.POST("register", register(service))
}

func login(service services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *services.LoginRequest
		if err := c.BindJSON(&request); err != nil {
			logrus.Info("error binding login data ", err)
			return
		}
		user, token, err := service.Login(request)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "login success",
			"user":    user,
			"token":   token,
		})
	}
}

func register(service services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request *services.RegisterRequest
		if err := c.BindJSON(&request); err != nil {
			logrus.Info("error binding register data ", err)
			c.JSON(400, gin.H{
				"error": err,
			})
			return
		}
		fmt.Println(request.Email, request.Password, request.Birthday, request.Sex)
		user, err := service.Register(request)
		fmt.Println("have got user from register service ", user)
		fmt.Println("have got err from register service ", err)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "user created",
			"user":    user,
		})
	}
}
