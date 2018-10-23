package main

import (
	"github.com/minchao/slack-tw-weather/internal/app/funcs/radar"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(radar.Handler)
}
