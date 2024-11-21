package middleware

import (
	"app/internal/app/ckey"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Middleware to associate an ID with each incoming request.
func TraceIdMiddleware(header string) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get(header)
		if len(traceId) == 0 {
			newTraceId, _ := uuid.NewV7()
			traceId = newTraceId.String()
		}
		c.Set(string(ckey.TraceId), traceId)

		c.Next()
	}
}
