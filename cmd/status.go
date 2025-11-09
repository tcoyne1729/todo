package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "get the status of the active task",
	Long:  `get the status of the active task`,
	Run: func(cmd *cobra.Command, args []string) {
		addCmd := commands.StatusCmd{}
		addCmd.Run(store.Store)
	},
}
