package cmd

import (
	"github.com/luka2220/devtasks/tui/initialization"
	"github.com/spf13/cobra"
)

// NOTE:
// Subcommand to initialize a new development task board
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new development task board",
	Long: `This sub-command is used create and initialize
a new development task board.

Propmts for the new board name and wether to set it as
the active board.
`,
	Run: func(cmd *cobra.Command, args []string) {
		initialization.StartProjectInitTui()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
