package routes

import (
	//"net/http"

	"practice-sales-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {

	auth := new(controllers.Auth)

	//router.GET("/checkAuth", c.String(200, "checkAuth") })
	router.POST("/login", auth.Login)
	router.PUT("/logout", auth.Logout)

}
