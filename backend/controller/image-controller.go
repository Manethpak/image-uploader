package controller

import (
	"net/http"

	"image-uploader/service"

	"github.com/gin-gonic/gin"
)

type ImageController interface {
	UploadImage(ctx *gin.Context)
}

type controller struct {
	service service.ImageService
}

func NewImageController(service service.ImageService) ImageController {
	return &controller{
		service,
	}
}

func (c *controller) UploadImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "File not found",
		})
	}

	outputPath, error := c.service.SaveImage(file, header)

	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to upload image to the server",
		})
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Image uploaded successfully",
		"path":    outputPath,
	})
}
