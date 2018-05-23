package main

import (
	"fmt"
	"../xaccounting"
)

func main() {
	x,e := xaccounting.New().Symbol("£")
	x,e = x.Precision(2)
	fmt.Println(x.Money(1000000.123456))
	fmt.Println("MONEY: ",x)
	fmt.Println("ERROR: ",e)
}
