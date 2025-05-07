package logging

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"log/slog"
	"os"
)

var baseLogger = slog.New(
	slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

func NewLogger(ctx context.Context) (*slog.Logger, error) {
	lambdaCtx, ok := lambdacontext.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("lambda context not found")
	}

	logger := baseLogger.With("requestId", lambdaCtx.AwsRequestID)
	return logger, nil
}
