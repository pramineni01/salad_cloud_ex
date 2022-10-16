package utils

import (
	"context"
	"os"
	"signal"
)

func GetContext(ctx context.Context) (context.Context, context.CancelFunc) {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		defer cancel()

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

		select {
		case sigCh:
		case ctx.Done():
		}
	}()

	return ctx, cancel
}