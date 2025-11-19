package handlers

import (
	"fmt"
	"pulumiGo/types"

	"github.com/spf13/cobra"
)

type ImportHandler struct {
	types.BaseHandler
}

func NewImportHandler() *ImportHandler {
	cmd := &cobra.Command{
		Use:   "import [type] [name] [id]",
		Short: "Import resources into an existing stack",
		Long:  `Import resources into an existing stack.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmdArgs := []string{"import"}
			cmdArgs = append(cmdArgs, args...)

			// String array flags
			if cmd.Flag("config") != nil && cmd.Flag("config").Changed {
				config, _ := cmd.Flags().GetStringArray("config")
				for _, c := range config {
					cmdArgs = append(cmdArgs, "--config", c)
				}
			}
			if cmd.Flag("properties") != nil && cmd.Flag("properties").Changed {
				properties, _ := cmd.Flags().GetStringArray("properties")
				for _, p := range properties {
					cmdArgs = append(cmdArgs, "--properties", p)
				}
			}

			// String flags
			if cmd.Flag("config-file") != nil && cmd.Flag("config-file").Changed {
				configFile, _ := cmd.Flags().GetString("config-file")
				cmdArgs = append(cmdArgs, "--config-file", configFile)
			}
			if cmd.Flag("file") != nil && cmd.Flag("file").Changed {
				file, _ := cmd.Flags().GetString("file")
				cmdArgs = append(cmdArgs, "--file", file)
			}
			if cmd.Flag("from") != nil && cmd.Flag("from").Changed {
				from, _ := cmd.Flags().GetString("from")
				cmdArgs = append(cmdArgs, "--from", from)
			}
			if cmd.Flag("generate-resources") != nil && cmd.Flag("generate-resources").Changed {
				generateResources, _ := cmd.Flags().GetString("generate-resources")
				cmdArgs = append(cmdArgs, "--generate-resources", generateResources)
			}
			if cmd.Flag("message") != nil && cmd.Flag("message").Changed {
				message, _ := cmd.Flags().GetString("message")
				cmdArgs = append(cmdArgs, "--message", message)
			}
			if cmd.Flag("out") != nil && cmd.Flag("out").Changed {
				out, _ := cmd.Flags().GetString("out")
				cmdArgs = append(cmdArgs, "--out", out)
			}
			if cmd.Flag("parent") != nil && cmd.Flag("parent").Changed {
				parent, _ := cmd.Flags().GetString("parent")
				cmdArgs = append(cmdArgs, "--parent", parent)
			}
			if cmd.Flag("provider") != nil && cmd.Flag("provider").Changed {
				provider, _ := cmd.Flags().GetString("provider")
				cmdArgs = append(cmdArgs, "--provider", provider)
			}
			if cmd.Flag("stack") != nil && cmd.Flag("stack").Changed {
				stack, _ := cmd.Flags().GetString("stack")
				cmdArgs = append(cmdArgs, "--stack", stack)
			}
			if cmd.Flag("suppress-permalink") != nil && cmd.Flag("suppress-permalink").Changed {
				suppressPermalink, _ := cmd.Flags().GetString("suppress-permalink")
				cmdArgs = append(cmdArgs, "--suppress-permalink", suppressPermalink)
			}

			// Boolean flags
			if cmd.Flag("debug") != nil && cmd.Flag("debug").Changed {
				cmdArgs = append(cmdArgs, "--debug")
			}
			if cmd.Flag("diff") != nil && cmd.Flag("diff").Changed {
				cmdArgs = append(cmdArgs, "--diff")
			}
			if cmd.Flag("generate-code") != nil && cmd.Flag("generate-code").Changed {
				cmdArgs = append(cmdArgs, "--generate-code")
			}
			if cmd.Flag("json") != nil && cmd.Flag("json").Changed {
				cmdArgs = append(cmdArgs, "--json")
			}
			if cmd.Flag("preview-only") != nil && cmd.Flag("preview-only").Changed {
				cmdArgs = append(cmdArgs, "--preview-only")
			}
			if cmd.Flag("protect") != nil && cmd.Flag("protect").Changed {
				cmdArgs = append(cmdArgs, "--protect")
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
			if cmd.Flag("yes") != nil && cmd.Flag("yes").Changed {
				cmdArgs = append(cmdArgs, "--yes")
			}

			// Int32 flag
			if cmd.Flag("parallel") != nil && cmd.Flag("parallel").Changed {
				parallel, _ := cmd.Flags().GetInt32("parallel")
				cmdArgs = append(cmdArgs, "--parallel", fmt.Sprintf("%d", parallel))
			}

			return executeCommand(cmd, cmdArgs)
		},
	}

	// String array flags
	cmd.Flags().StringArray("config", nil, "Config to use during the import")
	cmd.Flags().StringArray("properties", nil, "The property names to use for the import in the format name1,name2")

	// String flags
	cmd.Flags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	cmd.Flags().StringP("file", "f", "", "The path to a JSON-encoded file containing a list of resources to import")
	cmd.Flags().String("from", "", "Invoke a converter to import the resources")
	cmd.Flags().String("generate-resources", "", "When used with --from, always write a JSON-encoded file containing a list of importable resources discovered by conversion to the specified path")
	cmd.Flags().StringP("message", "m", "", "Optional message to associate with the update operation")
	cmd.Flags().StringP("out", "o", "", "The path to the file that will contain the generated resource declarations")
	cmd.Flags().String("parent", "", "The name and URN of the parent resource in the format name=urn, where name is the variable name of the parent resource")
	cmd.Flags().String("provider", "", "The name and URN of the provider to use for the import in the format name=urn, where name is the variable name for the provider resource")
	cmd.Flags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	cmd.Flags().String("suppress-permalink", "", "Suppress display of the state permalink")

	// Boolean flags
	cmd.Flags().BoolP("debug", "d", false, "Print detailed debugging output during resource operations")
	cmd.Flags().Bool("diff", false, "Display operation as a rich diff showing the overall change")
	cmd.Flags().Bool("generate-code", true, "Generate resource declaration code for the imported resources")
	cmd.Flags().BoolP("json", "j", false, "Serialize the import diffs, operations, and overall output as JSON")
	cmd.Flags().Bool("preview-only", false, "Only show a preview of the import, but don't perform the import itself")
	cmd.Flags().Bool("protect", true, "Allow resources to be imported with protection from deletion enabled")
	cmd.Flags().Bool("skip-preview", false, "Do not calculate a preview before performing the import")
	cmd.Flags().Bool("suppress-outputs", false, "Suppress display of stack outputs (in case they contain sensitive values)")
	cmd.Flags().Bool("suppress-progress", false, "Suppress display of periodic progress dots")
	cmd.Flags().BoolP("yes", "y", false, "Automatically approve and perform the import after previewing it")

	// Int32 flag
	cmd.Flags().Int32P("parallel", "p", 16, "Allow P resource operations to run in parallel at once (1 for no parallelism).")

	return &ImportHandler{BaseHandler: types.BaseHandler{Command: cmd}}
}

func (h *ImportHandler) RegisterSubcommands(cmd *cobra.Command) {
	// import 命令目前沒有子命令
}
