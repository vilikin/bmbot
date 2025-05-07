package main

import (
	"bmbot/application/helloworld"
	"bmbot/util/logging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	service = helloworld.NewService()
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, event json.RawMessage) {
	logger, err := logging.NewLogger(ctx)
	if err != nil {
		panic(err)
	}

	ctx = logging.NewContext(ctx, logger)

	err = service.SayHello(ctx)
	if err != nil {
		logger.ErrorContext(ctx, fmt.Sprintf("error saying hallo: %v", err))
	}
}
