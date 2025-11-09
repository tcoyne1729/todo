package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands/tags"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	TagAddCmd.Flags().StringP("context", "c", "", "Context or description for the new tag (e.g., department name)")
}

var TagAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add tags to the database",
	Long:  `all tags used are stored along with context which can be used to generate useful reports`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Because Args: cobra.ExactArgs(1) is set,
		// we are guaranteed that args[0] exists and contains the single tag name.
		tagName := args[0]

		// The flag value is retrieved here
		context, _ := cmd.Flags().GetString("context")

		fmt.Printf("Adding Tag: '%s' with Context: '%s'\n", tagName, context)
		cmdAdd := tags.TagAddCmd{
			Name:    tagName,
			Context: context,
		}
		cmdAdd.Run(store.Store)
	},
}
