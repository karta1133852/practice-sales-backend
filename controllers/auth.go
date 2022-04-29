package controllers

import "github.com/gin-gonic/gin"

// AuthController
type authController struct{} // 方便閱讀 private
type Auth struct {           // 包裝給外部使用
	*authController
}

func (_ *authController) Login(c *gin.Context) {
	c.String(200, "POST Login()")
}

func (_ *authController) Logout(c *gin.Context) {
	c.String(200, "PUT Logout()")
}
