package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "all tasks (include done)")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  `list the active tasks. Use flags to list all tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		listCmd := commands.ListCmd{
			All: all,
		}
		err := listCmd.Run(store.Store)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	},
}
