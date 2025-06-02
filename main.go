package main

import (
    "fmt"
    "os"
    "pulumiGo/command"
    "pulumiGo/command/handlers"
    "pulumiGo/iac_modules"
    "github.com/spf13/cobra"
)

var (
    debugMode bool
)

func main() {
    handlers.InitExecuteFunc(command.ExecuteCmd)
    rootCmd := &cobra.Command{
        Use:   "pulumiGo [command]",
        Short: "Pulumi Go Wrapper - Manage infrastructure with Pulumi",
        Long: `Pulumi Go Wrapper is a tool that enhances Pulumi with additional capabilities.
It allows you to manage your infrastructure as code using Pulumi commands
while providing extra features for stack management.`,
        PersistentPreRun: func(cmd *cobra.Command, args []string) {
            if debugMode {
                fmt.Println("Debug mode enabled")
                iac_modules.SetDebugMode(true)
            }
        },
        RunE: func(cmd *cobra.Command, args []string) error {
            // 空情況下顯示幫助訊息
            if len(args) == 0 {
                return cmd.Help()
            }
            return command.ExecuteCmd(cmd, args)
        },
    }
    
    // Add global flags
    rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "Enable debug logging")
        
    // Add all commands to the root command
    command.AddCommands(rootCmd)
    
    // Execute the command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    // Set debug mode after command parsing
    if debugMode {
        iac_modules.SetDebugMode(true)
    }
}