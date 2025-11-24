package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/tamadamas/magic_stream/server/go/internal/models"
	"github.com/tamadamas/magic_stream/server/go/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsersRepository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewUsersRepository(db *mongo.Database) *UsersRepository {
	return &UsersRepository{
		db:  db,
		col: db.Collection("users"),
	}
}

func (r *UsersRepository) Create(ctx context.Context, user *models.User) error {

	if err := validator.New().Struct(user); err != nil {
		return fmt.Errorf("Failed to register: %w", err)
	}

	checkUser, err := r.col.CountDocuments(ctx, bson.M{"email": user.Email})

	if err != nil || checkUser > 0 {
		return fmt.Errorf("Failed to register: User already exists")
	}

	passHash, err := utils.HashPassword(user.Password)

	if err != nil {
		return fmt.Errorf("Failed to register: Invalid Password")
	}

	user.UserID = bson.NewObjectID().Hex()
	user.Password = passHash
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = r.col.InsertOne(ctx, user)

	if err != nil {
		return fmt.Errorf("Failed to register: %w", err)

	}

	return nil
}
