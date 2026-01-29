package middlewares

import (
	"time"

	"site-admin-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs all HTTP requests
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)
		latencyMs := float64(latency.Nanoseconds()) / 1e6

		// Get status code
		statusCode := c.Writer.Status()

		// Get request details
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		// Log request
		logger.LogRequest(method, path, clientIP, statusCode, latencyMs)

		// Log errors if any
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				logger.Error(err.Error())
			}
		}
	}
}
