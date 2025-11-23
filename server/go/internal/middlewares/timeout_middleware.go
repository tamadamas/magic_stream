package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		finished := make(chan struct{})
		panicChan := make(chan any)

		// Run handlers in a goroutine
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			c.Next()
			finished <- struct{}{}
		}()

		// Wait for one of the cases
		select {
		case <-ctx.Done():
			// Timeout happened
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "request timeout",
			})
			c.Abort()
		case p := <-panicChan:
			// Pass panic to default behavior
			panic(p)
		case <-finished:
			// Completed normally
		}
	}
}
