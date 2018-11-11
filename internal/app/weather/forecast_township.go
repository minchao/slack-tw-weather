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
		[]string{"Wx,PoP6h,T,CI,RH"},
	)
	if err != nil {
		return err
	}
	if len(forecast.Records.Locations[0].Location) == 0 {
		return errors.New("township not found")
	}

	var messages []string
	var j = 0
	for i := 0; i < 24; i++ {
		if msg := getTownshipForecastDescription(forecast.Records.Locations[0].Location[0].WeatherElement, i); msg != "" {
			messages = append(messages, msg)
			j++
		}
		if j == 6 {
			break
		}
	}
	cmd.Println(strings.Join(messages, "\n"))

	return nil
}

func getTownshipForecastDescription(elements []cwb.FTWWeatherElement, position int) string {
	var date, d, wx, pop6h, t, ci, rh string

	date = *elements[0].Time[position].StartTime
	switch date[11:16] {
	case "00:00", "03:00", "09:00", "15:00", "21:00":
		return ""
	case "06:00":
		d = "上午"
	case "12:00":
		d = "下午"
	case "18:00":
		d = "晚上"
	}

	for _, element := range elements {
		switch element.ElementName {
		case "Wx":
			wx = element.Time[position].ElementValue[0].Value
		case "T":
			t = element.Time[position].ElementValue[0].Value
		case "CI":
			ci = element.Time[position].ElementValue[1].Value
		case "RH":
			rh = element.Time[position].ElementValue[0].Value
		case "PoP6h":
			for _, pop := range element.Time {
				if *pop.StartTime == date {
					pop6h = pop.ElementValue[0].Value
				}
			}
		}
	}
	return fmt.Sprintf("%s/%s %s，天氣%s，降雨機率 %s%%，溫度 %s 度，%s，相對濕度 %s%%",
		date[5:7], date[8:10], d, wx, pop6h, t, ci, rh)
}
