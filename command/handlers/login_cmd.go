package handlers

import (
    "github.com/spf13/cobra"
    "pulumiGo/interfaces"
)

// LoginHandler 處理 login 命令
type LoginHandler struct {
    interfaces.BaseHandler
    localFlag bool
}

// NewLoginHandler 創建 login 命令處理器
func NewLoginHandler() *LoginHandler {
    handler := &LoginHandler{}
    
    cmd := &cobra.Command{
        Use:   "login [url]",
        Short: "Log in to the Pulumi Cloud",
        Long:  `Log in to the Pulumi Cloud backend at the specified URL or the default.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            pulumiArgs := []string{"login"}
            pulumiArgs = append(pulumiArgs, args...)
            if handler.localFlag {
                pulumiArgs = append(pulumiArgs, "--local")
            }
            return executeCommand(cmd, pulumiArgs)
        },
    }
    
    // 添加 local 標誌
    cmd.Flags().BoolVar(&handler.localFlag, "local", false, "Use Pulumi in local mode")
    
    handler.Command = cmd
    return handler
}