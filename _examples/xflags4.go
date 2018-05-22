package main

import (
	"../xflags"
	"fmt"
)

// go run xflags4.go subcommandExample -t test nestedSubcommand -f 2 -y

func main() {
	// Declare variables and their defaults
	var stringFlagF = "defaultValueF"
	var intFlagT = 3
	var boolFlagB bool

	// Create the subcommand
	subcommandExample := xflags.NewSubcommand("subcommandExample")
	nestedSubcommand := xflags.NewSubcommand("nestedSubcommand")

	// Add a flag to the subcommand
	subcommandExample.String(&stringFlagF, "t", "testFlag", "A test string flag")
	nestedSubcommand.Int(&intFlagT, "f", "flag", "A test int flag")

	// add a global bool flag for fun
	xflags.Bool(&boolFlagB, "y", "yes", "A sample boolean flag")

	//  the nested subcommand to the parent subcommand at position 1
	subcommandExample.AttachSubcommand(nestedSubcommand, 1)

	//  the base subcommand to the parser at position 1
	xflags.AttachSubcommand(subcommandExample, 1)

	// Parse the subcommand and all flags
	xflags.Parse()

	// Use the flags and trailing arguments
	print(stringFlagF)
	print(intFlagT)

	// we can check if a subcommand was used easily
	if nestedSubcommand.Used {
		print(boolFlagB)
	}
	print(xflags.TrailingArguments[0:])
	fmt.Println("")
}