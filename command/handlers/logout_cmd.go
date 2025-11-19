package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// LogoutHandler 處理 logout 命令
type LogoutHandler struct {
	types.BaseHandler
}

// NewLogoutHandler 創建 logout 命令處理器
func NewLogoutHandler() *LogoutHandler {
	cmd := &cobra.Command{
		Use:   "logout [url]",
		Short: "Log out of the Pulumi Cloud",
		Long:  `Log out of the Pulumi Cloud and remove saved credentials.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			pulumiArgs := []string{"logout"}
			pulumiArgs = append(pulumiArgs, args...)

			// 添加標誌
			if cmd.Flag("all") != nil && cmd.Flag("all").Changed {
				pulumiArgs = append(pulumiArgs, "--all")
			}
			if cmd.Flag("cloud-url") != nil && cmd.Flag("cloud-url").Changed {
				cloudURL, _ := cmd.Flags().GetString("cloud-url")
				pulumiArgs = append(pulumiArgs, "--cloud-url", cloudURL)
			}
			if cmd.Flag("local") != nil && cmd.Flag("local").Changed {
				pulumiArgs = append(pulumiArgs, "--local")
			}

			return executeCommand(cmd, pulumiArgs)
		},
	}

	cmd.Flags().Bool("all", false, "Logout of all backends")
	cmd.Flags().StringP("cloud-url", "c", "", "A cloud URL to log out of (defaults to current cloud)")
	cmd.Flags().BoolP("local", "l", false, "Log out of using local mode")

	return &LogoutHandler{
		BaseHandler: types.BaseHandler{Command: cmd},
	}
}
