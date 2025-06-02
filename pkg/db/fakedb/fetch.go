package fakedb

import (
	"math/rand/v2"
	"time"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

const name = "github.com/evgenymng/go-template/pkg/db/fakedb"

var tracer = otel.Tracer(name)

func FetchFromDb(ctx context.Context) int {
	ctx, span := tracer.Start(context.Background(), "FetchFromDb")
	defer span.End()
	time.Sleep(1 * time.Second)
	return rand.IntN(10)
}
