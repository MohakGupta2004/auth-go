package users

import (
	"context"
	"net/http"
	"time"

	"github.com/MohakGupta2004/auth-go/database"
	"github.com/MohakGupta2004/auth-go/models"
	"github.com/MohakGupta2004/auth-go/utils/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func GetAllUsers() {

}

func GetOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")

		if err := auth.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.UserModel
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"success": user,
		})
	}
}
