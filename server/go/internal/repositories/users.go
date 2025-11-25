package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/tamadamas/magic_stream/server/go/internal/app_errors"
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

func (r *UsersRepository) Login(ctx context.Context, userLogin *models.UserLogin) (*models.User, error) {
	var user models.User

	err := r.col.FindOne(ctx, bson.M{"email": userLogin.Email}).Decode(&user)

	if err != nil {
		return nil, app_errors.NewNotFoundError(nil, "User is not found")
	}

	correntPass := utils.VerifyPassword(userLogin.Password, user.Password)

	if !correntPass {
		return nil, errors.New("Invalid email or passwpord")
	}

	token, refreshToken, err := utils.GenerateTokens(&user)

	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	updateAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	updateData := bson.M{
		"$set": bson.M{
			"token":         token,
			"refresh_token": refreshToken,
			"update_at":     updateAt,
		},
	}

	_, err = r.col.UpdateOne(ctx, bson.M{"user_id": user.UserID}, updateData)

	return &user, nil

}
