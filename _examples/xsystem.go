package main

import (
	"fmt"
	"../xsystem"
)

func main() {
	x := xsystem.New()
	fmt.Println(x.GetMem())
	fmt.Println(x.GetCPU())
}