package routes

import (
	//"net/http"

	"practice-sales-backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrdersRoutes(router *gin.RouterGroup) {

	orders := controllers.Orders{}

	router.GET("/:order_id", orders.GetOrder)
	// 修改訂單
	router.PATCH("/:order_id", orders.UpdateOrder)
	// 刪除訂單
	router.DELETE("/:order_id", orders.DeleteOrder)
	// 查詢訂單內產品
	// router.GET("/:order_id/products", orders.GetOrder)

}
