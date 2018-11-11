package weather

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/minchao/go-epa"
	"github.com/spf13/cobra"
)

var (
	local *time.Location

	forecastPm25Cmd = &cobra.Command{
		Use:   "forecast:pm25",
		Short: "Air quality forecasts",
		RunE:  forecastPm25Func,
	}
)

func init() {
	local, _ = time.LoadLocation("Asia/Taipei")
}

func forecastPm25Func(cmd *cobra.Command, _ []string) error {
	client := epa.NewClient("", nil)

	options := url.Values{}
	options.Set("sort", "PublishTime")
	resp, _, err := client.GetAirQualityForecast(context.Background(), options)
	if err != nil {
		return err
	}

	cmd.Println(getForecastPm25Description(resp.Result.Records))

	return nil
}

func getForecastPm25Description(forecasts []epa.AirQualityForecast) string {
	var messages []string
	for _, record := range forecasts {
		switch record.Area {
		case "北部", "中部", "高屏":
			if record.ForecastDate != time.Now().In(local).Add(24*time.Hour).Format("2006-01-02") {
				continue
			}

			content := record.Content
			content = strings.Replace(content, "\r", "\n", -1)
			content = regexp.MustCompile("\\d\\.").ReplaceAllString(content, "• ")

			msg := fmt.Sprintf("%s/%s %s地區\n%s",
				record.ForecastDate[5:7],
				record.ForecastDate[8:],
				record.Area,
				content)

			messages = append(messages, msg)
		}
	}
	return strings.Join(messages, "\n")
}
