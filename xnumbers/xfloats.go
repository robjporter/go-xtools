package xnumbers

import (
	"math/rand"
)

func RandFloat(low, high float64) float64 {
	if low >= high {
		tmp := high
		high = low
		low = tmp
	}
	return rand.Float64()*(high-low) + low
}
