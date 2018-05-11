package main

import (
	"fmt"
	"time"

	"../xformatting"
)

func main() {
	fmt.Println(xformatting.FormatRoman(50))
	fmt.Println(xformatting.ShortDuration(time.Duration(44 * time.Second)))
	fmt.Println(xformatting.ShortDuration(time.Duration(44 * time.Minute)))
	fmt.Println(xformatting.ShortDuration(time.Duration(44 * time.Hour)))
	fmt.Println(xformatting.LongDuration(time.Duration(44 * (24 * time.Hour))))
	fmt.Println(xformatting.LongDuration(time.Duration(44 * (7 * 24 * time.Hour))))
	fmt.Println(xformatting.LongDuration(time.Duration(44 * (4 * 7 * 24 * time.Hour))))

	fmt.Println(xformatting.ToPrec(3.14567890, 2))
	fmt.Println(xformatting.ToPrec(3.14567890, 3))
	fmt.Println(xformatting.ToPrec(3.14567890, 4))
	fmt.Println(xformatting.ToPrec(3.14567890, 5))

	fmt.Println(xformatting.ToFormattedBytes(1234))
	fmt.Println(xformatting.ToFormattedBytes(12345))
	fmt.Println(xformatting.ToFormattedBytes(123456))
	fmt.Println(xformatting.ToFormattedBytes(1234567))
	fmt.Println(xformatting.ToFormattedBytes(12345678))
	fmt.Println(xformatting.ToFormattedBytes(123456789))
	fmt.Println(xformatting.ToFormattedBytes(1234567890))
	fmt.Println(xformatting.ToFormattedBytes(12345678901))
	fmt.Println(xformatting.ToFormattedBytes(123456789012))
	fmt.Println(xformatting.ToFormattedBytes(1234567890123))
	fmt.Println(xformatting.ToFormattedBytes(12345678901234))
	fmt.Println(xformatting.ToFormattedBytes(123456789012345))
	fmt.Println(xformatting.ToFormattedBytes(1234567890123456))

}
