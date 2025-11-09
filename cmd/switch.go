package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(switchCmd)
}

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "switch to a new task",
	Long:  `switch to a new task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		// we assume the user doesnt want to use the other parameters and keep them internal
		switchCmd := commands.SwitchCmd{
			ID: id,
		}
		switchCmd.Run(store.Store)
	},
}
