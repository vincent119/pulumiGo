package handlers

import (
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
			a := append([]string{"refresh"}, args...)

			for _, f := range []string{
				"exclude", "import-pending-creates", "remote-env",
				"remote-env-secret", "remote-pre-run-command", "target",
			} {
				a = forwardStringArrayFlag(cmd, a, f)
			}
			for _, f := range []string{
				"config-file", "message", "remote-agent-pool-id",
				"remote-executor-image", "remote-executor-image-password",
				"remote-executor-image-username", "remote-git-auth-access-token",
				"remote-git-auth-password", "remote-git-auth-ssh-private-key",
				"remote-git-auth-ssh-private-key-path", "remote-git-auth-username",
				"remote-git-branch", "remote-git-commit", "remote-git-repo-dir",
				"stack", "suppress-permalink",
			} {
				a = forwardStringFlag(cmd, a, f)
			}
			for _, f := range []string{
				"clear-pending-creates", "debug", "diff", "exclude-dependents",
				"expect-no-changes", "json", "neo", "preview-only", "remote",
				"remote-inherit-settings", "remote-skip-install-dependencies",
				"run-program", "show-replacement-steps", "show-sames",
				"skip-pending-creates", "skip-preview", "suppress-outputs",
				"suppress-progress", "suppress-stream-logs", "target-dependents", "yes",
			} {
				a = forwardBoolFlag(cmd, a, f)
			}
			a = forwardInt32Flag(cmd, a, "parallel")

			return executeCommand(cmd, a)
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
