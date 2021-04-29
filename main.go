package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/api/v1/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello from GoLang",
		})
	})

	server.Run(":8080")
}
