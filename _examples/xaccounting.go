package main

import (
	"fmt"
	"../xaccounting"
)

func main() {
	x,e := xaccounting.New().Symbol("£")
	x.Precision(2)
	m,e:=x.Money(1000000.123456)
	fmt.Println("MONEY: ",m)
	fmt.Println("ERROR: ",e)
	fmt.Println("ISNEGATIVE: ",x.IsNegative())
	fmt.Println("ISPOSITIVE: ",x.IsPositive())
	fmt.Println("ISZERO: ",x.IsZero())
	fmt.Println("TAX: ",x.Tax(40))
	fmt.Println("SPLIT 3: ",x.Split(3))
	fmt.Println("SPLIT 5: ",x.Split(5))
	fmt.Println("SPLIT 10: ",x.Split(10))
	fmt.Println("SPLIT 25: ",x.Split(25))
}
