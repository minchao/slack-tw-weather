package app

import "github.com/spf13/cobra"

var (
	radarCmd = &cobra.Command{
		Use:   "radar",
		Short: "Weather radar (Composite reflectivity)",
		Run:   radarFunc,
	}
)

func radarFunc(cmd *cobra.Command, args []string) {
}
