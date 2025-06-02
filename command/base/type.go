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