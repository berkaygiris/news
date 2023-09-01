package command

import (
	"github.com/spf13/cobra"
	"newsletter/persistence"
)

func init() {
	RootCmd.AddCommand(clearCmd)
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear " + persistence.Path,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_ = persistence.Clear()
	},
}
