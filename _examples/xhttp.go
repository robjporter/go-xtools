package main

import (
	"fmt"

	"../xhttp"
)

func main() {
	s := xhttp.NewSite()
	s.SetURL("http://www.google.co.uk")
	fmt.Println(s.Status())
	fmt.Println(s.LastStatus())
	fmt.Println(s.LastStatusString())

	// TODO:
}
