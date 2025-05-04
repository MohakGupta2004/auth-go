package main

import (
	"os"

	"github.com/MohakGupta2004/auth-go/routes"
	"github.com/MohakGupta2004/auth-go/utils/env"
	"github.com/gin-gonic/gin"
)

func main() {
	port := env.GetString("PORT", ":8080")

	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	router.GET("/api/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "True"})
	})

	router.Run(port)
}
