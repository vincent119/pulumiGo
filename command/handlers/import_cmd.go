package handlers

import (
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
			a := append([]string{"import"}, args...)

			for _, f := range []string{"config", "properties"} {
				a = forwardStringArrayFlag(cmd, a, f)
			}
			for _, f := range []string{
				"config-file", "file", "from", "generate-resources",
				"message", "out", "parent", "provider", "stack", "suppress-permalink",
			} {
				a = forwardStringFlag(cmd, a, f)
			}
			for _, f := range []string{
				"debug", "diff", "generate-code", "json", "preview-only",
				"protect", "skip-preview", "suppress-outputs", "suppress-progress", "yes",
			} {
				a = forwardBoolFlag(cmd, a, f)
			}
			a = forwardInt32Flag(cmd, a, "parallel")

			return executeCommand(cmd, a)
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
