package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"practice-sales-backend/models"
	"practice-sales-backend/models/db"
)

// AuthController
type usersController struct{} // 方便閱讀 private
type Users struct {           // 包裝給外部使用
	*usersController
}

var usersModel = models.Users{}

func (_ *usersController) CreateUser(c *gin.Context) (err error) {

	body := struct {
		Username string
		Password string
	}{}
	c.ShouldBindJSON(&body)

	// TODO 加鹽
	// TODO 新增至 Database

	// query := c.Request.URL.Query()
	// c.String(200, userData.name)
	c.JSON(200, body)
	return
}

func (_ *usersController) GetUser(c *gin.Context) (err error) {

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
		FROM public.users WHERE uid=$1 LIMIT 1
	`
	err = db.GetDB().SelectOne(&user, queryStr, uid)
	if err != nil {
		return
	}

	c.JSON(200, user)
	return
}

func (_ *usersController) UpdateUser(c *gin.Context) (err error) {
	c.String(200, "PATCH UpdateUser()")
	return
}

func (_ *usersController) GetUserOrders(c *gin.Context) (err error) {
	c.String(200, "GET GetUserOrders()")
	return
}

func (_ *usersController) NewUserOrders(c *gin.Context) (err error) {

	// Parse order data
	var body = models.OrderData{}
	c.ShouldBindJSON(&body)

	// 檢查參數
	if body.OriginalPrice == 0 {
		c.Error(errors.New("訂單金額需大於 0 元"))
		return
	}

	// 檢查付款金額是否相符
	if err = usersModel.CheckTotal(body); err != nil {
		return
	}

	// 取得使用者與優惠折扣資料
	uid := c.Param("uid")
	strTimeNow := time.Now().Format(time.RFC3339)
	querySelect := `
		WITH u AS (
			SELECT *
			FROM users
			WHERE uid=$1
		)
		SELECT * FROM
			(SELECT coin, point, accumulated_spent FROM u) a,
			(SELECT pi.percentage_off, pi.exchange
			FROM (
			SELECT *
				FROM promotions AS p
				WHERE
				($2 BETWEEN p.start_time AND p.end_time) 
				OR
				COALESCE(p.start_time, p.end_time) IS NULL
				ORDER BY p.p_no DESC LIMIT 1
			) AS p
			INNER JOIN (
				SELECT p_no, pt[1] percentage_off, pt[2] exchange
				FROM (
				SELECT p_no, ARRAY_AGG(value ORDER BY promotion_type) pt
				FROM promotion_items
				WHERE vip_type=(SELECT vip_type FROM u)
				GROUP BY 1) z
			) AS pi
			USING (p_no)) b;
	`

	data := struct {
		Coin             int
		Point            int
		AccumulatedSpent int `db:"accumulated_spent"`
		PercentageOff    int `db:"percentage_off"`
		Exchange         int
	}{}
	err = db.GetDB().SelectOne(&data, querySelect, uid, strTimeNow)
	if err != nil {
		return
	}

	// 檢查金額是否足夠
	//if user.Coin
	// TODO 讀取 SQL error 判斷

	// 檢查優惠資料是否相符
	if data.PercentageOff != body.Discount || data.Exchange != body.Exchange {
		err = errors.New("優惠資料錯誤！")
		return
	}

	strNo, strQuantity := usersModel.FormatProductItems(body.Products)

	// Start transaction
	txn, err := db.GetDB().Begin()
	if err != nil {
		return
	}

	// Rollback the transaction
	defer txn.Rollback()

	var orderID int // 新產生的訂單編號
	row := txn.QueryRow(
		`INSERT INTO orders (cost_coin, cost_point, time) VALUES ($1, $2, $3) RETURNING order_id;`,
		body.PayedCoin, body.PayedPoint, strTimeNow,
	)
	if err = row.Scan(&orderID); err != nil {
		return
	}

	// ARRAY[$]內使用 Query format 會被強制加上 ' '
	sqlInsertItem := fmt.Sprintf(`
		INSERT INTO order_items (order_id, product_no, quantity)
		VALUES (%d, unnest(ARRAY[%s]), unnest(ARRAY[%s]));`,
		orderID, strNo, strQuantity,
	)
	_, err = txn.Exec(sqlInsertItem)
	if err != nil {
		return
	}

	rowUpdated := txn.QueryRow(
		`UPDATE users
		SET coin=coin-$1, point=point-$2, accumulated_spent=accumulated_spent+$3
		WHERE uid=$4 RETURNING uid, coin, point, accumulated_spent;`,
		body.PayedCoin, body.PayedPoint, body.PayedCoin, uid,
	)
	if err = rowUpdated.Err(); err != nil {
		//db.PrintDbError(err)
		return
	}

	res := struct {
		Uid              int
		Coin             int
		Point            int
		AccumelatedSpent int `db:"accumulated_spent"`
		orderId          int // 僅用於 Response
	}{}
	err = rowUpdated.Scan(&res.Uid, &res.Coin, &res.Point, &res.AccumelatedSpent)
	if err != nil {
		return
	}

	// Commit transaction
	err = txn.Commit()
	if err != nil {
		return
	}

	// 加上 orderId
	res.orderId = orderID
	c.JSON(200, res)
	return
}
