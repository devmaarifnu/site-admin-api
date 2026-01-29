package middlewares

import (
	"fmt"
	"runtime/debug"

	"site-admin-api/pkg/logger"
	"site-admin-api/pkg/response"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware recovers from panics and logs the error
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				logger.Errorf("PANIC: %v\nStack trace:\n%s", err, string(debug.Stack()))

				// Return error response
				response.InternalServerError(c, fmt.Sprintf("Internal server error: %v", err))
				c.Abort()
			}
		}()
		c.Next()
	}
}
