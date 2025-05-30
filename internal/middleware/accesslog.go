package middleware

import (
	"time"

	"go-template/pkg/log"

	"github.com/gin-gonic/gin"
)

// Middleware to log every incoming and processed request.
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		log.S.Infow(
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
		log.S.Infow(
			"Response",
			"status_code",
			status,
			"elapsed",
			time.Since(startTime).Seconds(),
		)
	}
}
