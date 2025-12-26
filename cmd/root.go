package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tcoyne1729/todo/cmd/block"
	"github.com/tcoyne1729/todo/cmd/tag"
	"github.com/tcoyne1729/todo/internal/storage"
	"github.com/tcoyne1729/todo/store"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A task manager app.",
	Long: `The purpose of the app is to fill the gap between a task manager like
	Github or Gitlab and a notebook. We want to collect data on what we are working on
	which can then be used to generate summaries and statistics later.
	We want to capture what work is done, what is left to do, what got blocked, etc.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		store.Store = storage.NewStore() // loads JSON files
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(tag.TagCmd)
	rootCmd.AddCommand(block.BlockCmd)
}
