package block

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

var blockCompleteCmd = &cobra.Command{
	Use:   "close",
	Short: "mark the current blocker as closed",
	Long:  `you must provide the id of the blocker that you want to close as an argument.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := store.Store
		if store.Current == "" {
			return errors.New("no curent todo task")
		}
		blockId := args[0]
		cmdBlock := commands.BlockCmd{ID: store.Current}
		cmdBlock.Complete(store, blockId)
		return nil
	},
}
