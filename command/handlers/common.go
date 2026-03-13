package handlers

import (
	"fmt"
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

// executeCommand is the package-internal command dispatcher.
func executeCommand(cmd *cobra.Command, args []string) error {
    if executeCommandFunc == nil {
        return fmt.Errorf("execute function not initialized")
    }
    return executeCommandFunc(cmd, args)
}

// executeCommandFunc 是全局執行函數指針，由 main 設置
var executeCommandFunc types.ExecuteCmdFunc

// InitExecuteFunc sets the command execution function.
func InitExecuteFunc(fn types.ExecuteCmdFunc) {
    executeCommandFunc = fn
}

// forwardStringArrayFlag appends --flagName value... for each value when the flag was changed.
func forwardStringArrayFlag(cmd *cobra.Command, args []string, flagName string) []string {
    if f := cmd.Flag(flagName); f != nil && f.Changed {
        vals, _ := cmd.Flags().GetStringArray(flagName)
        for _, v := range vals {
            args = append(args, "--"+flagName, v)
        }
    }
    return args
}

// forwardStringFlag appends --flagName value when the flag was changed.
func forwardStringFlag(cmd *cobra.Command, args []string, flagName string) []string {
    if f := cmd.Flag(flagName); f != nil && f.Changed {
        val, _ := cmd.Flags().GetString(flagName)
        args = append(args, "--"+flagName, val)
    }
    return args
}

// forwardBoolFlag appends --flagName when the flag was changed.
func forwardBoolFlag(cmd *cobra.Command, args []string, flagName string) []string {
    if f := cmd.Flag(flagName); f != nil && f.Changed {
        args = append(args, "--"+flagName)
    }
    return args
}

// forwardInt32Flag appends --flagName value when the flag was changed.
func forwardInt32Flag(cmd *cobra.Command, args []string, flagName string) []string {
    if f := cmd.Flag(flagName); f != nil && f.Changed {
        val, _ := cmd.Flags().GetInt32(flagName)
        args = append(args, "--"+flagName, fmt.Sprintf("%d", val))
    }
    return args
}
