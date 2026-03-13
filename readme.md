# pulumiGo

Pulumi Go Wrapper - Manage infrastructure with Pulumi.

`pulumiGo` is a tool that enhances Pulumi with additional capabilities. It allows you to manage your infrastructure as code using Pulumi commands while providing extra features for stack management.

## Features

- **Wrapper for Pulumi**: Seamlessly executes Pulumi commands.
- **Debug Mode**: Enable debug logging with the `--debug` or `-d` flag.
- **Stack Management**: Enhanced commands for stack operations.

## Supported Commands

- `up`: Create or update resources in a stack.
- `preview`: Preview updates to a stack's resources.
- `stack`: Manage stacks.
- `config`: Manage configuration.
- `login` / `logout`: Manage user authentication.
- `whoami`: Display the current logged-in user.
- `refresh`: Refresh the resources in a stack.
- `import`: Import resources into an existing stack.
- `state`: Edit the current stack's state.
- `plugin`: Manage language and resource plugins.
- `org`: Manage Organization.
- `about`: Show information about the Pulumi CLI.
- `version`: Show version information.

## Installation

### Homebrew (macOS / Linux)

```zsh
brew tap vincent119/tap
brew install pulumiGo
```

## Build Instructions

### Build for Mac

```zsh
go build -o pulumiGo
```

### Build for Windows

```zsh
GOOS=windows GOARCH=amd64 go build -o pulumiGo.exe
```

### MacOS Quarantine Handling

If you encounter permission issues on macOS, you may need to remove the quarantine attribute.

Check attribute:

```zsh
xattr pulumiGo
# Output: com.apple.quarantine
```

Remove quarantine:

```zsh
xattr -d com.apple.quarantine pulumiGo
```

## Shell Completion (Zsh)

To enable zsh completion for `pulumiGo`:

```zsh
# Create the completion directory if it doesn't exist
mkdir -p ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/pulumiGo

# Generate the completion script
pulumiGo completion zsh > ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/pulumiGo/_pulumiGo

# Add 'pulumiGo' to your plugins list in ~/.zshrc
# plugins=(git ... pulumiGo)

# Reload your shell
source ~/.zshrc
```
