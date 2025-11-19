// Package command 提供指令註冊與管理功能。
// 包含命令註冊表結構、命令添加與執行等功能模組。
// 此 package 被 main.go 使用，用於組織與執行所有命令。
package command

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// CommandRegistry 管理所有命令
type CommandRegistry struct {
    handlers []types.CommandHandler
}

// NewCommandRegistry 創建新的命令註冊表
func NewCommandRegistry() *CommandRegistry {
    return &CommandRegistry{
        handlers: []types.CommandHandler{},
    }
}

// Register 註冊命令處理器
func (r *CommandRegistry) Register(handler types.CommandHandler) {
    r.handlers = append(r.handlers, handler)
}

// AddToRootCommand 將所有註冊的命令添加到根命令
func (r *CommandRegistry) AddToRootCommand(rootCmd *cobra.Command) {
    for _, handler := range r.handlers {
        cmd := handler.GetCommand()
        rootCmd.AddCommand(cmd)
        handler.RegisterSubcommands(cmd)
    }
}

// ExecuteCmd 使用默認執行器執行命令
// 將命令與參數傳遞給默認執行器的 Execute 方法
func ExecuteCmd(cmd *cobra.Command, args []string) error {
    return DefaultExecutor.Execute(cmd, args)
}

// DefaultExecutor 是全局默認的命令執行器
var DefaultExecutor types.CommandExecutor
