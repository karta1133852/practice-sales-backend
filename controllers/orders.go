package controllers

import "github.com/gin-gonic/gin"

// AuthController
type ordersController struct{} // 方便閱讀 private
type Orders struct {           // 包裝給外部使用
	*ordersController
}

func (_ *ordersController) GetOrder(c *gin.Context) {
	c.String(200, "GET GetOrder()")
}

func (_ *ordersController) UpdateOrder(c *gin.Context) {
	c.String(200, "PATCH UpdateOrder()")
}

func (_ *ordersController) DeleteOrder(c *gin.Context) {
	c.String(200, "DELETE DeleteOrder()")
}
