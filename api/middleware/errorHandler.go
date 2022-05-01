package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 參考 https://zhuanlan.zhihu.com/p/76967528
type HandlerFunc func(c *gin.Context) error

// Wrapper(error handler): 將 middleware 包起來並使其可以回傳 error
func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 執行包起來的 middleware
		handler(c)

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
