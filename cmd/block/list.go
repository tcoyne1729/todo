package block

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/internal/models"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	blockListCmd.Flags().BoolP("all", "a", false, "list all blockers including closed ones. default false.")
}

func printBlocker(index int, blocker *models.Blocker) {
	fmt.Println("---")
	fmt.Printf("Index: %d\nID: %s\nText: %s\nBlocked By: %s\n", index, blocker.ID, blocker.Text, blocker.BlockedBy)
	for i, note := range blocker.BlockerNotes.Data {
		fmt.Println("------")
		fmt.Printf("note %d: %s\n", i, note.Text)
		fmt.Println("------")
	}
	fmt.Println("---")
}

var blockListCmd = &cobra.Command{
	Use:   "list",
	Short: "list the blockers for the current task",
	Long: `list all the blockers which are active. If you want to list all including the closed blockers, then
	use the -a flag`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		store := store.Store
		taskID := store.Current
		if taskID == "" {
			return errors.New("no current active task")
		}
		cmdBlock := commands.BlockCmd{ID: taskID}
		all, _ := cmd.Flags().GetBool("all")
		fmt.Printf("task: %s\n", store.Current)
		lsBlockers, err := cmdBlock.List(store, all)
		fmt.Printf("got %d blockers for this task\n", len(lsBlockers))
		if err != nil {
			fmt.Printf("failed to load blockers for task: %s\n", store.Current)
			fmt.Printf("error: %v\n", err)
		}
		for i, blocker := range lsBlockers {
			printBlocker(i, blocker)
		}
		return nil
	},
}
