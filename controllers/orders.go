package controllers

import "github.com/gin-gonic/gin"

// AuthController
type ordersController struct{} // 方便閱讀 private
type Orders struct {           // 包裝給外部使用
	*ordersController
}

func (_ *ordersController) GetOrder(c *gin.Context) (err error) {
	c.String(200, "GET GetOrder()")
	return
}

func (_ *ordersController) UpdateOrder(c *gin.Context) (err error) {
	c.String(200, "PATCH UpdateOrder()")
	return
}

func (_ *ordersController) DeleteOrder(c *gin.Context) (err error) {
	c.String(200, "DELETE DeleteOrder()")
	return
}
