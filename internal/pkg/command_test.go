package pkg

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestExecuteCommandC(t *testing.T) {
	cmd := &cobra.Command{
		Run: func(c *cobra.Command, args []string) {
			c.Print(strings.Join(args, ", "))
		},
	}

	_, output, err := ExecuteCommandC(cmd, "Hello", "World")
	if output != "Hello, World" {
		t.Errorf("Unexpected output: %v", output)
	}
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCommand_FlagInvalidInput(t *testing.T) {
	cmd := &cobra.Command{
		Run: func(c *cobra.Command, args []string) {},
	}
	cmd.Flags().IntP("number", "n", 0, "Number")

	if _, _, err := ExecuteCommandC(cmd, "--number", "string"); err == nil {
		t.Error("Invalid flag value should return error")
	}
}
