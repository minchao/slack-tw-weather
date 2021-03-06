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
	forecastCmd = &cobra.Command{
		Use:   "forecast [county]",
		Short: "36 hour weather forecasts",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires county argument")
			}
			return nil
		},
		RunE: forecastFunc,
	}

	counties = []string{
		"宜蘭縣",
		"花蓮縣",
		"臺東縣",
		"澎湖縣",
		"金門縣",
		"連江縣",
		"臺北市",
		"新北市",
		"桃園市",
		"臺中市",
		"臺南市",
		"高雄市",
		"基隆市",
		"新竹縣",
		"新竹市",
		"苗栗縣",
		"彰化縣",
		"南投縣",
		"雲林縣",
		"嘉義縣",
		"嘉義市",
		"屏東縣",
	}
)

func forecastFunc(cmd *cobra.Command, args []string) error {
	county := mapCounty(args[0])
	if !inCounties(county) {
		return fmt.Errorf("invalid county specified: %s", county)
	}

	client := cwb.NewClient(os.Getenv("CWB_API_KEY"), nil)
	forecast, _, err := client.Forecasts.Get36HourWeather(context.Background(), []string{county}, nil)
	if err != nil {
		return err
	}

	var messages []string
	for i := 0; i < 3; i++ {
		messages = append(messages, getForecastDescription(forecast.Records.Location[0], i))
	}
	cmd.Println(strings.Join(messages, "\n"))

	return nil
}

func mapCounty(name string) (county string) {
	county = strings.Replace(name, "台", "臺", -1)
	county = strings.ToLower(county)
	switch county {
	case "臺北", "taipei":
		county = "臺北市"
	case "新北", "臺北縣", "new taipei":
		county = "新北市"
	case "桃園":
		county = "桃園市"
	case "臺中", "臺中縣", "taichung":
		county = "臺中市"
	case "彰化", "changhua":
		county = "彰化縣"
	case "臺南", "臺南縣", "tainan":
		county = "臺南市"
	case "高雄", "高雄縣", "kaohsiung":
		county = "高雄市"
	case "基隆":
		county = "基隆市"
	case "台東", "taitung":
		county = "臺東縣"
	}
	return county
}

func inCounties(county string) bool {
	for _, c := range counties {
		if c == county {
			return true
		}
	}
	return false
}

func getForecastDescription(location cwb.F36HWCountryLocation, position int) string {
	var date, wx, pop, minT, maxT string

	for _, element := range location.WeatherElement {
		switch element.ElementName {
		case "Wx":
			st := element.Time[position].StartTime
			switch element.Time[position].StartTime[11:] {
			case "00:00:00":
				date = " 凌晨到中午"
			case "06:00:00":
				date = " 白天"
			case "12:00:00":
				date = " 中午到凌晨"
			case "18:00:00":
				date = " 晚上"
			}
			date = fmt.Sprintf("%s/%s%s", st[5:7], st[8:10], date)
			wx = element.Time[position].Parameter.ParameterName
		case "PoP":
			pop = element.Time[position].Parameter.ParameterName
		case "MinT":
			minT = element.Time[position].Parameter.ParameterName
		case "MaxT":
			maxT = element.Time[position].Parameter.ParameterName
		}
	}
	return fmt.Sprintf("%s，天氣%s，降雨機率 %s%%，溫度 %s 至 %s 度", date, wx, pop, minT, maxT)
}
