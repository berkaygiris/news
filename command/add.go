package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"newsletter/domain"
	"newsletter/persistence"
	"time"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [Title] [Publisher] [URL] [Comment] [Type] [Date | Optional]",
	Short: "Add a new entry to the newsletter",
	Args:  cobra.MinimumNArgs(5),
	Run: func(cmd *cobra.Command, args []string) {
		date := func() string {
			if len(args) > 2 {
				parsedDate, err := time.Parse(time.DateOnly, args[5])
				if err == nil {
					return parsedDate.Format(time.DateOnly)
				}
			}
			return time.Now().Format(time.DateOnly)
		}()

		newEntry := domain.Entry{
			Title:     args[0],
			Publisher: args[1],
			URL:       args[2],
			Comment:   args[3],
			Type:      args[4],
			Date:      date,
		}

		// Load existing data from the JSON file

		existingEntries, err := persistence.Read()
		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

		// Append the new entry to existing data
		existingEntries = append(existingEntries, newEntry)

		err = persistence.SaveAll(existingEntries)
		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

		fmt.Println("Entry added successfully!")
	},
}
