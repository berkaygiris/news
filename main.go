package main

import "newsletter/command"

func main() {

	// Add subcommands here

	if err := command.RootCmd.Execute(); err != nil {
		// Handle error
		//panic(err)
	}
}
