package block

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	blockAddCmd.Flags().StringP("context", "c", "", "Context or description for the new blocker")
	blockAddCmd.Flags().StringP("who", "w", "", "Who is blocking?")
}

var blockAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add tags to the database",
	Long: `all tags used are stored along with context which can be used to generate useful reports.
	You must add the task id as a parameter.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Because Args: cobra.ExactArgs(1) is set,
		// we are guaranteed that args[0] exists and contains the single tag name.
		id := args[0]

		// The flag value is retrieved here
		context, _ := cmd.Flags().GetString("context")
		blocker, _ := cmd.Flags().GetString("who")

		fmt.Printf("Adding blocker to task '%s' with Context: '%s'\n", id, context)
		cmdBlock := commands.BlockCmd{ID: id}
		cmdBlock.Add(store.Store, blocker, context)
	},
}
