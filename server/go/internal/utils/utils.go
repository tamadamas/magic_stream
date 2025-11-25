package utils

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
	"github.com/tamadamas/magic_stream/server/go/internal/models"
)

func LoggerFromCtx(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value("logger").(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}

func HashPassword(pass string) (string, error) {
	argon := argon2.DefaultConfig()
	hash, err := argon.HashEncoded([]byte(pass))

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func VerifyPassword(pass, hash string) bool {
	ok, err := argon2.VerifyEncoded([]byte(pass), []byte(hash))

	if err != nil || !ok {
		return false
	}

	return true
}

func GenerateTokens(user *models.User) (string, string, error) {
	var secretKey string = os.Getenv("JWT_SECRET")
	var refreshKey string = os.Getenv("JWT_REFRESH")

	if secretKey == "" || refreshKey == "" {
		panic("JWT_SECRET and JWT_REFRESH env variables required")
	}

	claims := &models.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Genres:    user.Genres,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MagicStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", "", err
	}

	refreshClaims := &models.UserResponse{
		UserID:    user.UserID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Genres:    user.Genres,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MagicStream",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 7 * time.Hour)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString([]byte(refreshKey))

	if err != nil {
		return "", "", err
	}

	return signedToken, signedRefreshToken, nil

}
