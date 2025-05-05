package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID            primitive.ObjectID `bson:"_id"`
	Username      *string            `json:"username" validate:"required,min=2,max=100"`
	Email         *string            `json:"email" validate:"required"`
	Password      *string            `json:"password" validate:"required,min=6"`
	Access_token  *string            `json:"access_token"`
	Refresh_token *string            `json:"refresh_token"`
	Phone         *int64             `json:"phone_number" validate:"required"`
	User_type     *string            `json:"usertype" validate:"required,eq=ADMIN|eq=USER"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"userId"`
}
