package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// Middleware to log every incoming and processed request.
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// use context-aware logging to correlate with traces
		otelzap.Ctx(c).Sugar().Infow(
			"Request",
			"path", c.Request.URL.Path,
			"params", c.Request.URL.Query(),
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
		)

		c.Next()

		status := c.Writer.Status()
		otelzap.Ctx(c).Sugar().Infow(
			"Response",
			"status_code", status,
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"elapsed", time.Since(startTime).Seconds(),
		)
	}
}
