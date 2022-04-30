package controllers

import (
	"practice-sales-backend/models/db"

	"github.com/gin-gonic/gin"
)

// AuthController
type usersController struct{} // 方便閱讀 private
type Users struct {           // 包裝給外部使用
	*usersController
}

// 接收 c.Request.Body
// Field name 開頭要大寫 -> Public
type UserData struct {
	Username string
	Password string
}

type User struct {
	Uid              int
	Username         string
	Password         string
	Salt             string
	Coin             int
	Point            int
	Vip_Type         string
	Accumulate_spent int
}

func (_ *usersController) CreateUser(c *gin.Context) {

	var userData UserData
	c.BindJSON(&userData)

	// TODO 加鹽

	// query := c.Request.URL.Query()
	// c.String(200, userData.name)
	c.JSON(200, userData)
}

func (_ *usersController) GetUser(c *gin.Context) {

	uid := c.Param("uid")
	var user User
	db.GetDB().SelectOne(&user, "SELECT * FROM public.users WHERE uid=$1 LIMIT 1", uid)

	// if err != nil {
	// 	return user, token, err
	// }
	// var response map[string]interface{}
	// response

	c.JSON(200, user)
}

func (_ *usersController) UpdateUser(c *gin.Context) {
	c.String(200, "PATCH UpdateUser()")
}

func (_ *usersController) GetUserOrders(c *gin.Context) {
	c.String(200, "GET GetUserOrders()")
}
