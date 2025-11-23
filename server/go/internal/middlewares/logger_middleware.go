package middlewares

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(baseLogger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID, _ := c.Get("request_id")
		reqLogger := baseLogger.With("request_id", reqID)

		ctx := context.WithValue(c.Request.Context(), "logger", reqLogger)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
