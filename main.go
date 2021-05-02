package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/golangApi/controllers"
	"github.com/golangApi/middleware"
	"github.com/golangApi/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	// Log all the shit in a log file
	setLogOutput()

	// Set basic auth middleware
	server.Use(middleware.BasicAuth())

	// Load the static content
	server.Static("/css", "./templates/css")

	// Load the html templates
	server.LoadHTMLGlob("templates/*.html")

	// Set up basic debugging tool But this is just in case of debugging
	// server.Use(gindump.Dump())

	apiRoutes := server.Group("/api/v1")
	{
		apiRoutes.GET("/videos", func(c *gin.Context) {
			c.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(c *gin.Context) {
			err := videoController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Hey welcome again"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
