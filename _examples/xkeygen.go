package main

import (
	"fmt"

	"../xkeygen"
)

func main() {
	fmt.Println(xkeygen.NewKey(12))
	fmt.Println(xkeygen.NewPass(12))
	fmt.Println(xkeygen.NewAPIKey(12))
}
