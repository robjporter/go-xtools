package xmath

import "math"

func Divide(a, b int) int {
	if a > 0 && b > 0 {
		value := a / b
		return value
	} else {
		return 0.0
	}
}

func DivideDetail(a, b int) float64 {
	if a > 0 && b > 0 {
		value := float64(a) / float64(b)
		return value
	} else {
		return 0.0
	}
}

func DivideDetailRound(a, b, c int) float64 {
	if a > 0 && b > 0 {
		value := float64(a) / float64(b)
		return Round(value, c)
	} else {
		return 0.0
	}
}

func Remainder(a, b int) float64 {
	if a > 0 && b > 0 {
		value := a % b
		return float64(value)
	} else {
		return 0.0
	}
}

func Round(v float64, d int) float64 {
	return RoundPrecise(v, 1, d)
}

func RoundPrecise(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
