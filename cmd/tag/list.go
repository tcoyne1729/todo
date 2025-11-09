package tag

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands/tags"
	"github.com/tcoyne1729/todo/store"
)

// func init() {
// 	tagCmd.AddCommand(tagAddCmd)
// }

var TagListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tags",
	Long:  `list all the tags in the database`,
	Run: func(ccmd *cobra.Command, args []string) {
		tagCmd := tags.TagListCmd{}
		tagCmd.Run(store.Store)
	},
}
