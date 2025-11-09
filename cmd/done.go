package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "mark the active task as done",
	Long:  `mark the active task as done`,
	Run: func(cmd *cobra.Command, args []string) {
		addCmd := commands.DoneCmd{}
		addCmd.Run(store.Store)
	},
}
