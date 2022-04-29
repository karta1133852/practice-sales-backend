package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"

	"practice-sales-backend/api/middleware"
	"practice-sales-backend/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//fmt.Println("test: ", os.Getenv("TEST"))

	server := gin.Default()
	// 基本設定
	server.SetTrustedProxies(nil)
	server.Use(gzip.Gzip(gzip.DefaultCompression))

	// CORS
	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, PATCH, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          60 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	// 暫時回傳 err stack
	server.Use(middleware.ErrorHandler)

	// Use RouterGroup to do nested routes
	routes.MainRoutes(server.Group("/api"))

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"title":   "Root",
			"message": "Hi!",
		})
	})

	port := "3000"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	server.Run(":" + port) // 單次簡單 concat 所以不使用 strings.Builder
}
