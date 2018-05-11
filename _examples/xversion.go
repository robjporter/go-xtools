package main

import (
	"fmt"

	"../xversion"
)

func main() {
	parsestringversion()
	parsescompareversion()
}

func parsestringversion() {
	fmt.Println("")
	fmt.Println("PARSE STRING VERSION *******************************************************")
	fmt.Println("Parse Version 1.02.3.             >", xversion.ParseVersionString("1.02.3."))
	fmt.Println("Parse Version 1.2..3              >", xversion.ParseVersionString("1.2..3"))
	fmt.Println("Parse Version 1.00.2              >", xversion.ParseVersionString("1.00.2"))
	fmt.Println("Parse Version 1.02.a              >", xversion.ParseVersionString("1.02.a"))
	fmt.Println("Parse Version 1.02.a.b            >", xversion.ParseVersionString("1.02.a.b"))
	fmt.Println("Parse Version 1.02(a)             >", xversion.ParseVersionString("1.02(a)"))
	fmt.Println("Parse Version 1.02(4a)            >", xversion.ParseVersionString("1.02(4a)"))
	fmt.Println("Parse Version 1.02(4b)            >", xversion.ParseVersionString("1.02(4b)"))
	fmt.Println("Parse Version 1.02(5a)            >", xversion.ParseVersionString("1.02(5a)"))
}

func parsescompareversion() {
	fmt.Println("")
	fmt.Println("PARSE COMPARE VERSION *******************************************************")
	fmt.Println("Compare 3.1(4b) v 2.02(3a)        >", xversion.CompareStrings("3.1(4b)", "2.02(3a)"))
	fmt.Println("Compare 3.1(4b) v 3.1(4b)         >", xversion.CompareStrings("3.1(4b)", "3.1(4b)"))
	fmt.Println("Compare 2.02(3a) v 3.1(4b)        >", xversion.CompareStrings("2.02(3a)", "3.1(4b)"))
}
