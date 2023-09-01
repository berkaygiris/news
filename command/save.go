package command

import (
	"github.com/spf13/cobra"
	"newsletter/persistence"
)

func init() {
	RootCmd.AddCommand(saveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save as" + persistence.SavedPath,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		_ = persistence.Save()
	},
}
