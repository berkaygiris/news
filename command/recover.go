package command

import (
	"github.com/spf13/cobra"
	"newsletter/persistence"
)

func init() {
	RootCmd.AddCommand(recoverCmd)
}

var recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "recover " + persistence.SavedPath + " into " + persistence.Path,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_ = persistence.Recover()
	},
}
