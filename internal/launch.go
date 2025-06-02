package internal

import (
	"context"
	"net/http"
	"sync"

	"go-template/docs"
	"go-template/internal/middleware"
	"go-template/internal/routes"
	"go-template/internal/server"
	"go-template/pkg/config"
	"go-template/pkg/execution"
	"go-template/pkg/log"
	"go-template/pkg/telemetry"

	"go.uber.org/zap"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggofiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

const (
	traceIdKey    = "trace_id"
	traceIdHeader = "X-Trace-ID"
)

func Launch() {
	execution.Launch(func(ctx context.Context, wg *sync.WaitGroup) {
		if config.C.Dev {
			gin.SetMode("debug")
		} else {
			gin.SetMode("release")
		}

		r := gin.New()

		devGroup := r.Group("")
		{
			devGroup.GET("/ping", routes.GetPing)
			devGroup.GET("/version", routes.GetVersion)

			if config.C.Dev {
				// register pprof endpoints
				pprof.Register(devGroup)

				// register swagger docs
				docs.SwaggerInfo.Version = config.C.Version
				devGroup.GET(
					"/swagger/*any",
					swagger.WrapHandler(swaggofiles.Handler),
				)
				devGroup.GET("/docs", func(c *gin.Context) {
					c.Redirect(
						http.StatusMovedPermanently,
						"/swagger/index.html",
					)
				})
			}
		}

		mws := []gin.HandlerFunc{
			middleware.ResponseHandler(traceIdKey),
			middleware.TraceIdMiddleware(traceIdHeader, traceIdKey),
			middleware.AccessLogMiddleware(),
			middleware.OtelAccessLogMiddleware(),
		}

		// unused main group
		v1 := r.Group("", mws...)
		{
			v1.GET("/send-trace", routes.SendTrace)
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			shutdown, err := onStartup(ctx)
			if err != nil {
				log.S.Fatalw("Failed to start the application", zap.Error(err))
			}
			server.Start(ctx, r, config.C.Server)
			err = shutdown(ctx)
			if err != nil {
				log.S.Fatalw(
					"Failed to shutdown the application",
					zap.Error(err),
				)
			}
		}()
		log.S.Info("Application started")
	})
}

// Add all required startup logic here.
func onStartup(
	ctx context.Context,
) (shutdown func(context.Context) error, err error) {
	shutdown, err = telemetry.Init(ctx)
	if err != nil {
		return
	}
	log.S.Info("OpenTelemetry SDK initialized")

	return shutdown, nil
}
