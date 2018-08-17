package main

import (
	"fmt"

	"../xnumbers"
)

func main() {
	fmt.Println(xnumbers.IsInIntSlice(4, []int{2, 4, 6, 8, 10}))
	fmt.Println(xnumbers.PosInSlice(4, []int{2, 4, 6, 8, 10}))
	fmt.Println(xnumbers.RandInt(1, 10))
	fmt.Println(xnumbers.RandInt(100, 10))
	fmt.Println(xnumbers.RandFloat(1, 10))
	fmt.Println(xnumbers.RandFloat(100, 10))
	fmt.Println(xnumbers.RandomNumbers(10))
	fmt.Println(xnumbers.RandomNumbers(20))
	fmt.Println(xnumbers.RandomNumbers(50))
}
