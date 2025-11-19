// Package types 定義了指令處理相關的介面與基礎結構。
// 包含指令處理器介面、基礎處理器結構、以及命令執行函式類型定義等。
package types

import (
	"github.com/spf13/cobra"
)


type CommandHandler interface {
    GetCommand() *cobra.Command
    RegisterSubcommands(cmd *cobra.Command)
}

type BaseHandler struct {
    Command *cobra.Command
}

func (h *BaseHandler) GetCommand() *cobra.Command {
    return h.Command
}

func (h *BaseHandler) RegisterSubcommands(cmd *cobra.Command) {
}

type CommandExecutor interface {
    Execute(cmd *cobra.Command, args []string) error
}

type ExecuteCmdFunc func(cmd *cobra.Command, args []string) error
