package main

import (
	//"net/http"

	"practice-sales-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	// Basic settings
	server.SetTrustedProxies(nil)

	// Use RouterGroup to do nested routes
	routes.MainRoutes(server.Group("/api"))

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"title":   "Test root",
			"message": "Hi!",
		})
	})

	err := server.Run(":3000")
	if err != nil {
		panic(err)
	}
}
