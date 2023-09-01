package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"newsletter/persistence"
)

func init() {
	RootCmd.AddCommand(printCmd)
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Add a new entry to the newsletter",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		existingEntries, err := persistence.Read()

		existingEntries.Print()
		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}
	},
}
