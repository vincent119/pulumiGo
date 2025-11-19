package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// PluginCommand 處理 pulumi plugin 命令
type PluginCommand struct {
    types.BaseHandler
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
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}

// newPluginInstallCommand 創建 plugin install 子命令
func newPluginInstallCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "install [kind] [name] [version]",
        Short: "Install one or more plugins",
        Long:  `Install one or more language or resource provider plugins.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"plugin", "install"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("checksum") != nil && cmd.Flag("checksum").Changed {
                checksum, _ := cmd.Flags().GetString("checksum")
                cmdArgs = append(cmdArgs, "--checksum", checksum)
            }
            if cmd.Flag("exact") != nil && cmd.Flag("exact").Changed {
                cmdArgs = append(cmdArgs, "--exact")
            }
            if cmd.Flag("file") != nil && cmd.Flag("file").Changed {
                file, _ := cmd.Flags().GetString("file")
                cmdArgs = append(cmdArgs, "--file", file)
            }
            if cmd.Flag("reinstall") != nil && cmd.Flag("reinstall").Changed {
                cmdArgs = append(cmdArgs, "--reinstall")
            }
            if cmd.Flag("server") != nil && cmd.Flag("server").Changed {
                server, _ := cmd.Flags().GetString("server")
                cmdArgs = append(cmdArgs, "--server", server)
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().String("checksum", "", "The expected SHA256 checksum for the plugin archive")
    cmd.Flags().Bool("exact", false, "Force installation of an exact version match")
    cmd.Flags().StringP("file", "f", "", "Install a plugin from a binary, folder or tarball file, instead of downloading it")
    cmd.Flags().Bool("reinstall", false, "Reinstall a plugin even if it already exists")
    cmd.Flags().String("server", "", "A URL to download plugins from")

    return cmd
}

// RegisterSubcommands 註冊 plugin 的子命令
func (h *PluginCommand) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 install 子命令
    installCmd := newPluginInstallCommand()
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
