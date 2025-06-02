package command

import (
    "github.com/spf13/cobra"
    "pulumiGo/interfaces"
)

// CommandRegistry 管理所有命令
type CommandRegistry struct {
    handlers []interfaces.CommandHandler
}

// NewCommandRegistry 創建新的命令註冊表
func NewCommandRegistry() *CommandRegistry {
    return &CommandRegistry{
        handlers: []interfaces.CommandHandler{},
    }
}

// Register 註冊命令處理器
func (r *CommandRegistry) Register(handler interfaces.CommandHandler) {
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

// 簡化版的執行命令功能
func ExecuteCmd(cmd *cobra.Command, args []string) error {
    return DefaultExecutor.Execute(cmd, args)
}

// 默認執行器
var DefaultExecutor interfaces.CommandExecutor