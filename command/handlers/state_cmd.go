package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// StateCommand 處理 pulumi state 命令
type StateCommand struct {
    types.BaseHandler
}

// NewStateCommand 創建新的 state 命令處理器
func NewStateCommand() *StateCommand {
    cmd := &cobra.Command{
        Use:   "state",
        Short: "Edit the current stack's state",
        Long:  `Edit the current stack's state, including deleting and protecting resources.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"state"}, args...))
        },
    }

    return &StateCommand{
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}

// RegisterSubcommands 註冊 state 的子命令
func (h *StateCommand) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 delete 子命令
    deleteCmd := &cobra.Command{
        Use:   "delete <resource-urn>",
        Short: "Deletes a resource from the stack's state",
        Long:  `Deletes a resource from the stack's state, without performing the actual deletion of cloud resources.`,
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "delete"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }

    // 為 delete 命令添加標誌
    deleteCmd.Flags().Bool("force", false, "Force deletion of protected resources")
    deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")

    cmd.AddCommand(deleteCmd)

    // 添加 protect 子命令
    protectCmd := &cobra.Command{
        Use:   "protect <resource-urn>",
        Short: "Mark a resource as protected",
        Long:  `Mark a resource as protected. Protected resources cannot be deleted or replaced.`,
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "protect"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.AddCommand(protectCmd)

    // 添加 unprotect 子命令
    unprotectCmd := &cobra.Command{
        Use:   "unprotect <resource-urn>",
        Short: "Unmark a resource as protected",
        Long:  `Unmark a resource as protected, allowing it to be deleted or replaced.`,
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "unprotect"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.AddCommand(unprotectCmd)
}
