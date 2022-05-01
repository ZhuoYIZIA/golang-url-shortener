package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ZhuoYIZIA/url-shortener/controllers"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}

	router.Use(cors.New(config))

	v1Router := router.Group("v1")
	{
		v1Router.POST("/", controllers.CreateUrl)
		// v1Router.GET("/:id", controllers.GetUrl)
	}

	return router
}
