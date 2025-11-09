package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("body", "b", "", "Description of the task")
	addCmd.Flags().StringArrayP("tag", "t", []string{""}, "Tags to add to the task")
	addCmd.Flags().Int32P("priority", "p", 0, "Priority of the task")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Long:  `add a new task`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		body, _ := cmd.Flags().GetString("body")
		priority, _ := cmd.Flags().GetInt32("priority")
		tags, _ := cmd.Flags().GetStringArray("tag")
		if len(tags) > 1 {
			errors.New("not implemented more than one tag yet...\n")
		}
		if len(tags) == 0 {
			tags = []string{""}
		}
		addCmd := commands.AddCmd{
			Title:    title,
			Body:     body,
			Priority: int(priority),
			Tag:      tags[0],
		}
		addCmd.Run(store.Store)
	},
}
