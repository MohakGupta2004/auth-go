package routes

import (
	"github.com/MohakGupta2004/auth-go/controllers/auth/login"
	"github.com/MohakGupta2004/auth-go/controllers/auth/register"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/users/login", login.LoginController())
	router.POST("/users/signup", register.RegisterController())
}
