package routes

import (
	"go-template/pkg/db/fakedb"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
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
//	@success	200	{string}	string
//
//	@router		/send-trace [get]
func SendTrace(c *gin.Context) {
	ctx, span := tracer.Start(c, "SendTrace")
	defer span.End()

	number := fakedb.FetchFromDb(ctx)

	c.String(200, "You rolled: %d!", number)
}
