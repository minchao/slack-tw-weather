package weather

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/minchao/go-cwb/cwb"
	"github.com/spf13/cobra"
)

var (
	forecastTownshipCmd = &cobra.Command{
		Use:   "forecast:township [township]",
		Short: "3 day weather township forecasts",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires township argument")
			}
			return nil
		},
		RunE: forecastTownshipFunc,
	}
)

func init() {
	forecastTownshipCmd.Flags().StringP("county", "", "", "County")
}

func forecastTownshipFunc(cmd *cobra.Command, args []string) error {
	county, _ := cmd.Flags().GetString("county")
	township := args[0]

	location, err := cwb.FindLocationByName(county)
	if err != nil {
		return err
	}

	client := cwb.NewClient(os.Getenv("CWB_API_KEY"), nil)
	forecast, _, err := client.Forecasts.GetTownshipsWeatherByLocations(context.Background(),
		[]string{location.DataSet[0].Name},
		[]string{township},
		[]string{"WeatherDescription"},
	)
	if err != nil {
		return err
	}
	if len(forecast.Records.Locations[0].Location) == 0 {
		return errors.New("township not found")
	}

	var messages []string
	for i := 0; i < 24; i++ {
		messages = append(messages,
			getTownshipForecastDescription(forecast.Records.Locations[0].Location[0], i))
	}
	cmd.Println(strings.Join(messages, "\n"))

	return nil
}

func getTownshipForecastDescription(location cwb.FTWDatasetLocation, position int) string {
	var date, description string

	for _, element := range location.WeatherElement {
		if element.ElementName == "WeatherDescription" {
			date = *element.Time[position].StartTime
			description = element.Time[position].ElementValue[0].Value
		}
	}
	return fmt.Sprintf("%s，天氣%s", date[5:16], description)
}
