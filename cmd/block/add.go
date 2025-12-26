package block

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	blockAddCmd.Flags().StringP("context", "c", "", "Context or description for the new blocker")
	blockAddCmd.Flags().StringP("who", "w", "", "Who is blocking?")
	blockAddCmd.Flags().StringP("todo", "t", "", "todo id. if blank then current todo task.")
}

var blockAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add blocker to a todo",
	Long: `add blockers. the default is to add to the active task
	but if you want to add blockers to another task then you
	need to provide the todo id.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// The flag value is retrieved here
		context, _ := cmd.Flags().GetString("context")
		blockedBy, _ := cmd.Flags().GetString("who")
		todoId, _ := cmd.Flags().GetString("todo")

		localStore := store.Store
		currentId := localStore.Current

		if currentId == "" && todoId == "" {
			return errors.New("no current todo and no todo id provided.")
		}
		var id string
		if todoId == "" && currentId != "" {
			id = currentId
		} else {
			id = todoId
		}

		fmt.Printf("Adding blocker to task '%s' with Context: '%s'\n", id, context)
		cmdBlock := commands.BlockCmd{ID: id}
		cmdBlock.Add(store.Store, blockedBy, context)
		return nil
	},
}
