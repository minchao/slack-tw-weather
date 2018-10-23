package weather

import (
	"reflect"
	"testing"
)

func Test_prepareArgs(t *testing.T) {
	tests := []struct {
		args []string
		want []string
	}{
		{
			args: []string{},
			want: []string{},
		},
		{
			args: []string{""},
			want: []string{},
		},
		{
			args: []string{"-h"},
			want: []string{"-h"},
		},
		{
			args: []string{"--version"},
			want: []string{"--version"},
		},
		{
			args: []string{"radar"},
			want: []string{"radar"},
		},
		{
			args: []string{"forecast", "Taichung"},
			want: []string{"forecast", "Taichung"},
		},
		{
			args: []string{"Taichung"},
			want: []string{"forecast", "Taichung"},
		},
	}

	for i, test := range tests {
		if got := prepareArgs(rootCmd, test.args); !reflect.DeepEqual(test.want, got) {
			t.Errorf("(%v) Expected: %v, got: %v", i, test.want, got)
		}
	}
}
