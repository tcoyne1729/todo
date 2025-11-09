package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands/tags"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	TagEditCmd.Flags().StringP("context", "c", "", "Updated context or description for the tag (e.g., department name)")
	TagEditCmd.Flags().StringP("name", "n", "", "Updated name for the tag")
}

var TagEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit single tag in the database",
	Long:  `where you want to change the name or context of a tag use this command. Where you want to leave either the name or context unchanged, you can leave it blank`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Because Args: cobra.ExactArgs(1) is set,
		// we are guaranteed that args[0] exists and contains the single tag name.
		tagOldName := args[0]

		// The flag value is retrieved here
		name, _ := cmd.Flags().GetString("name")
		context, _ := cmd.Flags().GetString("context")

		fmt.Printf("Updating Tag: '%s'\n", tagOldName)
		cmdAdd := tags.TagEditCmd{
			OldName:    tagOldName,
			NewName:    name,
			NewContext: context,
		}
		cmdAdd.Run(store.Store)
	},
}
