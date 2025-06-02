package command

import (
    "github.com/spf13/cobra"
    "pulumiGo/command/handlers"
)

// Setup initializes and returns all command handlers
func Setup() *CommandRegistry {
    registry := NewCommandRegistry()
    
    registry.Register(NewSimpleCommand("up", "Create or update resources in a stack",
        "Update the resources in a stack to match the current configuration."))
    registry.Register(handlers.NewPreviewCommand())
    registry.Register(NewSimpleCommand("logout", "Log out of the Pulumi Cloud",
        "Log out of the Pulumi Cloud and remove saved credentials."))
        registry.Register(NewSimpleCommand("whoami", "Display the current logged-in user",
        "Displays the username of the currently logged in user."))
    

    registry.Register(handlers.NewConfigHandler())
    registry.Register(handlers.NewStackHandler())
    registry.Register(handlers.NewLoginHandler())
    registry.Register(handlers.NewPluginCommand())
    registry.Register(handlers.NewStateCommand())
    registry.Register(handlers.NewVersionCommand())
    registry.Register(handlers.NewAboutCommand())
    registry.Register(handlers.NewOrgCommand())

    
    return registry
}

// AddCommands adds all subcommands to the root command
func AddCommands(rootCmd *cobra.Command) {

    registry := Setup()
    
    registry.AddToRootCommand(rootCmd)
}