package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"pulumiGo/iac"

	"github.com/spf13/cobra"
)

func init() {
	DefaultExecutor = &pulumiBinaryExecutor{}
}

// pulumiBinaryExecutor 實現了 CommandExecutor 介面
type pulumiBinaryExecutor struct{}

// Execute 執行 Pulumi 命令
func (e *pulumiBinaryExecutor) Execute(cmd *cobra.Command, args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("no Pulumi command specified")
    }

    // 判斷是否是不需要堆疊的命令
    isStacklessCommand := false
    stacklessCommands := []string{"login", "logout", "version", "whoami"}
    isStackSelectCommand := len(args) >= 2 && args[0] == "stack" && args[1] == "select"

    if isStackSelectCommand {
        isStacklessCommand = true
    } else if len(args) > 0 {
        for _, cmdName := range stacklessCommands {
            if args[0] == cmdName {
                isStacklessCommand = true
                break
            }
        }
    }

    // 只有非特殊命令才需要獲取堆疊
    var stackName string
    var err error
    if !isStacklessCommand {
        stackName, err = iac.StackCheck()
        if err != nil {
            return fmt.Errorf("stack check error: %w", err)
        }
    }

    // Construct Pulumi command
    pulumiArgs := []string{"pulumi"}
    pulumiArgs = append(pulumiArgs, args...)

    // 檢查並添加標誌
    // 特別處理 config 命令的標誌
    if len(args) > 0 && args[0] == "config" {
        // 處理 --show-secrets 標誌
        if cmd.Flag("show-secrets") != nil && cmd.Flag("show-secrets").Changed {
            pulumiArgs = append(pulumiArgs, "--show-secrets")
        }

        // 處理 --json 標誌
        if cmd.Flag("json") != nil && cmd.Flag("json").Changed {
            pulumiArgs = append(pulumiArgs, "--json")
        }
    }

    // 處理 login 命令的 --local 標誌
    if len(args) > 0 && args[0] == "login" {
        if cmd.Flag("local") != nil && cmd.Flag("local").Changed {
            pulumiArgs = append(pulumiArgs, "--local")
        }
    }

    // 僅在 debug 模式下輸出日誌
    iac.DebugLog("Executing command: %v", pulumiArgs)

    // 對於 stackless 命令,直接執行並跳過 Join/Recovery
    if isStacklessCommand {
        cmdExec := exec.Command(pulumiArgs[0], pulumiArgs[1:]...)
        cmdExec.Stdout = os.Stdout
        cmdExec.Stderr = os.Stderr
        cmdExec.Stdin = os.Stdin

        return cmdExec.Run()
    }

    // 其他命令正常執行
    return e.Run(pulumiArgs, stackName)
}

// Run executes a Pulumi command with the given arguments and stack name.
// Recovery is always attempted even if the command fails.
func (e *pulumiBinaryExecutor) Run(commands []string, stackName string) error {
    if err := iac.Join(stackName); err != nil {
        return fmt.Errorf("join error: %w", err)
    }

    iac.DebugLog("Executing command: %v", commands)

    // Execute command with standard input support for interactive mode.
    cmd := exec.Command(commands[0], commands[1:]...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    cmdErr := cmd.Run()
    if cmdErr != nil {
        log.Println("Command error:", cmdErr)
    }

    // Always execute recovery.
    if err := iac.Recovery(); err != nil {
        log.Println("Recovery error:", err)
    }

    return cmdErr
}
