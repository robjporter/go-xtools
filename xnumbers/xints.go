package xnumbers

import "math/rand"

func IsInIntSlice(str int, list []int) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func PosInSlice(str int, list []int) int {
	for k, v := range list {
		if v == str {
			return k
		}
	}
	return -1
}

func RandInt(low, high int) int {
	if low >= high {
		tmp := high
		high = low
		low = tmp
	}
	return int(rand.Float64()*float64(high-low)) + low
}
