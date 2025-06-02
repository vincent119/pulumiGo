package handlers

import (
    "github.com/spf13/cobra"
    "pulumiGo/interfaces"
)

// AboutCommand 處理 about 命令
type AboutCommand struct {
    interfaces.BaseHandler
}

func NewAboutCommand() *AboutCommand {
    cmd := &cobra.Command{
        Use:   "about",
        Short: "About Pulumi",
        Long:  `Show information about the Pulumi CLI and runtime.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"about"}, args...))
        },
    }

    return &AboutCommand{BaseHandler: interfaces.BaseHandler{Command: cmd}}
}

func (h *AboutCommand) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 env 子命令
    envCmd := &cobra.Command{
        Use:   "env",
        Short: "An overview of the environmental variables used by pulumi",
        Long:  `Displays information about the environmental variables that Pulumi uses.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"about", "env"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }
    cmd.AddCommand(envCmd)
}