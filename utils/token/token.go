package token

import (
	"fmt"
	"log"
	"time"

	"github.com/MohakGupta2004/auth-go/database"
	"github.com/MohakGupta2004/auth-go/utils/env"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SignedDetails struct {
	Email     string
	Username  string
	User_type string
	Uid       string
	jwt.RegisteredClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var SECRET_KEY string = env.GetString("SECRET_KEY", "secretkey")

func ValidateToken(accessToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(accessToken, &SignedDetails{}, func(t *jwt.Token) (any, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("Token is invalid")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		msg = fmt.Sprintf("Token is expired")
		return
	}

	return claims, msg
}

func GenerateAllTokens(email string, username string, user_type string, uid string) (accessToken, refreshToken string, error error) {
	claims := &SignedDetails{
		Email:     email,
		Username:  username,
		User_type: user_type,
		Uid:       uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(164 * time.Hour)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return accessToken, refreshToken, err
}
