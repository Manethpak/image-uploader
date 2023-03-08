package controller

import (
	"image-uploader/service"
	"log"

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
		log.Fatal(err)
		ctx.JSON(400, gin.H{
			"message": "File not found",
			"output":  nil,
		})
	}

	path, err := c.service.SaveImage(file, header)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(400, gin.H{
			"message": "Error while uploading file",
			"output":  nil,
		})
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"output":  path,
	})
}
