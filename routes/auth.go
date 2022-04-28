package routes

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup) {

	router.GET("/checkAuth", func(c *gin.Context) { c.String(200, "checkAuth") })
	router.POST("/login", func(c *gin.Context) { c.String(200, "POST login") })

}
