package routes

import (
	//"net/http"

	"github.com/gin-gonic/gin"
)

func MainRoutes(router *gin.RouterGroup) {

	// AuthRoutes(router.Group("/auth"))
	// UserRoutes(router.Group("/users"))
	useRoutes("/auth", router)
	useRoutes("/users", router)
	useRoutes("/orders", router)

	router.GET("/someGET", func(c *gin.Context) { c.String(200, "someGET") })
	router.POST("/somePOST", func(c *gin.Context) { c.String(200, "somePOST") })

}

func useRoutes(path string, router *gin.RouterGroup) {

	switch path {
	case "/auth":
		AuthRoutes(router.Group(path))
	case "/users":
		UsersRoutes(router.Group(path))
	case "/orders":
		OrdersRoutes(router.Group(path))
	}
}
