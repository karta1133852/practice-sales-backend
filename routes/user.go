package routes

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {

	router.GET("/someGET", func(c *gin.Context) { c.String(200, "someGET") })
	router.POST("/somePOST", func(c *gin.Context) { c.String(200, "somePOST") })

}
