package server

import (
	"context"
	"fmt"
	"net/http"

	"go-template/pkg/config"
	"go-template/pkg/log"

	"github.com/gin-gonic/gin"
)

func Start(ctx context.Context, r *gin.Engine, cfg config.ServerConfig) {
	// disable trusted proxy warning
	if err := r.SetTrustedProxies(nil); err != nil {
		log.S.Fatal(
			"Failed to configure trusted proxies settings",
		)
	}

	// create new server
	srv := &http.Server{
		Handler: r,
		Addr: fmt.Sprintf(
			"%s:%d",
			cfg.Host,
			cfg.Port,
		),
	}

	log.S.Info(fmt.Sprintf(
		"Server is listening on %s:%d",
		cfg.Host,
		cfg.Port,
	))

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(context.Background()); err != nil {
			log.S.Fatal(
				"Failed to stop the server",
			)
		} else {
			log.S.Info(
				"Server is stopped",
			)
		}
	}()

	// server runs in a goroutine
	if err := srv.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.S.Fatal(
			"An error occurred, cannot listen for requests anymore",
		)
	}
}
