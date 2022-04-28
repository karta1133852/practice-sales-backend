package routes

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.RouterGroup) {

	//AuthRoutes(router.Group("/auth"))
	//UserRoutes(router.Group("/user"))
	useRoutes("/auth", router)
	useRoutes("/user", router)

	router.GET("/someGET", func(c *gin.Context) { c.String(200, "someGET") })
	router.POST("/somePOST", func(c *gin.Context) { c.String(200, "somePOST") })

}

func useRoutes(path string, router *gin.RouterGroup) {

	switch path {
	case "/auth":
		AuthRoutes(router.Group("/auth"))
	case "/user":
		UserRoutes(router.Group("/user"))
	}
}
