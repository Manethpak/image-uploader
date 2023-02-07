package main

import (
	"image-uploader/route"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	addr := ":8080"

	route.SetupRoutes(server)

	server.Static("/public", "./public/images")

	server.Run(addr)
}
