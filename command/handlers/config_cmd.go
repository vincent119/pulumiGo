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

// newConfigSetCommand 創建 config set 子命令
func newConfigSetCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "set [key] [value]",
        Short: "Set a configuration value",
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"config", "set"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("path") != nil && cmd.Flag("path").Changed {
                cmdArgs = append(cmdArgs, "--path")
            }
            if cmd.Flag("plaintext") != nil && cmd.Flag("plaintext").Changed {
                cmdArgs = append(cmdArgs, "--plaintext")
            }
            if cmd.Flag("secret") != nil && cmd.Flag("secret").Changed {
                cmdArgs = append(cmdArgs, "--secret")
            }
            if cmd.Flag("type") != nil && cmd.Flag("type").Changed {
                typeVal, _ := cmd.Flags().GetString("type")
                cmdArgs = append(cmdArgs, "--type", typeVal)
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().Bool("path", false, "The key contains a path to a property in a map or list to set")
    cmd.Flags().Bool("plaintext", false, "Save the value as plaintext (unencrypted)")
    cmd.Flags().Bool("secret", false, "Encrypt the value instead of storing it in plaintext")
    cmd.Flags().String("type", "string", "Save the value as the given type. Allowed values: bool, int, string, float, object, json")

    return cmd
}

// newConfigGetCommand 創建 config get 子命令
func newConfigGetCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "get [key]",
        Short: "Get a configuration value",
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"config", "get"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("json") != nil && cmd.Flag("json").Changed {
                cmdArgs = append(cmdArgs, "--json")
            }
            if cmd.Flag("open") != nil && cmd.Flag("open").Changed {
                cmdArgs = append(cmdArgs, "--open")
            }
            if cmd.Flag("path") != nil && cmd.Flag("path").Changed {
                cmdArgs = append(cmdArgs, "--path")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().BoolP("json", "j", false, "Emit output as JSON")
    cmd.Flags().Bool("open", false, "Open and resolve any environments listed in the stack config")
    cmd.Flags().Bool("path", false, "The key contains a path to a property in a map or list to get")

    return cmd
}

func (h *ConfigHandler) RegisterSubcommands(cmd *cobra.Command) {

    cpCmd := newSubcommand("config", "cp [source] [destination]", "Copy configuration from one stack to another", 2)
    cmd.AddCommand(cpCmd)

    // 添加 get 子命令
    getCmd := newConfigGetCommand()
    cmd.AddCommand(getCmd)

    // 添加 set 子命令
    setCmd := newConfigSetCommand()
    cmd.AddCommand(setCmd)

    // 添加 rm 子命令
    rmCmd := newSubcommand("config", "rm [key]", "Remove a configuration value", 1)
    cmd.AddCommand(rmCmd)

    // 添加 env 子命令
    envCmd := newSubcommand("config", "env", "Manage ESC environments for a stack", 0)
    cmd.AddCommand(envCmd)

    // 添加 refresh 子命令
    refreshCmd := newSubcommand("config", "refresh", "Update the local configuration based on the most recent deployment of the stack", 0)
    cmd.AddCommand(refreshCmd)
}

