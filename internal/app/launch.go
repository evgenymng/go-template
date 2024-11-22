package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"app/docs"
	"app/internal/app/middleware"
	"app/internal/app/routes"
	"app/internal/config"
	"app/internal/log"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
)

func Launch() {
	gin.SetMode(config.C.Server.Mode)

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
		middleware.ResponseHandler(),
		middleware.TraceIdMiddleware("X-Trace-ID"),
		middleware.AccessLogMiddleware(),
		middleware.ApiAuthMiddleware(
			config.C.ApiKeys,
			"Authorization",
			[]string{
				"^/ping$",
				"^/swagger/.*$",
				"^/docs$",
				"^/debug/pprof/.*$",
			}),
	)

	// register handlers
	r.GET("/ping", routes.GetPing)

	if config.C.EnablePprof {
		pprof.Register(r)
	}
	if config.C.EnableDocs {
		docs.SwaggerInfo.Version = config.C.Version
		r.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))
		r.GET("/docs", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		})
		log.S.Debug("Added /docs endpoint", log.L())
	}

	// disable trusted proxy warning
	if err := r.SetTrustedProxies(nil); err != nil {
		log.S.Fatal(
			"Failed to configure trusted proxies settings",
			log.L().Error(err),
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
		log.S.Fatal("Failed to create listener", log.L().Error(err))
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
					log.L().Error(err),
				)
			}
		}()

		// listen for the interrupt signal
		<-ctx.Done()

		// restore default behavior on the interrupt signal and notify user of shutdown
		cancel()
		log.S.Info(
			"Shutting down gracefully, press Ctrl+C to force",
			log.L(),
		)
		ctx, cancel = context.WithTimeout(
			context.Background(),
			time.Duration(config.C.Server.ShutdownTimeout)*time.Second,
		)
		defer cancel()
	}

	// perform shutdown logic
	if err := srv.Shutdown(ctx); err != nil {
		log.S.Error(
			"Server forced to shutdown",
			log.L(),
		)
	}
}
