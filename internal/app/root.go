package app

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "weather",
	}
)

func init() {
	rootCmd.AddCommand(forecastCmd)
	rootCmd.AddCommand(radarCmd)
}
