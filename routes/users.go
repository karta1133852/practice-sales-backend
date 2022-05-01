package routes

import (
	//"net/http"

	"practice-sales-backend/controllers"

	"github.com/gin-gonic/gin"
)

var usersController controllers.Users

func UsersRoutes(router *gin.RouterGroup) {

	router.Use(authModel.Authenticate)

	// 新增 User
	router.POST("/", usersController.CreateUser)

	router.GET("/:uid", usersController.GetUser)
	router.PATCH("/:uid", usersController.UpdateUser)
	// 取得使用者的訂單列表
	router.GET("/:uid/orders", usersController.GetUserOrders)
	// 新增訂單
	router.POST("/:uid/orders", usersController.NewUserOrders)

}
