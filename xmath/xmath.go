package xmath

func Operator(op string, a, b int) int {
	if op == "*" {
		return a * b
	} else if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "/" {
		return a / b
	} else if op == "%" {
		return a % b
	}
	return -1
}

func ToRomanNumeral(x int) string {
	switch {
	case x >= 1000:
		return "M" + ToRomanNumeral(x-1000)
	case x >= 900:
		return "CM" + ToRomanNumeral(x-900)
	case x >= 500:
		return "D" + ToRomanNumeral(x-500)
	case x >= 400:
		return "CD" + ToRomanNumeral(x-400)
	case x >= 100:
		return "C" + ToRomanNumeral(x-100)
	case x >= 90:
		return "XC" + ToRomanNumeral(x-90)
	case x >= 50:
		return "L" + ToRomanNumeral(x-50)
	case x >= 40:
		return "XL" + ToRomanNumeral(x-40)
	case x >= 10:
		return "X" + ToRomanNumeral(x-10)
	case x >= 9:
		return "IX" + ToRomanNumeral(x-9)
	case x >= 5:
		return "V" + ToRomanNumeral(x-5)
	case x >= 4:
		return "IV" + ToRomanNumeral(x-4)
	case x >= 1:
		return "I" + ToRomanNumeral(x-1)
	}
	return ""
}
