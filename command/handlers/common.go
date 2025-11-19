package handlers

import (
	"pulumiGo/types"
	"strings"

	"github.com/spf13/cobra"
)

// newSubcommand 創建子命令
func newSubcommand(parentCmd string, use string, desc string, argCount int) *cobra.Command {
    cmd := &cobra.Command{
        Use:   use,
        Short: desc,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{parentCmd}

            // 從 Use 中提取主要命令名稱（例如從 "get [key]" 中提取 "get"）
            subCmd := use
            if space := strings.Index(use, " "); space > 0 {
                subCmd = use[:space]
            }

            cmdArgs = append(cmdArgs, subCmd)
            cmdArgs = append(cmdArgs, args...)

            return executeCommand(cmd, cmdArgs)
        },
    }

    if argCount > 0 {
        cmd.Args = cobra.ExactArgs(argCount)
    }

    return cmd
}

// executeCommand 是處理器包內部的命令執行函數
func executeCommand(cmd *cobra.Command, args []string) error {
    return executeCommandFunc(cmd, args)
}

// executeCommandFunc 是全局執行函數指針，由 main 設置
var executeCommandFunc types.ExecuteCmdFunc

// InitExecuteFunc 設置命令執行函數
func InitExecuteFunc(fn types.ExecuteCmdFunc) {
    executeCommandFunc = fn
}
