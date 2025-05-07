package helloworld

import (
	"bmbot/util/logging"
	"context"
	"fmt"
)

type Service interface {
	SayHello(ctx context.Context) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (h service) SayHello(ctx context.Context) error {
	logger, ok := logging.FromContext(ctx)
	if !ok {
		return fmt.Errorf("logger not found in context")
	}

	logger.InfoContext(ctx, "Hello World!")

	return nil
}
