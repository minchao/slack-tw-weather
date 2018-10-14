package app

import (
	"testing"

	"github.com/minchao/slack-tw-weather/internal/pkg"
)

func Test_rootCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			"weather",
			[]string{""},
		},
		{
			"forecast",
			[]string{"臺中市"},
		},
		{
			"forecast",
			[]string{"Taichung"},
		},
		{
			"forecast",
			[]string{"forecast", "Taichung"},
		},
		{
			"radar",
			[]string{"radar"},
		},
	}

	for i, test := range tests {
		c, _, err := pkg.ExecuteCommandC(rootCmd, test.args...)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if c.Name() != test.name {
			t.Errorf("(%v) Expected: %v, got: %v", i, test.name, c.Name())
		}
	}
}
