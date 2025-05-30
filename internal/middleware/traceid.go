package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Middleware to associate an ID with each incoming request.
func TraceIdMiddleware(
	traceIdHeader string,
	traceIdKey string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get(traceIdHeader)
		if len(traceId) == 0 {
			newTraceId, _ := uuid.NewV7()
			traceId = newTraceId.String()
		}
		c.Set(traceIdKey, traceId)

		c.Next()
	}
}
