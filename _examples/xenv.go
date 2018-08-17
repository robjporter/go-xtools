package main

import (
	"fmt"
	"../xenv"
)

func main() {
	env := xenv.New()
	fmt.Println("SIZE: ", env.Size())
	fmt.Println(env.Get("GOPATH"))
	fmt.Println("SIZE: ", env.Size())
	env.Add("TEST", "TESTING")
	fmt.Println("SIZE: ", env.Size())
	fmt.Println("DATA: ", env.GetAll())
}
