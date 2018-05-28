package main

import (
	"fmt"

	"../xfiles"
)

func main() {
	test := "This is a random string to test the features of xcompression = test!"
	fmt.Println("== CHECKSUM =================================================================")
	check := xfiles.GetChecksum(test)
	fmt.Println("CHECKSUM: ", check)
	fmt.Println("INTEGRITY: ", xfiles.CheckIntegrity(test, check))
	fmt.Println("== FILES ====================================================================")
	fmt.Println("EXISTS: ", xfiles.Exists("xas.go"))
	fmt.Println("DIRECTORY EMPTY: ", xfiles.DirIsEmpty("./"))
	a, e := xfiles.LookupPath("./;~/", "xas.go")
	fmt.Println("LOOKUP PATH: ", a, e)
	fmt.Println("COPY FILE/DIR ", xfiles.Copy("xas.go", "as.go"))
	fmt.Println("CURRENT PATH: ",xfiles.CurrentPath())
}
