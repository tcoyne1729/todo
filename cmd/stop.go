package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop working on a task",
	Long:  `stop working on a task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// we assume the user doesnt want to use the other parameters and keep them internal
		stopCmd := commands.StopCmd{
			ID: id,
		}
		stopCmd.Run(store.Store)
	},
}
