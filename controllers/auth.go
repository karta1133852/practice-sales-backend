package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"practice-sales-backend/models"
	"practice-sales-backend/models/db"
)

// AuthController
type authController struct{} // 方便閱讀 private
type Auth struct {           // 包裝給外部使用
	*authController
}

var authModel models.Auth

func (_ *authController) HashAndSalt(strPwd string) (err error, hashedPwd string) {

	bytePwd := []byte(strPwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashedPwd = string(hash)
	return
}

func (_ *authController) ComparePasswords(hashedPwd string, strPwd string) bool {

	byteHash := []byte(hashedPwd)
	bytePwd := []byte(strPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

func (this *authController) Login(c *gin.Context) {

	body := struct {
		Username string
		Password string
	}{}
	c.ShouldBindJSON(&body)

	data := struct {
		Uid       int
		HashedPwd string `db:"password"`
	}{}
	err := db.GetDB().SelectOne(&data,
		`SELECT uid, password FROM users WHERE username=$1;`,
		body.Username,
	)
	if err != nil {
		c.Error(err)
		return
	}

	isPwdSame := this.ComparePasswords(data.HashedPwd, body.Password)
	if !isPwdSame {
		c.Error(errors.New("使用者名稱或密碼錯誤"))
		return
	}

	token, expiredTime, err := authModel.CreateToken(data.Uid, body.Username)
	if err != nil {
		c.Error(err)
		return
	}

	// TODO add to redis
	expiredTime.UTC()

	c.JSON(200, gin.H{
		"uid":   data.Uid,
		"token": token,
	})
}

func (_ *authController) Logout(c *gin.Context) {
	c.String(200, "PUT Logout()")
}
