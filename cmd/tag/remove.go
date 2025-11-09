package tag

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands/tags"
	"github.com/tcoyne1729/todo/store"
)

// func init() {
// 	tagCmd.AddCommand(tagAddCmd)
// }

var TagRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove tags from the database",
	Long:  `all tags used are stored along with context which can be used to generate useful reports`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("no tags given to delete")
			return
		}
		for _, tag := range args {
			cmdDel := tags.TagDeleteCmd{Name: tag}
			cmdDel.Run(store.Store)
		}
	},
}
