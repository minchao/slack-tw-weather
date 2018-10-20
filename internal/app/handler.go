package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/minchao/slack-tw-weather/internal/pkg"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/form"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	req, _ := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	contentType, _ := request.Headers["Content-Type"]
	req.Header.Add("Content-Type", contentType)
	if err := req.ParseForm(); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	var command pkg.SlashCommand
	if err := form.NewDecoder().Decode(&command, req.Form); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	args := strings.Split(command.Text, " ")
	_, output, err := pkg.ExecuteCommandC(rootCmd, args...)
	if err != nil {
		fmt.Printf("Command execution error: %s", err)

		output = err.Error()
	}

	message := pkg.Message{
		ResponseType: "in_channel",
		Text:         output,
	}
	body, _ := json.Marshal(&message)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}
