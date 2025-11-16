package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("body", "b", "", "Description of the task")
	updateCmd.Flags().StringP("title", "t", "", "Title of the task")
	updateCmd.Flags().Int32P("priority", "p", 0, "Priority of the task")
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update an existing task",
	Long:  `update an existing task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		body, _ := cmd.Flags().GetString("body")
		title, _ := cmd.Flags().GetString("title")
		priority, _ := cmd.Flags().GetInt32("priority")
		updateCmd := commands.UpdateCmd{
			ID:       id,
			Title:    title,
			Body:     body,
			Priority: int(priority),
		}
		updateCmd.Run(store.Store)
	},
}
