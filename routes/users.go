package routes

import (
	"github.com/gin-gonic/gin"

	. "practice-sales-backend/api/middleware"
	"practice-sales-backend/controllers"
)

func UsersRoutes(router *gin.RouterGroup) {

	var usersController controllers.Users

	router.Use(Wrapper(authModel.Authenticate))

	// 新增 User
	router.POST("/", Wrapper(usersController.CreateUser))
	// 取得使用者資料
	router.GET("/:uid", Wrapper(usersController.GetUser))
	// 修改使用者資料
	router.PATCH("/:uid", Wrapper(usersController.UpdateUser))
	// 取得使用者的訂單列表
	router.GET("/:uid/orders", Wrapper(usersController.GetUserOrders))
	// 新增訂單
	router.POST("/:uid/orders", Wrapper(usersController.NewUserOrders))

}
