package routes

import (
	//"net/http"

	"practice-sales-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.RouterGroup) {

	users := controllers.Users{}

	// 新增 User
	router.POST("/", users.CreateUser)

	router.GET("/:uid", users.GetUser)
	router.PATCH("/:uid", users.UpdateUser)
	// 取得使用者的訂單列表
	router.GET("/:uid/orders", users.GetUserOrders)

}
