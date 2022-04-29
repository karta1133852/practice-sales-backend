package controllers

import "github.com/gin-gonic/gin"

// AuthController
type usersController struct{} // 方便閱讀 private
type Users struct {           // 包裝給外部使用
	*usersController
}

func (_ *usersController) CreateUser(c *gin.Context) {
	c.String(200, "POST CreateUser()")
}

func (_ *usersController) GetUser(c *gin.Context) {
	c.String(200, "GET GetUser()")
}

func (_ *usersController) UpdateUser(c *gin.Context) {
	c.String(200, "PATCH UpdateUser()")
}

func (_ *usersController) GetUserOrders(c *gin.Context) {
	c.String(200, "GET GetUserOrders()")
}
