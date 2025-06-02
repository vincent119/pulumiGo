package command

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "pulumiGo/iac_modules"
    "github.com/spf13/cobra"
)



func init() {
    // 初始化默認執行器
    DefaultExecutor = &pulumiBinaryExecutor{}
}

// CommandExecutor 介面定義
type CommandExecutor interface {
    Execute(cmd *cobra.Command, args []string) error
    Run(commands []string, stackName string)
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
        stackName, err = iac_modules.StackCheck()
        if err != nil {
            return fmt.Errorf("stack check error: %v", err)
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
    iac_modules.DebugLog("Executing command: %v", pulumiArgs)
    
    // 對於 stackless 命令，直接執行並跳過 Join/Recovery
    if isStacklessCommand {
        cmdExec := exec.Command(pulumiArgs[0], pulumiArgs[1:]...)
        cmdExec.Stdout = os.Stdout
        cmdExec.Stderr = os.Stderr
        cmdExec.Stdin = os.Stdin
        
        return cmdExec.Run()
    }
    
    // 其他命令正常執行
    e.Run(pulumiArgs, stackName)
    
    return nil
}

// Run executes a Pulumi command with the given arguments and stack name
func (e *pulumiBinaryExecutor) Run(commands []string, stackName string) {
    err := iac_modules.Join(stackName)
    if err != nil {
        log.Println("Join error:", err)
        return
    }

    iac_modules.DebugLog("Executing command: %v", commands)

    // Execute command with standard input support for interactive mode
    cmd := exec.Command(commands[0], commands[1:]...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    err = cmd.Run()
    if err != nil {
        log.Println("Command error:", err)
    }

    // Always execute recovery
    err = iac_modules.Recovery()
    if err != nil {
        log.Println("Recovery error:", err)
    }
}
