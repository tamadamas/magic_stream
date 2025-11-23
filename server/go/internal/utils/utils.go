package utils

import (
	"context"
	"log/slog"
)

func LoggerFromCtx(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value("logger").(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
