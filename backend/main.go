package main

import (
	"image-uploader/route"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if _, err := os.Stat("public"); os.IsNotExist(err) {
		os.Mkdir("public", 0755)
	}
}

func main() {
	server := gin.Default()

	route.SetupRoutes(server)

	server.Static("/public", "./public")

	server.Run()
}
