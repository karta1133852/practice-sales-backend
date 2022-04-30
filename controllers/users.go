package controllers

import (
	"errors"
	"math"
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
	Uid               int
	Username          string
	Password          string
	Salt              string
	Coin              int
	Point             int
	Vip_Type          string
	Accumulated_spent int
}

func (_ *usersController) CreateUser(c *gin.Context) {

	userData := struct {
		Username string
		Password string
	}{}
	c.BindJSON(&userData)

	// TODO 加鹽
	// TODO 新增至 Database

	// query := c.Request.URL.Query()
	// c.String(200, userData.name)
	c.JSON(200, userData)
}

func (_ *usersController) GetUser(c *gin.Context) {

	uid := c.Param("uid")
	user := struct {
		Uid               int
		Username          string
		Coin              int
		Point             int
		Vip_Type          string
		Accumulated_spent int
	}{}

	queryStr := `
		SELECT uid, username, coin, point, vip_type, accumulated_spent
		FROM public.users WHERE uid=$1 LIMIT 1`

	err := db.GetDB().SelectOne(&user, queryStr, uid)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, user)
}

func (_ *usersController) UpdateUser(c *gin.Context) {
	c.String(200, "PATCH UpdateUser()")
}

func (_ *usersController) GetUserOrders(c *gin.Context) {
	c.String(200, "GET GetUserOrders()")
}

type product_item struct {
	Product_no uint
	Quantity   uint
}

func (_ *usersController) NewUserOrders(c *gin.Context) {

	// Parse order data
	var orderData = struct {
		Total       uint
		Payed_coin  uint
		Payed_point uint
		Exchange    uint
		Discount    uint
		Products    []product_item
	}{}
	c.BindJSON(&orderData)

	c.JSON(200, orderData)
	return

	// 檢查參數
	if orderData.Total == 0 {
		c.Error(errors.New("訂單金額需大於 0 元"))
	}

	discountTotal := uint(math.Round(float64(orderData.Total) * float64(orderData.Discount) / 100.0))
	equivalentTotal := orderData.Payed_coin + uint(math.Round(float64(orderData.Payed_point)*(float64(orderData.Exchange)/100.0)))
	if discountTotal != equivalentTotal {
		c.Error(errors.New("付款金額錯誤"))
	}

	// 取得使用者與優惠折扣資料
	uid := c.Param("uid")
	querySelect := `
		TODO
	`

	user := struct {
		Coin              int
		Point             int
		Vip_Type          string
		Accumulated_spent int
	}{}
	errSelect := db.GetDB().SelectOne(&user, querySelect, uid)
	if errSelect != nil {
		c.Error(errSelect)
	}

	// 檢查金額是否足夠
	//if user.Coin

	queryInsert := `
		WITH _order_id AS (
			INSERT INTO orders (cost_coin, cost_point)
			VALUES ($1, $2) RETURNING order_id;
		)

		INSERT INTO order_item (order_id, product_no, quantity)
		VALUES (_order_id, unnest(ARRAY[$3]), unnest(ARRAY[$4]));
	`

	rows, errInsert := db.GetDB().Query(queryInsert)
	if errInsert != nil {
		c.Error(errInsert)
	}

	rows.Close()

	c.String(200, "POST NewUserOrders()")
}
