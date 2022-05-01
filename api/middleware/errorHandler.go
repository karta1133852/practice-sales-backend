package middleware

import (
	"net/http"
	"practice-sales-backend/models"

	"github.com/gin-gonic/gin"
)

// 參考 https://zhuanlan.zhihu.com/p/76967528
type HandlerFunc func(c *gin.Context) error

// Wrapper(error handler): 將 middleware 包起來使其可以回傳 error
func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 執行包起來的 middleware
		err := handler(c)

		if err != nil {
			customErr, ok := err.(*models.CustomError)
			if ok { // 自訂 rrror
				c.JSON(customErr.StatusCode, customErr.ToJSON())
			} else { // 普通 error
				c.String(http.StatusInternalServerError, "Error: "+err.Error())
			}
		}

		// TODO github.com/pkg/errors warp error stack
	}
}

func ErrorHandler(c *gin.Context) {
	// 其他 middleware 先執行
	c.Next()

	// 再來處理 Error
	for _, err := range c.Errors {
		// log, handle, etc.
		// log.Printf("Error: %s", err.Error())
		c.String(http.StatusNotFound, "Error: "+err.Error())
		// c.Status(http.StatusNotFound)
		// c.Writer.Write([]byte(err.Error()))
	}
}
