package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nish1013/ai-claims-assessment/internal/api"
	"log"
	"os"
	"strings"
)

func main() {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		log.Println("Warning: No .env file found. Using default values.")
	}

	router := gin.Default()
	api.RegisterRoutes(router)

	proxyEnv := os.Getenv("TRUSTED_PROXIES")
	var trustedProxies []string
	if proxyEnv != "" {
		trustedProxies = strings.Split(proxyEnv, ",")
	}

	errSetTrustProxy := router.SetTrustedProxies(trustedProxies)
	if errSetTrustProxy != nil {
		log.Fatalf("Failed to set trusted proxies: %v", errSetTrustProxy)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4003"
	}

	log.Printf("Starting server on port %s...", port)

	router.Run(":" + port)
}
