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
	fmt.Println("PRE TAX: ",x.PreTax(40))
	fmt.Println("SPLIT 3: ",x.Split(3))
	fmt.Println("SPLIT 5: ",x.Split(5))
	fmt.Println("SPLIT 10: ",x.Split(10))
	fmt.Println("SPLIT 25: ",x.Split(25))


	x2,e2 := xaccounting.New().Symbol("£")
	x2.Precision(2)
	x2.Money(880000.123456)
	fmt.Println("EQUALS: ",x.Equals(x2))
	fmt.Println("GREATER THAN: ",x.GreaterThan(x2))
	fmt.Println("GREATER THAN EQUAL TO: ",x.GreaterThanEqual(x2))
	fmt.Println("LESS THAN: ",x.LessThan(x2))
	fmt.Println("LESS THAN EQUAL TO: ",x.LessThaEqual(x2))
	if e2 == nil {
		fmt.Println("ADDITION: ",x.Add(x2).String())
		fmt.Println("SUBTRACTION: ",x.Sub(x2).String())
	}
	fmt.Println("DIVIDE: ",x.Divide(4).String())
	fmt.Println("MULTIPLICATION: ",x.Multiply(16).String())


	fmt.Println("POSITIVE: ",x.Positive())
	fmt.Println("NEGATIVE: ",x.Negative())
}
