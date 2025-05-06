package main

import (
	"fmt"
	"go-url-shortener/api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	setupRouter(router)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}

func setupRouter(router *gin.Engine) {
	router.POST("/api/v1/", routes.ShortenURL)
	router.GET("api/v1/", routes.GetByShortID)
}
