package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const name = "github.com/evgenymng/go-template/internal/middleware"

var (
	meter        = otel.Meter(name)
	httpDuration metric.Float64Histogram
)

func init() {
	var err error
	httpDuration, err = meter.Float64Histogram(
		"http.server.duration_seconds",
		metric.WithDescription("Measures HTTP request duration in seconds"),
	)
	if err != nil {
		panic(err)
	}
}

func OtelAccessLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		httpDuration.Record(c, duration, metric.WithAttributes(
			attribute.String("method", c.Request.Method),
			attribute.String("path", c.FullPath()),
			attribute.Int("status", c.Writer.Status())))
	}
}
