package execution

import (
	"context"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Launch(f func(context.Context, *sync.WaitGroup)) {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()
	defer func() {
		// make sure all pending messages are flushed before exiting
		_ = zap.S().Sync()
	}()

	wg := sync.WaitGroup{}

	go f(ctx, &wg)

	<-ctx.Done() // wait for the termination signal
	zap.S().Info(
		"Shutting down gracefully, Ctrl+C to force.",
	)

	done := make(
		chan bool,
	)
	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer shutdownCancel()

	forceCtx, forceCancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer forceCancel()

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		zap.S().Info(
			"All work is done, shutting down",
		)
	case <-forceCtx.Done():
		zap.S().Info(
			"Server is shutdown forcefully",
		)
	case <-shutdownCtx.Done():
		zap.S().Info(
			"Shutdown timeout exceeded, forceful shutdown",
		)
	}
}
