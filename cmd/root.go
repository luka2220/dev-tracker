package cmd

import (
	"os"

	"github.com/luka2220/devtasks/tui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devtasks",
	Short: "Intract with your development task boards",
	Long: `The root command of the application allows
you to interact with the boards you have created. When 
the command is ran, the currently activte board will be 
displayed.

NOTE: If the are currently no boards created, you will
be notified and the program will exit.
`,
	Run: func(cmd *cobra.Command, args []string) {
		tui.StartRootTui()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
