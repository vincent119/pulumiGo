package handlers

import (
	"fmt"
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

// UpCommand 處理 pulumi up 命令
type UpCommand struct {
    types.BaseHandler
}

// NewUpCommand 創建新的 up 命令處理器
func NewUpCommand() *UpCommand {
    cmd := &cobra.Command{
        Use:   "up",
        Short: "Create or update resources in a stack",
        Long:  `Update the resources in a stack to match the current configuration.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            cmdArgs := []string{"up"}
            cmdArgs = append(cmdArgs, args...)

            // 添加字符串數組標誌
            if cmd.Flag("attach-debugger") != nil && cmd.Flag("attach-debugger").Changed {
                attachDebugger, _ := cmd.Flags().GetStringArray("attach-debugger")
                for _, v := range attachDebugger {
                    cmdArgs = append(cmdArgs, "--attach-debugger", v)
                }
            }
            if cmd.Flag("config") != nil && cmd.Flag("config").Changed {
                config, _ := cmd.Flags().GetStringArray("config")
                for _, v := range config {
                    cmdArgs = append(cmdArgs, "--config", v)
                }
            }
            if cmd.Flag("exclude") != nil && cmd.Flag("exclude").Changed {
                exclude, _ := cmd.Flags().GetStringArray("exclude")
                for _, v := range exclude {
                    cmdArgs = append(cmdArgs, "--exclude", v)
                }
            }
            if cmd.Flag("policy-pack") != nil && cmd.Flag("policy-pack").Changed {
                policyPack, _ := cmd.Flags().GetStringArray("policy-pack")
                for _, v := range policyPack {
                    cmdArgs = append(cmdArgs, "--policy-pack", v)
                }
            }
            if cmd.Flag("policy-pack-config") != nil && cmd.Flag("policy-pack-config").Changed {
                policyPackConfig, _ := cmd.Flags().GetStringArray("policy-pack-config")
                for _, v := range policyPackConfig {
                    cmdArgs = append(cmdArgs, "--policy-pack-config", v)
                }
            }
            if cmd.Flag("remote-env") != nil && cmd.Flag("remote-env").Changed {
                remoteEnv, _ := cmd.Flags().GetStringArray("remote-env")
                for _, v := range remoteEnv {
                    cmdArgs = append(cmdArgs, "--remote-env", v)
                }
            }
            if cmd.Flag("remote-env-secret") != nil && cmd.Flag("remote-env-secret").Changed {
                remoteEnvSecret, _ := cmd.Flags().GetStringArray("remote-env-secret")
                for _, v := range remoteEnvSecret {
                    cmdArgs = append(cmdArgs, "--remote-env-secret", v)
                }
            }
            if cmd.Flag("remote-pre-run-command") != nil && cmd.Flag("remote-pre-run-command").Changed {
                remotePreRun, _ := cmd.Flags().GetStringArray("remote-pre-run-command")
                for _, v := range remotePreRun {
                    cmdArgs = append(cmdArgs, "--remote-pre-run-command", v)
                }
            }
            if cmd.Flag("replace") != nil && cmd.Flag("replace").Changed {
                replace, _ := cmd.Flags().GetStringArray("replace")
                for _, v := range replace {
                    cmdArgs = append(cmdArgs, "--replace", v)
                }
            }
            if cmd.Flag("target") != nil && cmd.Flag("target").Changed {
                target, _ := cmd.Flags().GetStringArray("target")
                for _, v := range target {
                    cmdArgs = append(cmdArgs, "--target", v)
                }
            }
            if cmd.Flag("target-replace") != nil && cmd.Flag("target-replace").Changed {
                targetReplace, _ := cmd.Flags().GetStringArray("target-replace")
                for _, v := range targetReplace {
                    cmdArgs = append(cmdArgs, "--target-replace", v)
                }
            }

            // 添加字符串標誌
            if cmd.Flag("config-file") != nil && cmd.Flag("config-file").Changed {
                configFile, _ := cmd.Flags().GetString("config-file")
                cmdArgs = append(cmdArgs, "--config-file", configFile)
            }
            if cmd.Flag("message") != nil && cmd.Flag("message").Changed {
                message, _ := cmd.Flags().GetString("message")
                cmdArgs = append(cmdArgs, "--message", message)
            }
            if cmd.Flag("plan") != nil && cmd.Flag("plan").Changed {
                plan, _ := cmd.Flags().GetString("plan")
                cmdArgs = append(cmdArgs, "--plan", plan)
            }
            if cmd.Flag("refresh") != nil && cmd.Flag("refresh").Changed {
                refresh, _ := cmd.Flags().GetString("refresh")
                cmdArgs = append(cmdArgs, "--refresh", refresh)
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
            if cmd.Flag("secrets-provider") != nil && cmd.Flag("secrets-provider").Changed {
                secretsProvider, _ := cmd.Flags().GetString("secrets-provider")
                cmdArgs = append(cmdArgs, "--secrets-provider", secretsProvider)
            }
            if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
                stack, _ := cmd.Flags().GetString("stack")
                cmdArgs = append(cmdArgs, "--stack", stack)
            }
            if cmd.Flag("suppress-permalink") != nil && cmd.Flag("suppress-permalink").Changed {
                suppressPermalink, _ := cmd.Flags().GetString("suppress-permalink")
                cmdArgs = append(cmdArgs, "--suppress-permalink", suppressPermalink)
            }

            // 添加布爾標誌
            if cmd.Flag("config-path") != nil && cmd.Flag("config-path").Changed {
                cmdArgs = append(cmdArgs, "--config-path")
            }
            if cmd.Flag("continue-on-error") != nil && cmd.Flag("continue-on-error").Changed {
                cmdArgs = append(cmdArgs, "--continue-on-error")
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
            if cmd.Flag("show-config") != nil && cmd.Flag("show-config").Changed {
                cmdArgs = append(cmdArgs, "--show-config")
            }
            if cmd.Flag("show-full-output") != nil && cmd.Flag("show-full-output").Changed {
                cmdArgs = append(cmdArgs, "--show-full-output")
            }
            if cmd.Flag("show-policy-remediations") != nil && cmd.Flag("show-policy-remediations").Changed {
                cmdArgs = append(cmdArgs, "--show-policy-remediations")
            }
            if cmd.Flag("show-reads") != nil && cmd.Flag("show-reads").Changed {
                cmdArgs = append(cmdArgs, "--show-reads")
            }
            if cmd.Flag("show-replacement-steps") != nil && cmd.Flag("show-replacement-steps").Changed {
                cmdArgs = append(cmdArgs, "--show-replacement-steps")
            }
            if cmd.Flag("show-sames") != nil && cmd.Flag("show-sames").Changed {
                cmdArgs = append(cmdArgs, "--show-sames")
            }
            if cmd.Flag("show-secrets") != nil && cmd.Flag("show-secrets").Changed {
                cmdArgs = append(cmdArgs, "--show-secrets")
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

            // 添加 int32 標誌
            if cmd.Flag("parallel") != nil && cmd.Flag("parallel").Changed {
                parallel, _ := cmd.Flags().GetInt32("parallel")
                cmdArgs = append(cmdArgs, "--parallel", fmt.Sprintf("%d", parallel))
            }

            return executeCommand(cmd, cmdArgs)
        },
    }

    // String array flags
    cmd.Flags().StringArray("attach-debugger", nil, "Enable the ability to attach a debugger to the program and source based plugins being executed")
    cmd.Flags().StringArrayP("config", "c", nil, "Config to use during the update and save to the stack config file")
    cmd.Flags().StringArray("exclude", nil, "Specify a resource URN to ignore")
    cmd.Flags().StringArray("policy-pack", nil, "Run one or more policy packs as part of this update")
    cmd.Flags().StringArray("policy-pack-config", nil, "Path to JSON file containing the config for the policy pack")
    cmd.Flags().StringArray("remote-env", nil, "[EXPERIMENTAL] Environment variables to use in the remote operation")
    cmd.Flags().StringArray("remote-env-secret", nil, "[EXPERIMENTAL] Environment variables with secret values")
    cmd.Flags().StringArray("remote-pre-run-command", nil, "[EXPERIMENTAL] Commands to run before the remote operation")
    cmd.Flags().StringArray("replace", nil, "Specify a single resource URN to replace")
    cmd.Flags().StringArrayP("target", "t", nil, "Specify a single resource URN to update")
    cmd.Flags().StringArray("target-replace", nil, "Specify a single resource URN to replace")

    // String flags
    cmd.Flags().String("config-file", "", "Use the configuration values in the specified file")
    cmd.Flags().StringP("message", "m", "", "Optional message to associate with the update operation")
    cmd.Flags().String("plan", "", "[EXPERIMENTAL] Path to a plan file to use for the update")
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
    cmd.Flags().String("secrets-provider", "default", "The type of the provider that should be used to encrypt and decrypt secrets")
    cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on")
    cmd.Flags().String("suppress-permalink", "", "Suppress display of the state permalink")

    // Boolean flags
    cmd.Flags().Bool("config-path", false, "Config keys contain a path to a property in a map or list to set")
    cmd.Flags().Bool("continue-on-error", false, "Continue updating resources even if an error is encountered")
    cmd.Flags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
    cmd.Flags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
    cmd.Flags().Bool("exclude-dependents", false, "Allows ignoring of dependent targets discovered but not specified in --exclude list")
    cmd.Flags().Bool("expect-no-changes", false, "Return an error if any changes occur during this update")
    cmd.Flags().BoolP("json", "j", false, "Serialize the update diffs, operations, and overall output as JSON")
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
    cmd.Flags().Bool("show-sames", false, "Show resources that don't need be updated because they haven't changed")
    cmd.Flags().Bool("show-secrets", false, "Show secret outputs in the CLI output")
    cmd.Flags().BoolP("skip-preview", "f", false, "Do not calculate a preview before performing the update")
    cmd.Flags().Bool("suppress-outputs", false, "Suppress display of stack outputs")
    cmd.Flags().Bool("suppress-progress", false, "Suppress display of periodic progress dots")
    cmd.Flags().Bool("suppress-stream-logs", false, "[EXPERIMENTAL] Suppress log streaming of the deployment job")
    cmd.Flags().Bool("target-dependents", false, "Allows updating of dependent targets discovered but not specified in --target list")
    cmd.Flags().BoolP("yes", "y", false, "Automatically approve and perform the update after previewing it")

    // Int32 flag
    cmd.Flags().Int32P("parallel", "p", 16, "Allow P resource operations to run in parallel at once")

    return &UpCommand{
        BaseHandler: types.BaseHandler{Command: cmd},
    }
}
