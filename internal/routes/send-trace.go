package routes

import (
	"go-template/pkg/db/fakedb"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

const name = "github.com/evgenymng/go-template/internal/routes"

var tracer = otel.Tracer(name)

// Sample OpenTelemetry tracing endpoint
//
//	@summary	Send trace to the OpenTelemetry collector, as well as processing time.
//	@tags		misc
//	@accept		plain
//
//	@produce	json
//
//	@success	200	{object}	map[string]any
//
//	@router		/send-trace [get]
func SendTrace(c *gin.Context) {
	ctx, span := tracer.Start(c, "SendTrace")
	defer span.End()

	// add span attributes for better observability
	span.SetAttributes(
		attribute.String("http.method", c.Request.Method),
		attribute.String("http.route", "/send-trace"),
	)

	// simulate database operation
	number := fakedb.FetchFromDb(ctx)

	// add the result as a span attribute
	span.SetAttributes(
		attribute.Int("db.result", number),
	)

	// extract trace information
	spanContext := span.SpanContext()
	traceID := spanContext.TraceID().String()
	spanID := spanContext.SpanID().String()

	c.JSON(200, gin.H{
		"message": "Trace sent successfully!",
		"trace_info": gin.H{
			"trace_id": traceID,
			"span_id":  spanID,
		},
		"data": gin.H{
			"dice_roll": number,
		},
	})
}
