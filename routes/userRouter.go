package routes

import (
	"github.com/MohakGupta2004/auth-go/controllers/users"
	"github.com/MohakGupta2004/auth-go/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.GET("/api/v1/users/all", users.GetAllUsers())
	router.GET("/api/v1/users/:id", users.GetOneUser())
}
