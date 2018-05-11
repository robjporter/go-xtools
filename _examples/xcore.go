package main

import (
	"fmt"

	"../xcore"
)

func main() {
	fmt.Println(xcore.F(test())[0])
}

func test() (bool, int, string, interface{}) {
	type t struct {
		name string
		age  int
	}
	tmp := t{name: "NAME", age: 22}
	return true, 4, "test", tmp
}
