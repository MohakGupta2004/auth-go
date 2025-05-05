package login

import (
	"context"
	"net/http"
	"time"

	"github.com/MohakGupta2004/auth-go/database"
	"github.com/MohakGupta2004/auth-go/models"
	"github.com/MohakGupta2004/auth-go/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func VerifyPassword(userPassword string, foundUserPassword string) (ok bool, err error) {
	ok = true
	err = nil

	err = bcrypt.CompareHashAndPassword([]byte(foundUserPassword), []byte(userPassword))

	if err != nil {
		return false, err
	}

	return ok, err
}

type LoginUser struct {
	Email    *string
	Password *string
}

func LoginController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)

		var user LoginUser
		var result models.UserModel

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		if err := validate.Struct(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&result)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Email doesn't exist",
			})
			return
		}
		validPassoword, err := VerifyPassword(*user.Password, *result.Password)
		if validPassoword != true {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "password is not valid",
			})
		}
		//email string, username string, user_type string, uid string
		accessToken, refreshToken, _ := token.GenerateAllTokens(*result.Email, *result.Username, *result.User_type, *&result.User_id)

		result.Access_token = &accessToken
		result.Refresh_token = &refreshToken
		c.JSON(http.StatusOK, gin.H{"user": result})

	}
}
