package controllers

import (
	"github.com/gin-gonic/gin"

	"practice-sales-backend/models"
	"practice-sales-backend/models/db"
)

// AuthController
type authController struct{} // 方便閱讀 private
type Auth struct {           // 包裝給外部使用
	*authController
}

func (this *authController) Login(c *gin.Context) (err error) {

	body := struct {
		Username string
		Password string
	}{}
	c.ShouldBindJSON(&body)

	data := struct {
		Uid       int
		HashedPwd string `db:"password"`
	}{}
	err = db.GetDB().SelectOne(&data,
		`SELECT uid, password FROM users WHERE username=$1;`,
		body.Username,
	)
	if err != nil {
		return
	}

	var authModel models.Auth
	isPwdSame := authModel.ComparePasswords(data.HashedPwd, body.Password)
	if !isPwdSame {
		err = &models.CustomError{StatusCode: 401, Message: "使用者名稱或密碼錯誤"}
		return
	}

	token, expiredTime, err := authModel.CreateToken(data.Uid, body.Username)
	if err != nil {
		return
	}

	// TODO add to redis
	expiredTime.UTC()

	c.JSON(200, gin.H{
		"uid":   data.Uid,
		"token": token,
	})
	return
}

func (_ *authController) Logout(c *gin.Context) (err error) {
	// TODO 從 Redis 中刪除 Token
	c.String(200, "PUT Logout()")
	return
}
