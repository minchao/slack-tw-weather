package main

import (
	"github.com/minchao/slack-tw-weather/internal/app"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(app.Handler)
}
