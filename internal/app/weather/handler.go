package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/minchao/slack-tw-weather/internal/pkg"
	"github.com/minchao/slack-tw-weather/internal/pkg/slack"

	"github.com/aws/aws-lambda-go/events"
	"github.com/go-playground/form"
	"github.com/spf13/cobra"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var command slack.SlashCommand
	if err := parseSlashCommand(request, &command); err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	args := strings.Split(command.Text, " ")
	args = prepareArgs(rootCmd, args)
	_, output, err := pkg.ExecuteCommandC(rootCmd, args...)
	if err != nil {
		fmt.Printf("Command execution error: %s", err)

		output = err.Error()
	}

	message := slack.Message{
		ResponseType: "in_channel",
		Text:         output,
	}
	body, _ := json.Marshal(&message)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func parseSlashCommand(request events.APIGatewayProxyRequest, command *slack.SlashCommand) error {
	req, _ := http.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	contentType, _ := request.Headers["Content-Type"]
	req.Header.Add("Content-Type", contentType)
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := form.NewDecoder().Decode(command, req.Form); err != nil {
		return err
	}
	return nil
}

func prepareArgs(rootCmd *cobra.Command, args []string) []string {
	if len(args) != 1 {
		return args
	}
	if args[0] == "" {
		return []string{}
	}
	if args[0][:1] == "-" {
		return args
	}
	for _, c := range rootCmd.Commands() {
		if c.Name() == args[0] {
			return args
		}
	}
	// Use forecast command with county, if the command not found.
	return []string{forecastCmd.Name(), args[0]}
}
