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
            a := append([]string{"preview"}, args...)

            for _, f := range []string{
                "attach-debugger", "config", "exclude", "policy-pack",
                "policy-pack-config", "remote-env", "remote-env-secret",
                "remote-pre-run-command", "replace", "target", "target-replace",
            } {
                a = forwardStringArrayFlag(cmd, a, f)
            }
            for _, f := range []string{
                "config-file", "import-file", "message", "refresh",
                "remote-agent-pool-id", "remote-executor-image",
                "remote-executor-image-password", "remote-executor-image-username",
                "remote-git-auth-access-token", "remote-git-auth-password",
                "remote-git-auth-ssh-private-key", "remote-git-auth-ssh-private-key-path",
                "remote-git-auth-username", "remote-git-branch", "remote-git-commit",
                "remote-git-repo-dir", "save-plan", "stack", "suppress-permalink",
            } {
                a = forwardStringFlag(cmd, a, f)
            }
            for _, f := range []string{
                "config-path", "debug", "diff", "exclude-dependents",
                "expect-no-changes", "json", "neo", "remote",
                "remote-inherit-settings", "remote-skip-install-dependencies", "run-program",
                "show-config", "show-full-output", "show-policy-remediations", "show-reads",
                "show-replacement-steps", "show-sames", "show-secrets",
                "suppress-outputs", "suppress-progress", "suppress-stream-logs", "target-dependents",
            } {
                a = forwardBoolFlag(cmd, a, f)
            }
            a = forwardInt32Flag(cmd, a, "parallel")

            return executeCommand(cmd, a)
        },
    }

    // String array flags
    cmd.Flags().StringArray("attach-debugger", nil, "Enable the ability to attach a debugger to the program and source based plugins being executed")
    cmd.Flags().StringArrayP("config", "c", nil, "Config to use during the preview and save to the stack config file")
    cmd.Flags().StringArrayP("exclude", "x", nil, "Specify a resource URN to ignore")
    cmd.Flags().StringArray("policy-pack", nil, "Run one or more policy packs as part of this update")
    cmd.Flags().StringArray("policy-pack-config", nil, "Path to JSON file containing the config for the policy pack")
    cmd.Flags().StringArray("remote-env", nil, "[EXPERIMENTAL] Environment variables to use in the remote operation")
    cmd.Flags().StringArray("remote-env-secret", nil, "[EXPERIMENTAL] Environment variables with secret values")
    cmd.Flags().StringArray("remote-pre-run-command", nil, "[EXPERIMENTAL] Commands to run before the remote operation")
    cmd.Flags().StringArray("replace", nil, "Specify resources to replace")
    cmd.Flags().StringArrayP("target", "t", nil, "Specify a single resource URN to update")
    cmd.Flags().StringArray("target-replace", nil, "Specify a single resource URN to replace")

    // String flags
    cmd.Flags().String("config-file", "", "Use the configuration values in the specified file")
    cmd.Flags().String("import-file", "", "Save any creates seen during the preview into an import file")
    cmd.Flags().StringP("message", "m", "", "Optional message to associate with the preview operation")
    cmd.Flags().StringP("refresh", "r", "", "Refresh the state of the stack's resources before this update")
    cmd.Flags().String("remote-agent-pool-id", "", "[EXPERIMENTAL] The agent pool to use to run the deployment job")
    cmd.Flags().String("remote-executor-image", "", "[EXPERIMENTAL] The Docker image to use for the executor")
    cmd.Flags().String("remote-executor-image-password", "", "[EXPERIMENTAL] The password for the credentials")
    cmd.Flags().String("remote-executor-image-username", "", "[EXPERIMENTAL] The username for the credentials")
    cmd.Flags().String("remote-git-auth-access-token", "", "[EXPERIMENTAL] Git personal access token")
    cmd.Flags().String("remote-git-auth-password", "", "[EXPERIMENTAL] Git password")
    cmd.Flags().String("remote-git-auth-ssh-private-key", "", "[EXPERIMENTAL] Git SSH private key")
    cmd.Flags().String("remote-git-auth-ssh-private-key-path", "", "[EXPERIMENTAL] Git SSH private key path")
    cmd.Flags().String("remote-git-auth-username", "", "[EXPERIMENTAL] Git username")
    cmd.Flags().String("remote-git-branch", "", "[EXPERIMENTAL] Git branch to deploy")
    cmd.Flags().String("remote-git-commit", "", "[EXPERIMENTAL] Git commit hash of the commit to deploy")
    cmd.Flags().String("remote-git-repo-dir", "", "[EXPERIMENTAL] The directory to work from in the project's source repository")
    cmd.Flags().String("save-plan", "", "[PREVIEW] Save the operations proposed by the preview to a plan file")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on")
    cmd.Flags().String("suppress-permalink", "", "Suppress display of the state permalink")

    // Boolean flags
    cmd.Flags().Bool("config-path", false, "Config keys contain a path to a property in a map or list to set")
    cmd.Flags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
    cmd.Flags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
    cmd.Flags().Bool("exclude-dependents", false, "Allow ignoring of dependent targets discovered but not specified in --exclude list")
    cmd.Flags().Bool("expect-no-changes", false, "Return an error if any changes are proposed by this preview")
    cmd.Flags().BoolP("json", "j", false, "Serialize the preview diffs, operations, and overall output as JSON")
    cmd.Flags().Bool("neo", false, "Enable Pulumi Neo's assistance for improved CLI experience and insights")
    cmd.Flags().Bool("remote", false, "[EXPERIMENTAL] Run the operation remotely")
    cmd.Flags().Bool("remote-inherit-settings", false, "[EXPERIMENTAL] Inherit deployment settings from the current stack")
    cmd.Flags().Bool("remote-skip-install-dependencies", false, "[EXPERIMENTAL] Whether to skip the default dependency installation step")
    cmd.Flags().Bool("run-program", false, "Run the program to determine up-to-date state for providers to refresh resources")
    cmd.Flags().Bool("show-config", false, "Show configuration keys and variables")
    cmd.Flags().Bool("show-full-output", false, "Display full length of inputs & outputs")
    cmd.Flags().Bool("show-policy-remediations", false, "Show per-resource policy remediation details instead of a summary")
    cmd.Flags().Bool("show-reads", false, "Show resources that are being read in")
    cmd.Flags().Bool("show-replacement-steps", false, "Show detailed resource replacement creates and deletes")
    cmd.Flags().Bool("show-sames", false, "Show resources that needn't be updated because they haven't changed")
    cmd.Flags().Bool("show-secrets", false, "Show secrets in plaintext in the CLI output")
    cmd.Flags().Bool("suppress-outputs", false, "Suppress display of stack outputs")
    cmd.Flags().Bool("suppress-progress", false, "Suppress display of periodic progress dots")
    cmd.Flags().Bool("suppress-stream-logs", false, "[EXPERIMENTAL] Suppress log streaming of the deployment job")
    cmd.Flags().Bool("target-dependents", false, "Allow updating of dependent targets discovered but not specified in --target list")

    // Int32 flag
    cmd.Flags().Int32P("parallel", "p", 16, "Allow P resource operations to run in parallel at once")

    return &PreviewCommand{
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}
