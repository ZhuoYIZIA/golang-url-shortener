package main

import (
	_ "github.com/joho/godotenv/autoload"
	// _ "github.com/lib/pq"
	"github.com/gin-gonic/gin"

	// url-shortener
	"github.com/ZhuoYIZIA/url-shortener/internal/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		controllers.Convert()

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}