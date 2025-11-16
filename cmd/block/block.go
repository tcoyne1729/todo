package block

import (
	"github.com/spf13/cobra"
)

func init() {
	BlockCmd.AddCommand(blockAddCmd)
}

var BlockCmd = &cobra.Command{
	Use:   "block",
	Short: "update any blockers on this task",
	Long:  `you can add descriptive information about blockers to the task`,
}
