package handlers

import (
	"testing"

	"github.com/spf13/cobra"
)

func newTestCmd() *cobra.Command {
	return &cobra.Command{Use: "test"}
}

func TestForwardBoolFlag(t *testing.T) {
	t.Run("flag changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().Bool("yes", false, "")
		_ = cmd.Flags().Set("yes", "true")

		args := forwardBoolFlag(cmd, []string{"base"}, "yes")
		if len(args) != 2 || args[1] != "--yes" {
			t.Errorf("got %v", args)
		}
	})

	t.Run("flag not changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().Bool("yes", false, "")

		args := forwardBoolFlag(cmd, []string{"base"}, "yes")
		if len(args) != 1 {
			t.Errorf("expected no append, got %v", args)
		}
	})

	t.Run("flag not defined", func(t *testing.T) {
		cmd := newTestCmd()
		args := forwardBoolFlag(cmd, []string{"base"}, "nonexistent")
		if len(args) != 1 {
			t.Errorf("expected no append, got %v", args)
		}
	})
}

func TestForwardStringFlag(t *testing.T) {
	t.Run("flag changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().String("stack", "", "")
		_ = cmd.Flags().Set("stack", "mystack")

		args := forwardStringFlag(cmd, []string{"base"}, "stack")
		if len(args) != 3 || args[1] != "--stack" || args[2] != "mystack" {
			t.Errorf("got %v", args)
		}
	})

	t.Run("flag not changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().String("stack", "", "")

		args := forwardStringFlag(cmd, []string{"base"}, "stack")
		if len(args) != 1 {
			t.Errorf("expected no append, got %v", args)
		}
	})
}

func TestForwardStringArrayFlag(t *testing.T) {
	t.Run("multiple values", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().StringArray("target", nil, "")
		_ = cmd.Flags().Set("target", "urn:a")
		_ = cmd.Flags().Set("target", "urn:b")

		args := forwardStringArrayFlag(cmd, []string{"base"}, "target")
		// expect: base --target urn:a --target urn:b
		if len(args) != 5 {
			t.Errorf("got %v", args)
		}
	})

	t.Run("flag not changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().StringArray("target", nil, "")

		args := forwardStringArrayFlag(cmd, []string{"base"}, "target")
		if len(args) != 1 {
			t.Errorf("expected no append, got %v", args)
		}
	})
}

func TestForwardInt32Flag(t *testing.T) {
	t.Run("flag changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().Int32("parallel", 16, "")
		_ = cmd.Flags().Set("parallel", "4")

		args := forwardInt32Flag(cmd, []string{"base"}, "parallel")
		if len(args) != 3 || args[1] != "--parallel" || args[2] != "4" {
			t.Errorf("got %v", args)
		}
	})

	t.Run("flag not changed", func(t *testing.T) {
		cmd := newTestCmd()
		cmd.Flags().Int32("parallel", 16, "")

		args := forwardInt32Flag(cmd, []string{"base"}, "parallel")
		if len(args) != 1 {
			t.Errorf("expected no append, got %v", args)
		}
	})
}
