package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nish1013/ai-claims-assessment/internal/api"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func main() {
	initLogger()

	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		logrus.Warn("Warning: No .env file found. Using default values.")
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
		logrus.Fatalf("Failed to set trusted proxies: %v", errSetTrustProxy)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4003"
	}

	logrus.WithFields(logrus.Fields{
		"port": port,
	}).Info("Starting server...")

	router.Run(":" + port)
}

func initLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
