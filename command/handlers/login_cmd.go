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
            a := append([]string{"login"}, args...)
            for _, f := range []string{"cloud-url", "default-org"} {
                a = forwardStringFlag(cmd, a, f)
            }
            for _, f := range []string{"insecure", "interactive", "local"} {
                a = forwardBoolFlag(cmd, a, f)
            }
            return executeCommand(cmd, a)
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
