package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"pulumiGo/command"
	"pulumiGo/command/handlers"

	"github.com/spf13/cobra"
)

func TestRootCommandSuppressesUsageOnExecutionError(t *testing.T) {
	cmd := newRootCommand("test")
	t.Cleanup(func() {
		handlers.InitExecuteFunc(command.ExecuteCmd)
	})

	expectedErr := errors.New("exit status 1")
	handlers.InitExecuteFunc(func(_ *cobra.Command, _ []string) error {
		return expectedErr
	})

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	cmd.SetArgs([]string{"up"})

	err := cmd.Execute()
	if !errors.Is(err, expectedErr) {
		t.Fatalf("Execute() error = %v, want %v", err, expectedErr)
	}

	errOutput := stderr.String()
	for _, unwanted := range []string{"Usage:", "Flags:", "Error:"} {
		if strings.Contains(errOutput, unwanted) {
			t.Fatalf("stderr contains %q:\n%s", unwanted, errOutput)
		}
	}
}
