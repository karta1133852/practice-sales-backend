package routes

import (
	"practice-sales-backend/controllers"

	"github.com/gin-gonic/gin"
)

var ordersController controllers.Orders

func OrdersRoutes(router *gin.RouterGroup) {

	router.Use(authModel.Authenticate)

	router.GET("/:order_id", ordersController.GetOrder)
	// 修改訂單
	router.PATCH("/:order_id", ordersController.UpdateOrder)
	// 刪除訂單
	router.DELETE("/:order_id", ordersController.DeleteOrder)
	// 查詢訂單內產品
	// router.GET("/:order_id/products", ordersController.GetOrder)

}
