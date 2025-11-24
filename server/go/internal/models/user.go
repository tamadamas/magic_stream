package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID       string        `bson:"user_id" json:"user_id"`
	FirstName    string        `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName     string        `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email        string        `bson:"email" json:"email" validate:"required,email"`
	Password     string        `bson:"password" json:"password" validate:"required,min=10,max=100"`
	Token        string        `bson:"token" json:"token"`
	RefreshToken string        `bson:"refresh_token" json:"refresh_token"`
	Role         string        `bson:"role" json:"role" validate:"required,oneof=admin user"`
	Genres       []Genre       `bson:"genre" json:"genre" validate:"required,dive"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
