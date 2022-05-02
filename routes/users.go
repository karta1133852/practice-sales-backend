package routes

import (
	"github.com/gin-gonic/gin"

	. "practice-sales-backend/api/middleware"
	"practice-sales-backend/controllers"
)

func UsersRoutes(router *gin.RouterGroup) {

	var usersController controllers.Users

	// 新增 User
	router.POST("/", Wrapper(usersController.CreateUser))
	// 取得使用者資料
	router.GET("/:uid", Wrapper(authModel.Authenticate), Wrapper(usersController.GetUser))
	// 修改使用者資料
	router.PATCH("/:uid", Wrapper(authModel.Authenticate), Wrapper(usersController.UpdateUser))
	// 取得使用者的訂單列表
	router.GET("/:uid/orders", Wrapper(authModel.Authenticate), Wrapper(usersController.GetUserOrders))
	// 新增訂單
	router.POST("/:uid/orders", Wrapper(authModel.Authenticate), Wrapper(usersController.NewUserOrders))

}
