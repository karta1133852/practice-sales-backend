package models

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthModel
type authModel struct{} // 方便閱讀 private
type Auth struct {      // 包裝給外部使用
	*authModel
}

// custom claims
type Payload struct {
	Uid int
	jwt.StandardClaims
}

// 透過 uid 生成 JWT Token
func (_ *authModel) CreateToken(uid int, username string) (token string, err error) {

	now := time.Now()
	jwtId := strconv.Itoa(uid) + strconv.FormatInt(now.Unix(), 10)

	// JWT Payload 設定
	payload := Payload{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: now.Add(1 * time.Hour).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "practice-sales-backend",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   username,
		},
	}
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(jwtSecret)
	if err != nil {
		return
	}

	return token, nil
}
