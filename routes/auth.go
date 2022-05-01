package routes

import (
	"github.com/gin-gonic/gin"

	"practice-sales-backend/controllers"
	"practice-sales-backend/models"
)

var authController controllers.Auth

func AuthRoutes(router *gin.RouterGroup) {

	var authModel models.Auth

	//router.GET("/checkAuth", c.String(200, "checkAuth") })
	router.POST("/login", authController.Login)
	router.PUT("/logout", authModel.Authenticate, authController.Logout)
}
