// Package handlers 提供所有與指令執行邏輯相關的處理函式，
// 包含參數驗證、執行流程控制、與結果回傳等功能模組。
// 此 package 被 pulumiGo/command 使用，用於統一 API 輸出行為。
package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// AboutCommand 處理 about 命令
type AboutCommand struct {
    types.BaseHandler
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

    return &AboutCommand{BaseHandler: types.BaseHandler{Command: cmd}}
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
