package middleware

import (
	"time"

	"app/internal/app/ckey"
	"app/internal/log"

	"github.com/gin-gonic/gin"
)

// Middleware to log every incoming and processed request.
func AccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetString(string(ckey.TraceId))
		startTime := time.Now()
		l := log.L().TraceId(traceId).
			Add("path", c.Request.URL.Path).
			Add("params", c.Request.URL.Query()).
			Add("host", c.Request.URL.Hostname())
		log.S.Info("Request received", l)

		c.Next()

		status := c.Writer.Status()
		l = l.
			Add("status_code", status).
			Add("elapsed", time.Since(startTime).Seconds())
		log.S.Info("Response sent", l)
	}
}
