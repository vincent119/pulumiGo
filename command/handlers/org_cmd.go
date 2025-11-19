package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)


type OrgCommand struct {
    types.BaseHandler
}

func NewOrgCommand() *OrgCommand {
    cmd := &cobra.Command{
        Use:   "org",
        Short: "Manage Pulumi organizations",
        Long:  `Manage organizations in the Pulumi Cloud.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            return executeCommand(cmd, append([]string{"org"}, args...))
        },
    }

    return &OrgCommand{BaseHandler: types.BaseHandler{Command: cmd}}
}

func (h *OrgCommand) RegisterSubcommands(cmd *cobra.Command) {
    // 添加 get-default 子命令
    getDefaultCmd := &cobra.Command{
        Use:   "get-default",
        Short: "Get the default org for the current backend",
        Long:  `Get the default organization for the current backend.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"org", "get-default"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }
    cmd.AddCommand(getDefaultCmd)


    searchCmd := &cobra.Command{
        Use:   "search",
        Short: "Search for resources in Pulumi Cloud",
        Long:  `Search for resources in the Pulumi Cloud.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"org", "search"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }


    searchCmd.Flags().StringP("org", "o", "", "Organization name to filter search results")
    searchCmd.Flags().StringP("tag", "t", "", "Filter results with the given tag")
    searchCmd.Flags().StringP("project", "p", "", "Filter results with the given project")
    searchCmd.Flags().StringP("language", "l", "", "Filter results based on the language they use")
    searchCmd.Flags().StringP("stack", "s", "", "Filter results with the given stack name")
    searchCmd.Flags().String("resource-type", "", "Filter results with the given resource type")

    cmd.AddCommand(searchCmd)


    setDefaultCmd := &cobra.Command{
        Use:   "set-default [org]",
        Short: "Set the local default organization for the current backend",
        Long:  `Set the default organization for the current backend. This organization will be used for all commands that require an organization, unless otherwise specified.`,
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"org", "set-default"}
            cmdArgs = append(cmdArgs, args...)
            return executeCommand(cmd, cmdArgs)
        },
    }
    cmd.AddCommand(setDefaultCmd)
}
