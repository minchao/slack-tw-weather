package app

import "github.com/spf13/cobra"

var (
	forecastCmd = &cobra.Command{
		Use:     "forecast",
		Aliases: county,
		Run:     forecastFunc,
	}

	county = []string{
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
		"Taichung",
	}
)

func forecastFunc(cmd *cobra.Command, args []string) {
}
