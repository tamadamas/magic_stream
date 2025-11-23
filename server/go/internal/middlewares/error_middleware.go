package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tamadamas/magic_stream/server/go/internal/app_errors"
	"github.com/tamadamas/magic_stream/server/go/internal/utils"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		logger := utils.LoggerFromCtx(c.Request.Context())

		for _, ginErr := range c.Errors {
			logger.Error("request error", "error", ginErr.Err)
		}

		for _, ginErr := range c.Errors {
			err := ginErr.Err

			if httpErr, ok := err.(app_errors.HTTPError); ok {
				// return typed error first
				c.JSON(httpErr.Status(), gin.H{"error": httpErr.Error()})
				return
			}
		}

		c.JSON(500, gin.H{
			"error": "internal server error",
		})
	}
}
