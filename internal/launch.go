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

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

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
		}

		// unused main group
		_ = r.Group("", mws...)

		wg.Add(1)
		go func() {
			defer wg.Done()
			err := onStartup(ctx)
			if err != nil {
				log.S.Fatalw("Failed to startup application", zap.Error(err))
			}
			server.Start(ctx, r, config.C.Server)
			onShutdown(ctx)
		}()
		log.S.Info("Application started")
	})
}

// Add all required onStartup logic here.
func onStartup(_ context.Context) error {
	return nil
}

// Add all required onShutdown logic here.
func onShutdown(_ context.Context) {
}
