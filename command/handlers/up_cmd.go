package handlers

import (
    //"pulumiGo/command"
    "github.com/spf13/cobra"
    "pulumiGo/interfaces"
)

// UpCommand 處理 pulumi up 命令
type UpCommand struct {
    interfaces.BaseHandler
}

// NewUpCommand 創建新的 up 命令處理器
func NewUpCommand() *UpCommand {
    cmd := &cobra.Command{
        Use:   "up",
        Short: "Create or update resources in a stack",
        Long:  `Update the resources in a stack to match the current configuration.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"up"}, args...))
        },
    }
    
    // 添加 up 命令專用的標誌
    cmd.Flags().BoolP("yes", "y", false, "Automatically approve and perform the update")
    cmd.Flags().BoolP("diff", "", false, "Display operation as a rich diff showing the overall change")
    cmd.Flags().Bool("skip-preview", false, "Do not perform a preview before performing the update")
    
    return &UpCommand{
        BaseHandler: interfaces.BaseHandler{Command: cmd},
    }
}