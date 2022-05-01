package routes

import (
	//"net/http"

	"practice-sales-backend/controllers"

	"github.com/gin-gonic/gin"
)

var authController controllers.Auth

func AuthRoutes(router *gin.RouterGroup) {

	//router.GET("/checkAuth", c.String(200, "checkAuth") })
	router.POST("/login", authController.Login)
	router.PUT("/logout", authController.Logout)

}
