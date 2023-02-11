package main

import (
	"image-uploader/route"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	server := gin.Default()

	route.SetupRoutes(server)

	server.Static("/public", "./public")

	server.Run()
}
