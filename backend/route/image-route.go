package route

import (
	"image-uploader/controller"
	"image-uploader/service"

	"github.com/gin-gonic/gin"
)

func addImageRoutes(rg *gin.RouterGroup) {
	imageService := service.NewImageService()
	imageController := controller.NewImageController(imageService)

	rg.POST("/image", func(c *gin.Context) {
		imageController.UploadImage(c)
	})
}
