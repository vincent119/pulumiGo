package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// ConfigHandler 處理 config 命令
type ConfigHandler struct {
    types.BaseHandler
}

// NewConfigHandler 創建 config 命令處理器
func NewConfigHandler() *ConfigHandler {
    cmd := &cobra.Command{
        Use:   "config",
        Short: "Manage configuration",
        Long:  `Manage configuration settings for the current stack.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            // 直接調用具體執行邏輯，而不是通過 command 包
            return executeCommand(cmd, append([]string{"config"}, args...))
        },
    }

    // 添加 config 特有的標誌
    cmd.PersistentFlags().Bool("show-secrets", false, "Show secret values when listing config")
    cmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")

    return &ConfigHandler{types.BaseHandler{Command: cmd}}
}

func (h *ConfigHandler) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 get 子命令
    getCmd := newSubcommand("config", "get [key]", "Get a configuration value", 1)
    cmd.AddCommand(getCmd)

    // 添加 set 子命令
    setCmd := newSubcommand("config", "set [key] [value]", "Set a configuration value", 2)
    cmd.AddCommand(setCmd)

    // 添加 rm 子命令
    rmCmd := newSubcommand("config", "rm [key]", "Remove a configuration value", 1)
    cmd.AddCommand(rmCmd)

    // 添加 env 子命令
    envCmd := newSubcommand("config", "env", "Manage ESC environments for a stack", 0)
    cmd.AddCommand(envCmd)
}

