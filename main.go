package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nish1013/ai-claims-assessment/internal/api"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found. Using default values.")
	}

	router := gin.Default()
	api.RegisterRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4003"
	}

	log.Printf("Starting server on port %s...", port)

	router.Run(":" + port)
}
