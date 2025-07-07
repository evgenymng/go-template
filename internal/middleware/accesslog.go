package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Middleware to log every incoming and processed request.
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		zap.S().Infow(
			"Request",
			"path",
			c.Request.URL.Path,
			"params",
			c.Request.URL.Query(),
			"host",
			c.Request.URL.Hostname(),
		)

		c.Next()

		status := c.Writer.Status()
		zap.S().Infow(
			"Response",
			"status_code",
			status,
			"elapsed",
			time.Since(startTime).Seconds(),
		)
	}
}
