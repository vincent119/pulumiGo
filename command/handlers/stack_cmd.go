package handlers

import (
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

type StackHandler struct {
    types.BaseHandler
}


func NewStackHandler() *StackHandler {
    cmd := &cobra.Command{
        Use:   "stack",
        Short: "Manage stacks",
        Long:  `Manage stacks (e.g., ls, select, remove).`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"stack"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("show-ids") != nil && cmd.Flag("show-ids").Changed {
                cmdArgs = append(cmdArgs, "--show-ids")
            }
            if cmd.Flag("show-name") != nil && cmd.Flag("show-name").Changed {
                cmdArgs = append(cmdArgs, "--show-name")
            }
            if cmd.Flag("show-secrets") != nil && cmd.Flag("show-secrets").Changed {
                cmdArgs = append(cmdArgs, "--show-secrets")
            }
            if cmd.Flag("show-urns") != nil && cmd.Flag("show-urns").Changed {
                cmdArgs = append(cmdArgs, "--show-urns")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().BoolP("show-ids", "i", false, "Display each resource's provider-assigned unique ID")
    cmd.Flags().Bool("show-name", false, "Display only the stack name")
    cmd.Flags().Bool("show-secrets", false, "Display stack outputs which are marked as secret in plaintext")
    cmd.Flags().BoolP("show-urns", "u", false, "Display each resource's Pulumi-assigned globally unique URN")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")

    return &StackHandler{BaseHandler: types.BaseHandler{Command: cmd}}
}

// newStackLsCommand 創建 stack ls 子命令
func newStackLsCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "ls",
        Short: "List stacks",
        Long:  `List all stacks.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"stack", "ls"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("all") != nil && cmd.Flag("all").Changed {
                cmdArgs = append(cmdArgs, "--all")
            }
            if cmd.Flag("json") != nil && cmd.Flag("json").Changed {
                cmdArgs = append(cmdArgs, "--json")
            }
            if cmd.Flag("organization") != nil && cmd.Flag("organization").Changed {
                organization, _ := cmd.Flags().GetString("organization")
                cmdArgs = append(cmdArgs, "--organization", organization)
            }
            if cmd.Flag("project") != nil && cmd.Flag("project").Changed {
                project, _ := cmd.Flags().GetString("project")
                cmdArgs = append(cmdArgs, "--project", project)
            }
            if cmd.Flag("tag") != nil && cmd.Flag("tag").Changed {
                tag, _ := cmd.Flags().GetString("tag")
                cmdArgs = append(cmdArgs, "--tag", tag)
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().BoolP("all", "a", false, "List all stacks instead of just stacks for the current project")
    cmd.Flags().BoolP("json", "j", false, "Emit output as JSON")
    cmd.Flags().StringP("organization", "o", "", "Filter returned stacks to those in a specific organization")
    cmd.Flags().StringP("project", "p", "", "Filter returned stacks to those with a specific project name")
    cmd.Flags().StringP("tag", "t", "", "Filter returned stacks to those in a specific tag (tag-name or tag-name=tag-value)")

    return cmd
}

// newStackExportCommand 創建 stack export 子命令
func newStackExportCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "export",
        Short: "Export the current stack's state",
        Long:  `Export the current stack's state to stdout or a file.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"stack", "export"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("file") != nil && cmd.Flag("file").Changed {
                file, _ := cmd.Flags().GetString("file")
                cmdArgs = append(cmdArgs, "--file", file)
            }
            if cmd.Flag("show-secrets") != nil && cmd.Flag("show-secrets").Changed {
                cmdArgs = append(cmdArgs, "--show-secrets")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("version") != nil && cmd.Flag("version").Changed {
                version, _ := cmd.Flags().GetString("version")
                cmdArgs = append(cmdArgs, "--version", version)
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().String("file", "", "A filename to write stack output to")
    cmd.Flags().Bool("show-secrets", false, "Emit secrets in plaintext in exported stack. Defaults to false")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
    cmd.Flags().String("version", "", "Previous stack version to export. (If unset, will export the latest.)")

    return cmd
}

// newStackImportCommand 創建 stack import 子命令
func newStackImportCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "import",
        Short: "Import a deployment from standard in or a file",
        Long:  `Import a deployment from standard in or a file into an existing stack.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"stack", "import"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("file") != nil && cmd.Flag("file").Changed {
                file, _ := cmd.Flags().GetString("file")
                cmdArgs = append(cmdArgs, "--file", file)
            }
            if cmd.Flag("force") != nil && cmd.Flag("force").Changed {
                cmdArgs = append(cmdArgs, "--force")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().String("file", "", "A filename to read stack input from")
    cmd.Flags().BoolP("force", "f", false, "Force the import to occur, even if apparent errors are discovered beforehand (not recommended)")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")

    return cmd
}

// newStackInitCommand 創建 stack init 子命令
func newStackInitCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "init [name]",
        Short: "Create a new stack",
        Long:  `Create a new stack and set it as the active stack.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"stack", "init"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("copy-config-from") != nil && cmd.Flag("copy-config-from").Changed {
                copyConfigFrom, _ := cmd.Flags().GetString("copy-config-from")
                cmdArgs = append(cmdArgs, "--copy-config-from", copyConfigFrom)
            }
            if cmd.Flag("no-select") != nil && cmd.Flag("no-select").Changed {
                cmdArgs = append(cmdArgs, "--no-select")
            }
            if cmd.Flag("secrets-provider") != nil && cmd.Flag("secrets-provider").Changed {
                secretsProvider, _ := cmd.Flags().GetString("secrets-provider")
                cmdArgs = append(cmdArgs, "--secrets-provider", secretsProvider)
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("teams") != nil && cmd.Flag("teams").Changed {
                teams, _ := cmd.Flags().GetStringArray("teams")
                for _, team := range teams {
                    cmdArgs = append(cmdArgs, "--teams", team)
                }
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().String("copy-config-from", "", "The name of the stack to copy existing config from")
    cmd.Flags().Bool("no-select", false, "Do not select the stack")
    cmd.Flags().String("secrets-provider", "", "The type of the provider that should be used to encrypt and decrypt secrets")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to create")
    cmd.Flags().StringArray("teams", nil, "A list of team names that should have permission to read and update this stack")

    return cmd
}

// newStackRmCommand 創建 stack rm 子命令
func newStackRmCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "rm [stack-name]",
        Short: "Remove a stack and its configuration",
        Long:  `Remove a stack and its configuration from the system.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"stack", "rm"}
            cmdArgs = append(cmdArgs, args...)

            // 添加標誌
            if cmd.Flag("force") != nil && cmd.Flag("force").Changed {
                cmdArgs = append(cmdArgs, "--force")
            }
            if cmd.Flag("preserve-config") != nil && cmd.Flag("preserve-config").Changed {
                cmdArgs = append(cmdArgs, "--preserve-config")
            }
            if cmd.Flag("remove-backups") != nil && cmd.Flag("remove-backups").Changed {
                cmdArgs = append(cmdArgs, "--remove-backups")
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
                cmdArgs = append(cmdArgs, "--yes")
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    cmd.Flags().BoolP("force", "f", false, "Forces deletion of the stack, leaving behind any resources managed by the stack")
    cmd.Flags().Bool("preserve-config", false, "Do not delete the corresponding Pulumi.<stack-name>.yaml configuration file for the stack")
    cmd.Flags().Bool("remove-backups", false, "Additionally remove backups of the stack, if using the DIY backend")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
    cmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts, and proceed with removal anyway")

    return cmd
}

func (h *StackHandler) RegisterSubcommands(cmd *cobra.Command) {

    lsCmd := newStackLsCommand()
    cmd.AddCommand(lsCmd)

    exportCmd := newStackExportCommand()
    cmd.AddCommand(exportCmd)

    importCmd := newStackImportCommand()
    cmd.AddCommand(importCmd)

    initCmd := newStackInitCommand()
    cmd.AddCommand(initCmd)

    rmCmd := newStackRmCommand()
    cmd.AddCommand(rmCmd)

    selectCmd := &cobra.Command{
        Use:   "select [stack]",
        Short: "Select a stack",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            originalArgs := []string{"stack", "select"}
            originalArgs = append(originalArgs, args...)
            return executeCommand(cmd, originalArgs)
        },
    }
    cmd.AddCommand(selectCmd)
}
