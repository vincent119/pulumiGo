package command

import (
    "github.com/spf13/cobra"
    "pulumiGo/interfaces"
)

// SimpleCommandHandler 簡單命令處理器
type SimpleCommandHandler struct {
    interfaces.BaseHandler
}

// NewSimpleCommand 創建簡單的 Pulumi 命令
func NewSimpleCommand(name, short, long string) *SimpleCommandHandler {
    cmd := &cobra.Command{
        Use:   name,
        Short: short,
        Long:  long,
        RunE: func(cmd *cobra.Command, args []string) error {
            return ExecuteCmd(cmd, append([]string{name}, args...))
        },
    }
    
    return &SimpleCommandHandler{interfaces.BaseHandler{Command: cmd}}
}