package route

import (
	"image-uploader/controller"
	"image-uploader/service"
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

	imageService := service.NewImageService()
	imageController := controller.NewImageController(imageService)

	api.POST("/image", func(c *gin.Context) {
		imageController.UploadImage(c)
	})

}
