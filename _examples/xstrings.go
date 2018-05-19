package main

import (
	"fmt"

	"../xstrings"
)

func main() {
	fmt.Println(xstrings.Center("This is a test", "=", 50))
	fmt.Println(xstrings.Substring("This is a test", 5, 6))
	fmt.Println(xstrings.UUID4())
	fmt.Println(xstrings.ToTrain("ThisA_test"))
	fmt.Println(xstrings.ToSpinal("ThisA_test"))
	fmt.Println(xstrings.ToSnake("ThisA_test"))
	fmt.Println(xstrings.ToSnakeUpper("ThisA_test"))
	fmt.Println(xstrings.ToCamel("ThisA_test"))
	fmt.Println(xstrings.ToCamelLower("ThisA_test"))
	fmt.Println(xstrings.IsInSlice("test", []string{"a", "b", "tester", "testing", "test"}))
	fmt.Println(xstrings.PosInSlice("test", []string{"a", "b", "tester", "testing", "test"}))
	fmt.Println(xstrings.StringsBetween("[what is between]", "[", "]"))
	fmt.Println(xstrings.StringBetween("[what is between]", "[", "]"))
}
