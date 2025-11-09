package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(noteCmd)
	noteCmd.Flags().StringP("id", "i", "", "ID of task to add note to (defaults to current task)")
}

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "add a note to the task",
	Long:  `add a note to the current task. You can add a note to a non-current task with a flag`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		note := args[0]
		id, _ := cmd.Flags().GetString("id")
		noteCmd := commands.NoteCmd{
			Note: note,
			ID:   id,
		}
		noteCmd.Run(store.Store)
	},
}
