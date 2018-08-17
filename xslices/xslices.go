package xslices

import (
	"strconv"
)

func IntSliceToString(list []int, sep string) string {
	result := ""
	last := len(list) - 1
	for i, v := range list {
		result = result + strconv.Itoa(v)
		if i != last {
			result = result + sep
		}
	}
	return result
}

func IntSliceContains(slice []int, check int) bool {
	for _, v := range slice {
		if v == check {
			return true
		}
	}
	return false
}

func StringSliceContains(slice []string, check string) bool {
	for _, v := range slice {
		if v == check {
			return true
		}
	}
	return false
}
