package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 參考 https://zhuanlan.zhihu.com/p/76967528
type HandlerFunc func(c *gin.Context) error

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(c)
	}
}

func ErrorHandler(c *gin.Context) {
	// 確保其他 middleware 執行
	c.Next()

	for _, err := range c.Errors {
		// log, handle, etc.
		// log.Printf("Error: %s", err.Error())
		c.String(http.StatusNotFound, "Error: "+err.Error())
		// c.Status(http.StatusNotFound)
		// c.Writer.Write([]byte(err.Error()))
	}
}
