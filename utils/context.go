package utils

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func GetContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer cancel()

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

		select {
		case <-sigs:
		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}