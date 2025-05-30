package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"

	"go-template/docs"
	"go-template/internal/middleware"
	"go-template/internal/routes"
	"go-template/pkg/config"
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
	setMode()

	// server will run using this context
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	// new gin server engine
	r := gin.New()
	r.Use(
		middleware.ResponseHandler(traceIdKey),
		middleware.TraceIdMiddleware(traceIdHeader, traceIdKey),
		middleware.AccessLogMiddleware(),
	)

	// register handlers
	r.GET("/ping", routes.GetPing)

	if config.C.Debug {
		// register pprof endpoints
		pprof.Register(r)

		// register swagger docs
		docs.SwaggerInfo.Version = config.C.Release
		r.GET("/swagger/*any", swagger.WrapHandler(swaggofiles.Handler))
		r.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
		log.S.Debug("Added /docs endpoint")
	}

	// disable trusted proxy warning
	if err := r.SetTrustedProxies(nil); err != nil {
		log.S.Fatal(
			"Failed to configure trusted proxies settings",
			zap.Error(err),
		)
	}

	// create new server
	srv := &http.Server{
		Handler: r,
	}
	// setting onShutdown logic
	srv.RegisterOnShutdown(onShutdown)

	// create listener
	listener, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%d",
		config.C.Server.Host,
		config.C.Server.Port,
	))
	defer func() {
		_ = listener.Close()
	}()

	if err != nil {
		log.S.Fatal("Failed to create listener", zap.Error(err))
	}

	// perform startup logic
	err = onStartup(ctx)

	if err == nil {
		// server runs in a goroutine
		go func() {
			if err := srv.Serve(listener); err != nil &&
				err != http.ErrServerClosed {
				log.S.Fatal(
					"An error occurred, cannot listen for requests anymore",
					zap.Error(err),
				)
			}
		}()

		// listen for the interrupt signal
		<-ctx.Done()

		// restore default behavior of the interrupt signal and notify user
		cancel()
		log.S.Info("Shutting down gracefully, press Ctrl+C to force")
		ctx, cancel = context.WithTimeout(
			context.Background(),
			config.C.Server.ShutdownTimeout,
		)
		defer cancel()
	}

	// perform shutdown logic
	if err := srv.Shutdown(ctx); err != nil {
		log.S.Error("Server forced to shutdown")
	}
}

func setMode() {
	var mode string
	if config.C.Debug {
		mode = "debug"
	} else {
		mode = "release"
	}
	gin.SetMode(mode)
}

// Add all required onShutdown logic here.
func onStartup(_ context.Context) error {
	return nil
}

// Add all required onShutdown logic here.
func onShutdown() {
}
