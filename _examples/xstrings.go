package main

import (
	"fmt"

	"../xstrings"
)

func main() {
	fmt.Println(xstrings.Center("This is a test", "=", 50))
	fmt.Println(xstrings.SubString("This is a test", 5, 6))
	fmt.Println(xstrings.SubStringStart("This is a test", 6))
	fmt.Println(xstrings.SubStringEnd("This is a test", 6))
	fmt.Println(xstrings.Truncate("This is a test", 6, false))
	fmt.Println(xstrings.Truncate("This is a test", 6, true))

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
	fmt.Println(xstrings.Reverse("This is a test"))
	fmt.Println(xstrings.Format("The {} says {}", "cow", "MOO!"))
	fmt.Println(xstrings.RandStringWithLengthLimit(10))
	fmt.Println(xstrings.RandStringWithLengthLimit(20))
	fmt.Println(xstrings.RandStringWithLengthLimit(50))
	fmt.Println(xstrings.Sha1("InString"))
	fmt.Println(xstrings.Sha256("InString"))
	fmt.Println(xstrings.MaskString("TESTING", "TESTINGTESTING", 4, 22))
}
