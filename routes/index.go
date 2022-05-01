package routes

import (
	"github.com/gin-gonic/gin"

	"practice-sales-backend/models"
)

var authModel models.Auth

func MainRoutes(router *gin.RouterGroup) {

	// TODO auth check with /users, /orders
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
