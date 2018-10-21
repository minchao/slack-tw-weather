package app

import (
	"os"

	"github.com/spf13/cobra"
)

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
	args := prepareArgs(rootCmd, os.Args[1:])
	rootCmd.SetArgs(args)
	rootCmd.Execute()
}
