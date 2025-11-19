package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

type StackHandler struct {
    types.BaseHandler
}


func NewStackHandler() *StackHandler {
    cmd := &cobra.Command{
        Use:   "stack",
        Short: "Manage stacks",
        Long:  `Manage stacks (e.g., ls, select, remove).`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"stack"}, args...))
        },
    }

    return &StackHandler{BaseHandler: types.BaseHandler{Command: cmd}}
}


func (h *StackHandler) RegisterSubcommands(cmd *cobra.Command) {

    lsCmd := newSubcommand("stack", "ls", "List stacks", 0)
    cmd.AddCommand(lsCmd)

    selectCmd := &cobra.Command{
        Use:   "select [stack]",
        Short: "Select a stack",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            originalArgs := []string{"stack", "select"}
            originalArgs = append(originalArgs, args...)
            return executeCommand(cmd, originalArgs)
        },
    }
    cmd.AddCommand(selectCmd)
}
