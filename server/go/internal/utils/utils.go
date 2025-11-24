package utils

import (
	"context"
	"log/slog"

	"github.com/matthewhartstonge/argon2"
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
