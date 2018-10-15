package app

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use: "weather",
	}
)

func init() {
	rootCmd.Version = "0.0.1"
	rootCmd.AddCommand(forecastCmd)
	rootCmd.AddCommand(radarCmd)
}

func Execute() {
	rootCmd.Execute()
}
