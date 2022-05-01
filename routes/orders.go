package routes

import (
	"github.com/gin-gonic/gin"

	. "practice-sales-backend/api/middleware"
	"practice-sales-backend/controllers"
)

func OrdersRoutes(router *gin.RouterGroup) {

	var ordersController controllers.Orders

	router.Use(Wrapper(authModel.Authenticate))

	router.GET("/:order_id", Wrapper(ordersController.GetOrder))
	// 修改訂單
	router.PATCH("/:order_id", Wrapper(ordersController.UpdateOrder))
	// 刪除訂單
	router.DELETE("/:order_id", Wrapper(ordersController.DeleteOrder))
	// 查詢訂單內產品
	// router.GET("/:order_id/products", Wrapper(ordersController.GetOrder))

}
