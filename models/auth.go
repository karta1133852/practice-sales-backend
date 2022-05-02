package models

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthModel
type authModel struct{} // 方便閱讀 private
type Auth struct {      // 包裝給外部使用
	*authModel
}

// custom claims (payload)
type Payload struct {
	Uid int
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (_ *authModel) HashAndSalt(strPwd string) (hashedPwd string, err error) {

	bytePwd := []byte(strPwd)
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		return
	}

	hashedPwd = string(hash)
	return
}

func (_ *authModel) ComparePasswords(hashedPwd string, strPwd string) bool {

	byteHash := []byte(hashedPwd)
	bytePwd := []byte(strPwd)

	// 密碼一樣回傳 nil, 不一樣回傳 err
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		return false
	}
	return true
}

// CreateToken: 透過 uid 生成 JWT Token
func (_ *authModel) CreateToken(uid int, username string) (token string, expiredTime time.Time, err error) {

	now := time.Now()
	expiredTime = now.Add(1 * time.Hour)
	jwtId := strconv.Itoa(uid) + strconv.FormatInt(now.Unix(), 10)

	// JWT Payload 設定
	payload := Payload{
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			Audience:  username,
			ExpiresAt: expiredTime.Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "practice-sales-backend",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   username,
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(jwtSecret)
	if err != nil {
		return
	}

	return token, expiredTime, nil
}

// TODO 要不要搬到 api/middleware 底下
// Authenticate: 驗證用 middleware
func (_ *authModel) Authenticate(c *gin.Context) (err error) {

	auth := c.GetHeader("Authorization")
	s := strings.Split(auth, "Bearer ")
	if len(s) <= 1 {
		err = &CustomError{StatusCode: 401, Message: "Token not found"}
		c.Abort()
		return
	}
	token := s[1]

	tokenClaims, err := jwt.ParseWithClaims(token, &Payload{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors != 0 {
				if jwt.ValidationErrorMalformed != 0 {
					message = "Token error"
				} else if jwt.ValidationErrorSignatureInvalid != 0 {
					message = "Signature wrong"
				} else if jwt.ValidationErrorExpired != 0 {
					message = "Token expired"
				} else if jwt.ValidationErrorNotValidYet != 0 {
					message = "Token not yet valid"
				} else {
					message = "Token error"
				}
			}
		}
		// http.StatusUnauthorized
		err = &CustomError{StatusCode: 401, Message: message}
		c.Abort()
		return
	}

	// 判斷 token 是否 valid
	if tokenClaims.Valid {
		c.Next()
		return
	} else {
		err = &CustomError{StatusCode: 401, Message: "Token invalid"}
		c.Abort() // 取消執行接下來的 middleware
		return
	}
}
