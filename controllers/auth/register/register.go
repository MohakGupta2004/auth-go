package register

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/MohakGupta2004/auth-go/database"
	"github.com/MohakGupta2004/auth-go/models"
	"github.com/MohakGupta2004/auth-go/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(password string) *string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal("Can't able to create a hashed password")
	}
	hashedPassword := string(pass[:])
	return &hashedPassword
}

func RegisterController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), time.Second*100)
		var user models.UserModel

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		validatorErr := validate.Struct(user)
		if validatorErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validatorErr})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking the email"})
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "email already exists",
			})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		accessToken, refreshToken, _ := token.GenerateAllTokens(*user.Email, *user.Username, *user.User_type, *&user.User_id)
		user.Access_token = &accessToken
		user.Refresh_token = &refreshToken
		user.Password = HashPassword(*user.Password)
		result, err := userCollection.InsertOne(ctx, user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		defer cancel()
		c.JSON(http.StatusOK, gin.H{"success": result})

	}
}
