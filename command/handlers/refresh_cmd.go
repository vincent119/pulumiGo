package handlers

import (
	"fmt"
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

type RefreshHandler struct {
	types.BaseHandler
}

func NewRefreshHandler() *RefreshHandler {
	cmd := &cobra.Command{
		Use:   "refresh [url]",
		Short: "Refresh the resources in a stack",
		Long:  `Refresh the resources in a stack to match the current state of the infrastructure.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdArgs := []string{"refresh"}
			cmdArgs = append(cmdArgs, args...)

			// String array flags
			if cmd.Flag("exclude") != nil && cmd.Flag("exclude").Changed {
				exclude, _ := cmd.Flags().GetStringArray("exclude")
				for _, e := range exclude {
					cmdArgs = append(cmdArgs, "--exclude", e)
				}
			}
			if cmd.Flag("import-pending-creates") != nil && cmd.Flag("import-pending-creates").Changed {
				importPendingCreates, _ := cmd.Flags().GetStringArray("import-pending-creates")
				for _, ipc := range importPendingCreates {
					cmdArgs = append(cmdArgs, "--import-pending-creates", ipc)
				}
			}
			if cmd.Flag("remote-env") != nil && cmd.Flag("remote-env").Changed {
				remoteEnv, _ := cmd.Flags().GetStringArray("remote-env")
				for _, re := range remoteEnv {
					cmdArgs = append(cmdArgs, "--remote-env", re)
				}
			}
			if cmd.Flag("remote-env-secret") != nil && cmd.Flag("remote-env-secret").Changed {
				remoteEnvSecret, _ := cmd.Flags().GetStringArray("remote-env-secret")
				for _, res := range remoteEnvSecret {
					cmdArgs = append(cmdArgs, "--remote-env-secret", res)
				}
			}
			if cmd.Flag("remote-pre-run-command") != nil && cmd.Flag("remote-pre-run-command").Changed {
				remotePreRunCommand, _ := cmd.Flags().GetStringArray("remote-pre-run-command")
				for _, rprc := range remotePreRunCommand {
					cmdArgs = append(cmdArgs, "--remote-pre-run-command", rprc)
				}
			}
			if cmd.Flag("target") != nil && cmd.Flag("target").Changed {
				target, _ := cmd.Flags().GetStringArray("target")
				for _, t := range target {
					cmdArgs = append(cmdArgs, "--target", t)
				}
			}

			// String flags
			if cmd.Flag("config-file") != nil && cmd.Flag("config-file").Changed {
				configFile, _ := cmd.Flags().GetString("config-file")
				cmdArgs = append(cmdArgs, "--config-file", configFile)
			}
			if cmd.Flag("message") != nil && cmd.Flag("message").Changed {
				message, _ := cmd.Flags().GetString("message")
				cmdArgs = append(cmdArgs, "--message", message)
			}
			if cmd.Flag("remote-agent-pool-id") != nil && cmd.Flag("remote-agent-pool-id").Changed {
				remoteAgentPoolID, _ := cmd.Flags().GetString("remote-agent-pool-id")
				cmdArgs = append(cmdArgs, "--remote-agent-pool-id", remoteAgentPoolID)
			}
			if cmd.Flag("remote-executor-image") != nil && cmd.Flag("remote-executor-image").Changed {
				remoteExecutorImage, _ := cmd.Flags().GetString("remote-executor-image")
				cmdArgs = append(cmdArgs, "--remote-executor-image", remoteExecutorImage)
			}
			if cmd.Flag("remote-executor-image-password") != nil && cmd.Flag("remote-executor-image-password").Changed {
				remoteExecutorImagePassword, _ := cmd.Flags().GetString("remote-executor-image-password")
				cmdArgs = append(cmdArgs, "--remote-executor-image-password", remoteExecutorImagePassword)
			}
			if cmd.Flag("remote-executor-image-username") != nil && cmd.Flag("remote-executor-image-username").Changed {
				remoteExecutorImageUsername, _ := cmd.Flags().GetString("remote-executor-image-username")
				cmdArgs = append(cmdArgs, "--remote-executor-image-username", remoteExecutorImageUsername)
			}
			if cmd.Flag("remote-git-auth-access-token") != nil && cmd.Flag("remote-git-auth-access-token").Changed {
				remoteGitAuthAccessToken, _ := cmd.Flags().GetString("remote-git-auth-access-token")
				cmdArgs = append(cmdArgs, "--remote-git-auth-access-token", remoteGitAuthAccessToken)
			}
			if cmd.Flag("remote-git-auth-password") != nil && cmd.Flag("remote-git-auth-password").Changed {
				remoteGitAuthPassword, _ := cmd.Flags().GetString("remote-git-auth-password")
				cmdArgs = append(cmdArgs, "--remote-git-auth-password", remoteGitAuthPassword)
			}
			if cmd.Flag("remote-git-auth-ssh-private-key") != nil && cmd.Flag("remote-git-auth-ssh-private-key").Changed {
				remoteGitAuthSSHPrivateKey, _ := cmd.Flags().GetString("remote-git-auth-ssh-private-key")
				cmdArgs = append(cmdArgs, "--remote-git-auth-ssh-private-key", remoteGitAuthSSHPrivateKey)
			}
			if cmd.Flag("remote-git-auth-ssh-private-key-path") != nil && cmd.Flag("remote-git-auth-ssh-private-key-path").Changed {
				remoteGitAuthSSHPrivateKeyPath, _ := cmd.Flags().GetString("remote-git-auth-ssh-private-key-path")
				cmdArgs = append(cmdArgs, "--remote-git-auth-ssh-private-key-path", remoteGitAuthSSHPrivateKeyPath)
			}
			if cmd.Flag("remote-git-auth-username") != nil && cmd.Flag("remote-git-auth-username").Changed {
				remoteGitAuthUsername, _ := cmd.Flags().GetString("remote-git-auth-username")
				cmdArgs = append(cmdArgs, "--remote-git-auth-username", remoteGitAuthUsername)
			}
			if cmd.Flag("remote-git-branch") != nil && cmd.Flag("remote-git-branch").Changed {
				remoteGitBranch, _ := cmd.Flags().GetString("remote-git-branch")
				cmdArgs = append(cmdArgs, "--remote-git-branch", remoteGitBranch)
			}
			if cmd.Flag("remote-git-commit") != nil && cmd.Flag("remote-git-commit").Changed {
				remoteGitCommit, _ := cmd.Flags().GetString("remote-git-commit")
				cmdArgs = append(cmdArgs, "--remote-git-commit", remoteGitCommit)
			}
			if cmd.Flag("remote-git-repo-dir") != nil && cmd.Flag("remote-git-repo-dir").Changed {
				remoteGitRepoDir, _ := cmd.Flags().GetString("remote-git-repo-dir")
				cmdArgs = append(cmdArgs, "--remote-git-repo-dir", remoteGitRepoDir)
			}
			if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
				stack, _ := cmd.Flags().GetString("stack")
				cmdArgs = append(cmdArgs, "--stack", stack)
			}
			if cmd.Flag("suppress-permalink") != nil && cmd.Flag("suppress-permalink").Changed {
				suppressPermalink, _ := cmd.Flags().GetString("suppress-permalink")
				cmdArgs = append(cmdArgs, "--suppress-permalink", suppressPermalink)
			}

			// Boolean flags
			if cmd.Flag("clear-pending-creates") != nil && cmd.Flag("clear-pending-creates").Changed {
				cmdArgs = append(cmdArgs, "--clear-pending-creates")
			}
			if cmd.Flag("debug") != nil && cmd.Flag("debug").Changed {
				cmdArgs = append(cmdArgs, "--debug")
			}
			if cmd.Flag("diff") != nil && cmd.Flag("diff").Changed {
				cmdArgs = append(cmdArgs, "--diff")
			}
			if cmd.Flag("exclude-dependents") != nil && cmd.Flag("exclude-dependents").Changed {
				cmdArgs = append(cmdArgs, "--exclude-dependents")
			}
			if cmd.Flag("expect-no-changes") != nil && cmd.Flag("expect-no-changes").Changed {
				cmdArgs = append(cmdArgs, "--expect-no-changes")
			}
			if cmd.Flag("json") != nil && cmd.Flag("json").Changed {
				cmdArgs = append(cmdArgs, "--json")
			}
			if cmd.Flag("neo") != nil && cmd.Flag("neo").Changed {
				cmdArgs = append(cmdArgs, "--neo")
			}
			if cmd.Flag("preview-only") != nil && cmd.Flag("preview-only").Changed {
				cmdArgs = append(cmdArgs, "--preview-only")
			}
			if cmd.Flag("remote") != nil && cmd.Flag("remote").Changed {
				cmdArgs = append(cmdArgs, "--remote")
			}
			if cmd.Flag("remote-inherit-settings") != nil && cmd.Flag("remote-inherit-settings").Changed {
				cmdArgs = append(cmdArgs, "--remote-inherit-settings")
			}
			if cmd.Flag("remote-skip-install-dependencies") != nil && cmd.Flag("remote-skip-install-dependencies").Changed {
				cmdArgs = append(cmdArgs, "--remote-skip-install-dependencies")
			}
			if cmd.Flag("run-program") != nil && cmd.Flag("run-program").Changed {
				cmdArgs = append(cmdArgs, "--run-program")
			}
			if cmd.Flag("show-replacement-steps") != nil && cmd.Flag("show-replacement-steps").Changed {
				cmdArgs = append(cmdArgs, "--show-replacement-steps")
			}
			if cmd.Flag("show-sames") != nil && cmd.Flag("show-sames").Changed {
				cmdArgs = append(cmdArgs, "--show-sames")
			}
			if cmd.Flag("skip-pending-creates") != nil && cmd.Flag("skip-pending-creates").Changed {
				cmdArgs = append(cmdArgs, "--skip-pending-creates")
			}
			if cmd.Flag("skip-preview") != nil && cmd.Flag("skip-preview").Changed {
				cmdArgs = append(cmdArgs, "--skip-preview")
			}
			if cmd.Flag("suppress-outputs") != nil && cmd.Flag("suppress-outputs").Changed {
				cmdArgs = append(cmdArgs, "--suppress-outputs")
			}
			if cmd.Flag("suppress-progress") != nil && cmd.Flag("suppress-progress").Changed {
				cmdArgs = append(cmdArgs, "--suppress-progress")
			}
			if cmd.Flag("suppress-stream-logs") != nil && cmd.Flag("suppress-stream-logs").Changed {
				cmdArgs = append(cmdArgs, "--suppress-stream-logs")
			}
			if cmd.Flag("target-dependents") != nil && cmd.Flag("target-dependents").Changed {
				cmdArgs = append(cmdArgs, "--target-dependents")
			}
			if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
				cmdArgs = append(cmdArgs, "--yes")
			}

			// Int32 flag
			if cmd.Flag("parallel") != nil && cmd.Flag("parallel").Changed {
				parallel, _ := cmd.Flags().GetInt32("parallel")
				cmdArgs = append(cmdArgs, "--parallel", fmt.Sprintf("%d", parallel))
			}

			return executeCommand(cmd, cmdArgs)
		},
	}

	// String array flags
	cmd.Flags().StringArrayP("exclude", "x", nil, "Specify a resource URN to ignore. These resources will not be refreshed. Multiple resources can be specified using --exclude urn1 --exclude urn2. Wildcards (*, **) are also supported")
	cmd.Flags().StringArray("import-pending-creates", nil, "A list of form [[URN ID]...] describing the provider IDs of pending creates")
	cmd.Flags().StringArray("remote-env", nil, "[EXPERIMENTAL] Environment variables to use in the remote operation of the form NAME=value (e.g. --remote-env FOO=bar)")
	cmd.Flags().StringArray("remote-env-secret", nil, "[EXPERIMENTAL] Environment variables with secret values to use in the remote operation of the form NAME=secretvalue (e.g. --remote-env FOO=secret)")
	cmd.Flags().StringArray("remote-pre-run-command", nil, "[EXPERIMENTAL] Commands to run before the remote operation")
	cmd.Flags().StringArrayP("target", "t", nil, "Specify a single resource URN to refresh. Multiple resource can be specified using: --target urn1 --target urn2")

	// String flags
	cmd.Flags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	cmd.Flags().StringP("message", "m", "", "Optional message to associate with the update operation")
	cmd.Flags().String("remote-agent-pool-id", "", "[EXPERIMENTAL] The agent pool to use to run the deployment job. When empty, the Pulumi Cloud shared queue will be used.")
	cmd.Flags().String("remote-executor-image", "", "[EXPERIMENTAL] The Docker image to use for the executor")
	cmd.Flags().String("remote-executor-image-password", "", "[EXPERIMENTAL] The password for the credentials with access to the Docker image to use for the executor")
	cmd.Flags().String("remote-executor-image-username", "", "[EXPERIMENTAL] The username for the credentials with access to the Docker image to use for the executor")
	cmd.Flags().String("remote-git-auth-access-token", "", "[EXPERIMENTAL] Git personal access token")
	cmd.Flags().String("remote-git-auth-password", "", "[EXPERIMENTAL] Git password; for use with username or with an SSH private key")
	cmd.Flags().String("remote-git-auth-ssh-private-key", "", "[EXPERIMENTAL] Git SSH private key; use --remote-git-auth-password for the password, if needed")
	cmd.Flags().String("remote-git-auth-ssh-private-key-path", "", "[EXPERIMENTAL] Git SSH private key path; use --remote-git-auth-password for the password, if needed")
	cmd.Flags().String("remote-git-auth-username", "", "[EXPERIMENTAL] Git username")
	cmd.Flags().String("remote-git-branch", "", "[EXPERIMENTAL] Git branch to deploy; this is mutually exclusive with --remote-git-commit; either value needs to be specified")
	cmd.Flags().String("remote-git-commit", "", "[EXPERIMENTAL] Git commit hash of the commit to deploy (if used, HEAD will be in detached mode); this is mutually exclusive with --remote-git-branch; either value needs to be specified")
	cmd.Flags().String("remote-git-repo-dir", "", "[EXPERIMENTAL] The directory to work from in the project's source repository where Pulumi.yaml is located; used when Pulumi.yaml is not in the project source root")
	cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	cmd.Flags().String("suppress-permalink", "", "Suppress display of the state permalink")

	// Boolean flags
	cmd.Flags().Bool("clear-pending-creates", false, "Clear all pending creates, dropping them from the state")
	cmd.Flags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	cmd.Flags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	cmd.Flags().Bool("exclude-dependents", false, "Allows ignoring of dependent targets discovered but not specified in --exclude list")
	cmd.Flags().Bool("expect-no-changes", false, "Return an error if any changes occur during this refresh. This check happens after the refresh is applied")
	cmd.Flags().BoolP("json", "j", false, "Serialize the refresh diffs, operations, and overall output as JSON")
	cmd.Flags().Bool("neo", false, "Enable Pulumi Neo's assistance for improved CLI experience and insights (can also be set with PULUMI_NEO environment variable)")
	cmd.Flags().Bool("preview-only", false, "Only show a preview of the refresh, but don't perform the refresh itself")
	cmd.Flags().Bool("remote", false, "[EXPERIMENTAL] Run the operation remotely")
	cmd.Flags().Bool("remote-inherit-settings", false, "[EXPERIMENTAL] Inherit deployment settings from the current stack")
	cmd.Flags().Bool("remote-skip-install-dependencies", false, "[EXPERIMENTAL] Whether to skip the default dependency installation step")
	cmd.Flags().Bool("run-program", false, "Run the program to determine up-to-date state for providers to refresh resources")
	cmd.Flags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes instead of a single step")
	cmd.Flags().Bool("show-sames", false, "Show resources that needn't be updated because they haven't changed, alongside those that do")
	cmd.Flags().Bool("skip-pending-creates", false, "Skip importing pending creates in interactive mode")
	cmd.Flags().BoolP("skip-preview", "f", false, "Do not calculate a preview before performing the refresh")
	cmd.Flags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	cmd.Flags().Bool("suppress-progress", false, "Suppress display of periodic progress dots")
	cmd.Flags().Bool("suppress-stream-logs", false, "[EXPERIMENTAL] Suppress log streaming of the deployment job")
	cmd.Flags().Bool("target-dependents", false, "Allows updating of dependent targets discovered but not specified in --target list")
	cmd.Flags().BoolP("yes", "y", false, "Automatically approve and perform the refresh after previewing it")

	// Int32 flag
	cmd.Flags().Int32P("parallel", "p", 16, "Allow P resource operations to run in parallel at once (1 for no parallelism).")

	return &RefreshHandler{BaseHandler: types.BaseHandler{Command: cmd}}
}

func (h *RefreshHandler) RegisterSubcommands(cmd *cobra.Command) {
	// refresh 命令目前沒有子命令
}
