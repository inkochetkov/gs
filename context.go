package gs

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NewContext GracefulShutdown only
func NewContext(ctx context.Context) (context.Context, context.CancelFunc) {

	ctxCancel, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)

	return ctxCancel, stop
}

// NewModuleContext to module
func NewModuleContext(ctx context.Context) context.Context {

	ctxCancel, _ := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)

	return ctxCancel
}
