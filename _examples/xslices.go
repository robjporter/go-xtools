package main

import (
	"fmt"
	"../xslices"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	fmt.Println(xslices.IntSliceToString(a, ":"))
	fmt.Println(xslices.IntSliceContains(a, 2))
	fmt.Println(xslices.StringSliceContains(b, "eight"))
}
