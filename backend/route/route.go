package route

import (
	"image-uploader/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api.Use(middleware.CORSMiddleware())

	addImageRoutes(api)

}
