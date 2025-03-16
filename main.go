package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nish1013/ai-claims-assessment/internal/api"
)

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)

	router.Run(":4003")
}
