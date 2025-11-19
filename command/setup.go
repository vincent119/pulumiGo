package command

import (
	"pulumiGo/command/handlers"

	"github.com/spf13/cobra"
)

// Setup initializes and returns all command handlers
func Setup() *CommandRegistry {
    registry := NewCommandRegistry()

    registry.Register(NewSimpleCommand("up", "Create or update resources in a stack",
        "Update the resources in a stack to match the current configuration."))
    registry.Register(handlers.NewPreviewCommand())
    registry.Register(handlers.NewLogoutHandler())
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
    registry.Register(handlers.NewImportHandler())
    registry.Register(handlers.NewRefreshHandler())


    return registry
}

// AddCommands adds all subcommands to the root command
func AddCommands(rootCmd *cobra.Command) {

    registry := Setup()

    registry.AddToRootCommand(rootCmd)
}
