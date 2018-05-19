package main

import (
	"fmt"

	"../xregex"
)

func main() {
	fmt.Println(xregex.Email("test"))
	fmt.Println(xregex.Email("test@"))
	fmt.Println(xregex.Email("test@gmail"))
	fmt.Println(xregex.Email("test@gmail.c"))
	fmt.Println(xregex.Email("test@gmail.com"))
}
