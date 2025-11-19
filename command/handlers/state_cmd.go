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

// newStateDeleteCommand 創建 state delete 子命令
func newStateDeleteCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "delete [resource-urn]",
        Short: "Deletes a resource from the stack's state",
        Long:  `Deletes a resource from the stack's state, without performing the actual deletion of cloud resources.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "delete"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("all") != nil && cmd.Flag("all").Changed {
                cmdArgs = append(cmdArgs, "--all")
            }
            if cmd.Flag("force") != nil && cmd.Flag("force").Changed {
                cmdArgs = append(cmdArgs, "--force")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("target-dependents") != nil && cmd.Flag("target-dependents").Changed {
                cmdArgs = append(cmdArgs, "--target-dependents")
            }
            if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
                cmdArgs = append(cmdArgs, "--yes")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().Bool("all", false, "Delete all resources in the stack")
    cmd.Flags().Bool("force", false, "Force deletion of protected resources")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
    cmd.Flags().Bool("target-dependents", false, "Delete the URN and all its dependents")
    cmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")

    return cmd
}

// newStateProtectCommand 創建 state protect 子命令
func newStateProtectCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "protect [resource-urn...]",
        Short: "Mark a resource as protected",
        Long:  `Mark a resource as protected. Protected resources cannot be deleted or replaced.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "protect"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("all") != nil && cmd.Flag("all").Changed {
                cmdArgs = append(cmdArgs, "--all")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
                cmdArgs = append(cmdArgs, "--yes")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().Bool("all", false, "Protect all resources in the checkpoint")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
    cmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")

    return cmd
}

// newStateUnprotectCommand 創建 state unprotect 子命令
func newStateUnprotectCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "unprotect [resource-urn...]",
        Short: "Unmark a resource as protected",
        Long:  `Unmark a resource as protected, allowing it to be deleted or replaced.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "unprotect"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("all") != nil && cmd.Flag("all").Changed {
                cmdArgs = append(cmdArgs, "--all")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
                cmdArgs = append(cmdArgs, "--yes")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().Bool("all", false, "Unprotect all resources in the checkpoint")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
    cmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")

    return cmd
}

// newStateMoveCommand 創建 state move 子命令
func newStateMoveCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "move [urn...]",
        Short: "Move resources from one stack to another",
        Long:  `Move resources from one stack to another.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "move"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("dest") != nil && cmd.Flag("dest").Changed {
                dest, _ := cmd.Flags().GetString("dest")
                cmdArgs = append(cmdArgs, "--dest", dest)
            }
            if cmd.Flag("include-parents") != nil && cmd.Flag("include-parents").Changed {
                cmdArgs = append(cmdArgs, "--include-parents")
            }
            if cmd.Flag("source") != nil && cmd.Flag("source").Changed {
                source, _ := cmd.Flags().GetString("source")
                cmdArgs = append(cmdArgs, "--source", source)
            }
            if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
                cmdArgs = append(cmdArgs, "--yes")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().String("dest", "", "The name of the stack to move resources to")
    cmd.Flags().Bool("include-parents", false, "Include all the parents of the moved resources as well")
    cmd.Flags().String("source", "", "The name of the stack to move resources from")
    cmd.Flags().BoolP("yes", "y", false, "Automatically approve and perform the move")

    return cmd
}

// newStateRenameCommand 創建 state rename 子命令
func newStateRenameCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "rename [resource-urn] [new-name]",
        Short: "Rename a resource in the stack's state",
        Long:  `Rename a resource in the stack's state.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"state", "rename"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
                cmdArgs = append(cmdArgs, "--yes")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
    cmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")

    return cmd
}

// RegisterSubcommands 註冊 state 的子命令
func (h *StateCommand) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 delete 子命令
    deleteCmd := newStateDeleteCommand()
    cmd.AddCommand(deleteCmd)

    // 添加 protect 子命令
    protectCmd := newStateProtectCommand()
    cmd.AddCommand(protectCmd)

    // 添加 unprotect 子命令
    unprotectCmd := newStateUnprotectCommand()
    cmd.AddCommand(unprotectCmd)

    // 添加 move 子命令
    moveCmd := newStateMoveCommand()
    cmd.AddCommand(moveCmd)

    // 添加 rename 子命令
    renameCmd := newStateRenameCommand()
    cmd.AddCommand(renameCmd)
}
