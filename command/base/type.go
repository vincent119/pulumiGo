// Package base 提供指令處理的基礎結構與介面定義。
// 包含基本的指令處理器結構、指令介面定義、以及執行函式類型定義等。
package base

import (
    "github.com/spf13/cobra"
)


type CommandHandler interface {

    GetCommand() *cobra.Command
}


type BaseHandler struct {
    Command *cobra.Command
}

func (h *BaseHandler) GetCommand() *cobra.Command {
    return h.Command
}


type ExecuteCmdFunc func(cmd *cobra.Command, args []string) error


type CommandExecutor interface {
    Execute(cmd *cobra.Command, args []string) error
}
