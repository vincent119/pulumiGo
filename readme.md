# pulumiGo

[![GitHub](https://img.shields.io/badge/github-vincent119%2FpulumiGo-blue?logo=github)](https://github.com/vincent119/pulumiGo)
[![License](https://img.shields.io/github/license/vincent119/pulumiGo)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/vincent119/pulumiGo?logo=go)](go.mod)
[![Build](https://img.shields.io/github/actions/workflow/status/vincent119/pulumiGo/release.yaml?label=build)](https://github.com/vincent119/pulumiGo/actions/workflows/release.yaml)
[![Release](https://img.shields.io/github/v/release/vincent119/pulumiGo)](https://github.com/vincent119/pulumiGo/releases/latest)
[![Stars](https://img.shields.io/github/stars/vincent119/pulumiGo)](https://github.com/vincent119/pulumiGo/stargazers)

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

## Project Structure

`pulumiGo` expects a specific directory layout to manage multi-stack infrastructure. See [`example/iac_project`](example/iac_project) for a working reference.

```
<project>/
├── Pulumi.yaml                          # Project definition (name, runtime)
├── Pulumi.<stack>.yaml                  # Stack config (tags, secrets, env vars)
├── stack_share_variables/               # Global variables shared across all stacks
│   └── aws.yaml                         # e.g. AWS account ID, shared constants
└── stacks/
    └── <stack>/                         # Per-stack resource definitions
        └── aws/
            ├── providers/
            │   └── aws.yaml             # AWS provider definitions (region, account)
            ├── private_variables/
            │   └── variables.yaml       # Stack references and dynamic variables
            └── <service>/
                └── <resource>/
                    └── *.yaml           # Resource definitions (S3, EC2, RDS, etc.)
```

### How it works

When you run `pulumiGo up` or `pulumiGo preview`, it:

1. Reads the current stack name via `pulumi stack ls`
2. Merges `resources`, `variables`, and `outputs` from `stacks/<stack>/` into `Pulumi.yaml`
3. Executes the Pulumi command
4. Restores `Pulumi.yaml` to its original state (recovery)

This allows you to split large infrastructure definitions across multiple YAML files organized by service, while keeping `Pulumi.yaml` clean.

### The `stacks/` directory

The `stacks/` directory is the **required entry point** for all resource definitions. `pulumiGo` will not function without it.

Each subdirectory under `stacks/<stack>/` is recursively scanned for YAML files. All `resources`, `variables`, and `outputs` blocks found are merged into `Pulumi.yaml` before execution.

```
stacks/
└── dev/                        # Must match the active Pulumi stack name
    └── aws/
        ├── providers/
        │   └── aws.yaml        # Provider configuration
        ├── private_variables/
        │   └── variables.yaml  # StackReferences and dynamic lookups
        └── S3/
            └── TestBucket/
                └── bucket.yaml # Resource definition
```

> The stack subdirectory name (e.g. `dev`) must exactly match the active stack selected via `pulumi stack select`.

> For best practices on organizing Pulumi projects and stacks, see the [Pulumi documentation](https://www.pulumi.com/docs/iac/guides/basics/organizing-projects-stacks/).
