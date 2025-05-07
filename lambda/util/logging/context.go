package logging

import (
	"context"
	"log/slog"
)

type key int

var loggerKey key

func NewContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func FromContext(ctx context.Context) (*slog.Logger, bool) {
	u, ok := ctx.Value(loggerKey).(*slog.Logger)
	return u, ok
}
