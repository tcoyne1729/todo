package block

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	blockNoteCmd.Flags().StringP("note", "n", "", "note to add to the blocker")
}

var blockNoteCmd = &cobra.Command{
	Use:   "note",
	Short: "add a note to a blocker",
	Long:  `you must provide the id of the blocker that you want to add a note to as an argument.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store := store.Store
		if store.Current == "" {
			return errors.New("no curent todo task")
		}
		blockId := args[0]
		note, err := cmd.Flags().GetString("note")
		if err != nil {
			return err
		}
		cmdBlock := commands.BlockCmd{ID: store.Current}
		cmdBlock.AddNote(store, blockId, note)
		return nil
	},
}
