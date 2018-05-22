package main

import (
	"../xflags"
)

// go run xflags3.go subcommandExample -f test

func main() {
		// Declare variables and their defaults
		var stringFlag = "defaultValue"

		// Create the subcommand
		subcommand := xflags.NewSubcommand("subcommandExample")

		// Add a flag to the subcommand
		subcommand.String(&stringFlag, "f", "flag", "A test string flag")

		// Add the subcommand to the parser at position 1
		xflags.AttachSubcommand(subcommand, 1)

		// Parse the subcommand and all flags
		xflags.Parse()

		// Use the flag
		print(stringFlag)
}