package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// LoginHandler 處理 login 命令
type LoginHandler struct {
    types.BaseHandler
}

// NewLoginHandler 創建 login 命令處理器
func NewLoginHandler() *LoginHandler {
    cmd := &cobra.Command{
        Use:   "login [url]",
        Short: "Log in to the Pulumi Cloud",
        Long:  `Log in to the Pulumi Cloud backend at the specified URL or the default.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            pulumiArgs := []string{"login"}
            pulumiArgs = append(pulumiArgs, args...)

            // 添加標誌
            if cmd.Flag("cloud-url") != nil && cmd.Flag("cloud-url").Changed {
                cloudURL, _ := cmd.Flags().GetString("cloud-url")
                pulumiArgs = append(pulumiArgs, "--cloud-url", cloudURL)
            }
            if cmd.Flag("default-org") != nil && cmd.Flag("default-org").Changed {
                defaultOrg, _ := cmd.Flags().GetString("default-org")
                pulumiArgs = append(pulumiArgs, "--default-org", defaultOrg)
            }
            if cmd.Flag("insecure") != nil && cmd.Flag("insecure").Changed {
                pulumiArgs = append(pulumiArgs, "--insecure")
            }
            if cmd.Flag("interactive") != nil && cmd.Flag("interactive").Changed {
                pulumiArgs = append(pulumiArgs, "--interactive")
            }
            if cmd.Flag("local") != nil && cmd.Flag("local").Changed {
                pulumiArgs = append(pulumiArgs, "--local")
            }

            return executeCommand(cmd, pulumiArgs)
        },
    }

    cmd.Flags().StringP("cloud-url", "c", "", "A cloud URL to log in to")
    cmd.Flags().String("default-org", "", "A default org to associate with the login")
    cmd.Flags().Bool("insecure", false, "Allow insecure server connections when using SSL")
    cmd.Flags().Bool("interactive", false, "Show interactive login options based on known identity providers")
    cmd.Flags().BoolP("local", "l", false, "Use Pulumi in local-only mode")

    return &LoginHandler{
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}
