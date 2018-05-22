package main

import (
	"../xflags"
)

// go run xflags1.go -f test

func main() {
	// Declare variables and their defaults
	var stringFlag = "defaultValue"

	// Add a flag
	xflags.String(&stringFlag, "f", "flag", "A test string flag")

	// Parse the flag
	xflags.Parse()

	// Use the flag
	print(stringFlag)
}