package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// VersionCommand 處理 version 命令
type VersionCommand struct {
    types.BaseHandler
}

// NewVersionCommand 創建新的 version 命令處理器
func NewVersionCommand() *VersionCommand {
    cmd := &cobra.Command{
        Use:   "version",
        Short: "Display the current Pulumi version",
        Long:  `Display the current Pulumi CLI version and runtime information.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"version"}

            // 檢查是否需要詳細信息
            if cmd.Flag("verbose") != nil && cmd.Flag("verbose").Changed {
                cmdArgs = append(cmdArgs, "--verbose")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    // 添加 version 命令的標誌
    cmd.Flags().BoolP("verbose", "v", false, "Display additional version information")

    return &VersionCommand{
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}

// RegisterSubcommands 註冊 version 的子命令
func (h *VersionCommand) RegisterSubcommands(cmd *cobra.Command) {
    // version 命令沒有子命令
}
