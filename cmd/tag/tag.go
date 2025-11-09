package tag

import (
	"github.com/spf13/cobra"
)

func init() {
	TagCmd.AddCommand(TagAddCmd)
	TagCmd.AddCommand(TagListCmd)
	TagCmd.AddCommand(TagRemoveCmd)
	TagCmd.AddCommand(TagEditCmd)
}

var TagCmd = &cobra.Command{
	Use:   "tag",
	Short: "manipulate the tags on a task",
	Long:  `tag are used to add metadata to a task which can be reused across many tasks (this could be for example the name of the person or department the work is for).`,
}
