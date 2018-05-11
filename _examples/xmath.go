package main

import (
	"fmt"

	"../xmath"
)

func main() {
	mathsdivide()
	mathsoperators()
}

func mathsdivide() {
	fmt.Println("")
	fmt.Println("MATHS DIVIDE *******************************************************")
	fmt.Println("DivideDetail(44,6):           >", xmath.DivideDetail(44, 6))
	fmt.Println("RoundPlus_Divide(44,6):       >", xmath.RoundPrecise(xmath.DivideDetail(44, 6), 0.5, 5))
	fmt.Println("RoundPlus_Divide(44,6):       >", xmath.Round(xmath.DivideDetail(44, 6), 4))
	fmt.Println("DivideDetailRound(44,6,2):    >", xmath.DivideDetailRound(44, 6, 2))
	fmt.Println("Round_DivideDetail(44,6):     >", xmath.Round(xmath.DivideDetail(44, 6), 2))
	fmt.Println("Divide(44,6):                 >", xmath.Divide(44, 6))
	fmt.Println("Remainder(44,6):              >", xmath.Remainder(44, 6))
}

func mathsoperators() {
	fmt.Println("")
	fmt.Println("MATHS *******************************************************")
	fmt.Println("Maths('*',5,2):               >", xmath.Operator("*", 5, 2))
	fmt.Println("Maths('+',5,2):               >", xmath.Operator("+", 5, 2))
	fmt.Println("Maths('-',5,2):               >", xmath.Operator("-", 5, 2))
	fmt.Println("Maths('/',5,2):               >", xmath.Operator("/", 5, 2))
	fmt.Println("Maths('%',5,2):               >", xmath.Operator("%", 5, 2))
	fmt.Println("Maths(1):                     >", xmath.ToRomanNumeral(1))
	fmt.Println("Maths(10):                    >", xmath.ToRomanNumeral(10))
	fmt.Println("Maths(100):                   >", xmath.ToRomanNumeral(100))
	fmt.Println("Maths(1000):                  >", xmath.ToRomanNumeral(1000))
	fmt.Println("Maths(10000):                 >", xmath.ToRomanNumeral(10000))
}
