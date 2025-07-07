package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

// Sample OpenTelemetry logs endpoint
//
//	@summary	Send logs to the OpenTelemetry collector with trace correlation
//	@tags		misc
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{object}	map[string]any
//
//	@router		/send-logs [get]
func SendLogs(c *gin.Context) {
	ctx, span := tracer.Start(c, "SendLogs")
	defer span.End()

	// add span attributes for context
	span.SetAttributes(
		attribute.String("http.method", c.Request.Method),
		attribute.String("http.route", "/send-logs"),
		attribute.String("operation.type", "logging_demo"),
	)

	// extract trace information for response
	spanContext := span.SpanContext()
	traceID := spanContext.TraceID().String()
	spanID := spanContext.SpanID().String()

	// Regular logging (no trace correlation)
	zap.S().Info("Regular log message without trace correlation")

	// Regular logging (with trace correlation)
	otelzap.Ctx(ctx).Debug("Debug level log entry")

	otelzap.Ctx(ctx).Info("Info level log entry")

	// Structured logging
	otelzap.Ctx(ctx).Sugar().Warnw(
		"Warning level log entry",
		"warning_type", "rate_limit_approaching",
		"current_usage", 85,
		"limit", 100,
	)

	// Simulate an error scenario
	simulatedError := errors.New("simulated error for demonstration")
	otelzap.Ctx(ctx).Sugar().Errorw(
		"Error level log entry",
		"error", simulatedError,
	)

	c.JSON(200, gin.H{
		"message": "Logs sent successfully with trace correlation!",
		"trace_info": gin.H{
			"trace_id":  traceID,
			"span_id":   spanID,
			"operation": "SendLogs",
		},
	})
}
