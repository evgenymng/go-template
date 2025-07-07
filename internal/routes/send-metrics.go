package routes

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const metricsName = "github.com/evgenymng/go-template/internal/routes/metrics"

var meter = otel.Meter(metricsName)

var (
	requestCounter    metric.Int64Counter
	responseHistogram metric.Float64Histogram
	activeConnections metric.Int64UpDownCounter
)

func init() {
	var err error

	requestCounter, err = meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("1"),
	)
	if err != nil {
		panic(err)
	}

	responseHistogram, err = meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("Duration of HTTP requests in seconds"),
		metric.WithUnit("s"),
	)
	if err != nil {
		panic(err)
	}

	activeConnections, err = meter.Int64UpDownCounter(
		"active_connections",
		metric.WithDescription("Number of active connections"),
		metric.WithUnit("1"),
	)
	if err != nil {
		panic(err)
	}
}

// Sample OpenTelemetry metrics endpoint
//
//	@summary	Send metrics to the OpenTelemetry collector
//	@tags		misc
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{object}	map[string]any
//
//	@router		/send-metrics [get]
func SendMetrics(c *gin.Context) {
	_, span := tracer.Start(c, "SendMetrics")
	defer span.End()

	// Add span attributes
	span.SetAttributes(
		attribute.String("http.method", c.Request.Method),
		attribute.String("http.route", "/send-metrics"),
		attribute.String("operation.type", "metrics_demo"),
	)

	// increment request counter
	requestCounter.Add(c.Request.Context(), 1, metric.WithAttributes(
		attribute.String("method", c.Request.Method),
		attribute.String("endpoint", "/send-metrics"),
		attribute.String("status", "200"),
	))

	// record response time (simulate random duration)
	responseTime := rand.Float64() * 2.0 // Random duration between 0-2 seconds
	responseHistogram.Record(
		c.Request.Context(),
		responseTime,
		metric.WithAttributes(
			attribute.String("method", c.Request.Method),
			attribute.String("endpoint", "/send-metrics"),
		),
	)

	// simulate active connections change
	connectionChange := rand.Intn(10) - 5 // Random change between -5 to +4
	activeConnections.Add(
		c.Request.Context(),
		int64(connectionChange),
		metric.WithAttributes(
			attribute.String("service", "go-template"),
		),
	)

	c.JSON(200, gin.H{
		"message": "Metrics sent successfully!",
		"metrics": gin.H{
			"request_count":     1,
			"response_time":     responseTime,
			"connection_change": connectionChange,
			"timestamp":         time.Now().Unix(),
		},
	})
}
