package weather

import (
	"github.com/spf13/cobra"
)

var (
	radarCmd = &cobra.Command{
		Use:   "radar",
		Short: "Weather radar (Composite reflectivity)",
		RunE:  radarFunc,
	}
)

func radarFunc(cmd *cobra.Command, args []string) error {
	return nil
}
