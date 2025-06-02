package handlers

import (
    "github.com/spf13/cobra"
    "pulumiGo/interfaces"
)

// PluginCommand 處理 pulumi plugin 命令
type PluginCommand struct {
    interfaces.BaseHandler
}

// NewPluginCommand 創建新的 plugin 命令處理器
func NewPluginCommand() *PluginCommand {
    cmd := &cobra.Command{
        Use:   "plugin",
        Short: "Manage plugins",
        Long:  `Manage language and resource provider plugins.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"plugin"}, args...))
        },
    }
    
    return &PluginCommand{
        BaseHandler: interfaces.BaseHandler{Command: cmd},
    }
}

// RegisterSubcommands 註冊 plugin 的子命令
func (h *PluginCommand) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 install 子命令
    installCmd := &cobra.Command{
        Use:   "install [kind] [name] [version]",
        Short: "Install one or more plugins",
        Long:  `Install one or more language or resource provider plugins.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"plugin", "install"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }
    
    // 為 install 命令添加標誌
    installCmd.Flags().Bool("exact", false, "Force installation of an exact version match")
    installCmd.Flags().String("file", "", "Install a plugin from a tarball file, instead of downloading it")
    installCmd.Flags().Bool("reinstall", false, "Reinstall the plugin even if it already exists")
    
    cmd.AddCommand(installCmd)
    
    // 添加 ls 子命令
    lsCmd := &cobra.Command{
        Use:   "ls",
        Short: "List plugins",
        Long:  `List all currently installed plugins.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"plugin", "ls"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }
    
    // 為 ls 命令添加標誌
    lsCmd.Flags().BoolP("json", "j", false, "Emit output as JSON")
    lsCmd.Flags().StringP("project", "p", "", "Project to list plugins for")
    
    cmd.AddCommand(lsCmd)
    
    // 添加 rm 子命令
    rmCmd := &cobra.Command{
        Use:   "rm [kind] [name] [version]",
        Short: "Remove one or more plugins",
        Long:  `Remove one or more plugins from the cache.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"plugin", "rm"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }
    
    // 為 rm 命令添加標誌
    rmCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")
    rmCmd.Flags().BoolP("all", "a", false, "Remove all plugins")
    
    cmd.AddCommand(rmCmd)
}