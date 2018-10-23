package main

import (
	"github.com/minchao/slack-tw-weather/internal/app/weather"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(weather.Handler)
}
