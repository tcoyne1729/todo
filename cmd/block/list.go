package block

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/internal/commands"
	"github.com/tcoyne1729/todo/store"
)

func init() {
	blockListCmd.Flags().BoolP("all", "a", false, "list all blockers including closed ones. default false.")
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
			fmt.Printf("%d: %v", i, blocker)
		}
		return nil
	},
}
