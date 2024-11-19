package main

import (
	"image-uploader/route"
	"log"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		log.Print("App running with default environment variables.")
		gin.SetMode(gin.DebugMode)
	}

	if _, err := os.Stat("public"); os.IsNotExist(err) {
		os.Mkdir("public", 0755)
	}
}

func main() {
	server := gin.Default()

	server.Use(static.Serve("/public", static.LocalFile("./public", true)))
	server.Use(static.Serve("/", static.LocalFile("../frontend/dist", true)))

	route.SetupRoutes(server)

	server.Run()
}
