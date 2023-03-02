package controller

import (
	"image-uploader/service"
	"log"
	"os"

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
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)

	fileoutput := os.Getenv("URL")
	if os.Getenv("GIN_MODE") == "release" {
		fileoutput += "/public/" + file.Filename
	} else {
		fileoutput += ":" + os.Getenv("PORT") + "/public/" + file.Filename
	}

	err = ctx.SaveUploadedFile(file, "public/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"message": "success",
		"output":  fileoutput,
	})
}
