package main

import (
	"fmt"
	"time"

	"../xif"
)

func main() {
	time.Sleep(1 * time.Second)
	ifthen()
	time.Sleep(1 * time.Second)
	ifequals()
	time.Sleep(1 * time.Second)
	ifequalsmultipleint()
	time.Sleep(1 * time.Second)
	ifequalsmultiplestring()
}

func ifthen() {
	fmt.Println("")
	fmt.Println("IF THEN / IF THEN ELSE ******************************************")
	fmt.Println("IfThen(4>2):                  >", xif.IfThen(4 > 2, "4 is greater than 2"))
	fmt.Println("IfThen(4>2):                  >", xif.IfThen(4 > 2, true))
	fmt.Println("IfThenElse(4>2):              >", xif.IfThenElse(4 > 2, "4 is greater than 2", "4 is less than 2"))
	fmt.Println("IfThenElse(4>2):              >", xif.IfThenElse(4 > 2, true, false))
	fmt.Println("IfThen(4>2):                  >", xif.IfThen(1 == 1, "Yes"))            // "Yes"
	fmt.Println("IfThen(4>2):                  >", xif.IfThen(1 != 1, "Woo"))            // nil
	fmt.Println("IfThen(4>2):                  >", xif.IfThen(1 < 2, "Less"))            // "Less"
	fmt.Println("IfThen(4>2):                  >", xif.IfThenElse(1 == 1, "Yes", false)) // "Yes"
	fmt.Println("IfThen(4>2):                  >", xif.IfThenElse(1 != 1, nil, 1))       // 1
	fmt.Println("IfThen(4>2):                  >", xif.IfThenElse(1 < 2, nil, "No"))     // nil
}

func ifequals() {
	fmt.Println("")
	fmt.Println("IF EQUALS *******************************************************")
	fmt.Println("IfEquals(4>2):                >", xif.IfEquals(4 > 2))
	fmt.Println("IfEquals(4<2):                >", xif.IfEquals(4 < 2))
}

func ifequalsmultipleint() {
	fmt.Println("")
	fmt.Println("IF EQUALS MULTIPLE INT *******************************************************")
	fmt.Println("IfEqualsMultipleInt(4,0,2,4): >", xif.IfEqualsMultipleInt(4, 0, 2, 4))
	fmt.Println("IfEqualsMultipleInt(4,0,1):   >", xif.IfEqualsMultipleInt(4, 0, 1))
}

func ifequalsmultiplestring() {
	fmt.Println("")
	fmt.Println("IF EQUALS MULTIPLE STRING *******************************************************")
	fmt.Println("IfEqualsMultipleString():     >", xif.IfEqualsMultipleString("a", "a", "b"))
	fmt.Println("IfEqualsMultipleString():     >", xif.IfEqualsMultipleString("a", "b"))
}
