package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start working on a task",
	Long:  `start working on a task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		startCmd := commands.StartCmd{
			ID: id,
		}
		startCmd.Run(store.Store)
	},
}
