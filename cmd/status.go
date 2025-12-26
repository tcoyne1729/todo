package cmd

import (
	"fmt"
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
		statusCmd := commands.StatusCmd{}
		err := statusCmd.Run(store.Store)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
