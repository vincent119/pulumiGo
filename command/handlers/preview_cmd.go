package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// PreviewCommand 處理 pulumi preview 命令
type PreviewCommand struct {
    types.BaseHandler
}

// NewPreviewCommand 創建新的 preview 命令處理器
func NewPreviewCommand() *PreviewCommand {
    cmd := &cobra.Command{
        Use:   "preview",
        Short: "Preview changes to resources in a stack",
        Long:  `Show a preview of changes that would be made by an update operation.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"preview"}, args...))
        },
    }

    // 添加 preview 命令專用的標誌
    cmd.Flags().BoolP("diff", "", false, "Display operation as a rich diff showing the overall change")
    cmd.Flags().Bool("show-sames", false, "Show resources that needn't be updated because they haven't changed")
    cmd.Flags().BoolP("json", "j", false, "Emit output as JSON")

    return &PreviewCommand{
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}
