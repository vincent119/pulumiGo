package interfaces

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