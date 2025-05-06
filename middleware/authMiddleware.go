package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/MohakGupta2004/auth-go/utils/token"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Authorization")
		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "jwt expired",
			})
			return
		}
		userToken := strings.Split(accessToken, " ")[1]
		log.Print(userToken)
		claims, err := token.ValidateToken(userToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Set("username", claims.Username)
		c.Set("user_type", claims.User_type)

	}
}
