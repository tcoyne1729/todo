package block

import (
	"github.com/spf13/cobra"
)

func init() {
	BlockCmd.AddCommand(blockAddCmd)
	BlockCmd.AddCommand(blockCompleteCmd)
	BlockCmd.AddCommand(blockListCmd)
	BlockCmd.AddCommand(blockNoteCmd)
}

var BlockID string

var BlockCmd = &cobra.Command{
	Use:   "block",
	Short: "update any blockers on the active task",
	Long: `you can add descriptive information about blockers to the active task.
	This is only availble for the active task unless you provide the task id with a flag.`,
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	if len(args) > 0 {
	// 		BlockID = args[0]
	// 	}
	// },
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
