package main

import (
	"fmt"

	"../xenvironment"
)

func main() {
	systemenvironments()
}

func systemenvironments() {
	fmt.Println("")
	fmt.Println("SYSTEM ENVIRONMENTS *******************************************************")
	fmt.Println("ENVIRONMENT GO PATH BIN:          >", xenvironment.GOPATHBIN())
	fmt.Println("ENVIRONMENT GO PATH:              >", xenvironment.GOPATH())
	fmt.Println("ENVIRONMENT PATH SEPARATOR:       >", xenvironment.PathSeparator())
	fmt.Println("ENVIRONMENT LIST SEPARATOR:       >", xenvironment.ListSeparator())
	fmt.Println("ENVIRONMENT IS COMPILED:          >", xenvironment.IsCompiled())
	fmt.Println("ENVIRONMENT BUILD DEBUG:          >", xenvironment.BuildDebug())
	fmt.Println("ENVIRONMENT CHECK ARCH:           >", xenvironment.CheckArchitecture())
	fmt.Println("ENVIRONMENT BUILD STAMP:          >", xenvironment.BuildStamp())
	fmt.Println("ENVIRONMENT BUILD HOST:           >", xenvironment.BuildHost())
	fmt.Println("ENVIRONMENT COMPILER:             >", xenvironment.Compiler())
	fmt.Println("ENVIRONMENT GO ARCH:              >", xenvironment.GOARCH())
	fmt.Println("ENVIRONMENT GO OS:                >", xenvironment.GOOS())
	fmt.Println("ENVIRONMENT GO ROOT:              >", xenvironment.GOROOT())
	fmt.Println("ENVIRONMENT GO VERSION:           >", xenvironment.GOVER())
	fmt.Println("ENVIRONMENT NUMBER CPU:           >", xenvironment.NumCPU())
	fmt.Println("ENVIRONMENT FORMATTED TIME:       >", xenvironment.GetFormattedTime())
	fmt.Println("ENVIRONMENT USERNAME:             >", xenvironment.GetUsername())
	fmt.Println("ENVIRONMENT VARIABLES:            >", xenvironment.GetAllEnvironment())
}
