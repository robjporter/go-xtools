package gojsonq

import (
	"reflect"
	"strings"
)

const (
	signEq           = "="
	signEqEng        = "eq"
	signNotEq        = "!="
	signNotEqEng     = "neq"
	signNotEqAnother = "<>"
	signGt           = ">"
	signGtEng        = "gt"
	signLt           = "<"
	signLtEng        = "lt"
	signGtE          = ">="
	signGtEEng       = "gte"
	signLtE          = "<="
	signLtEEng       = "lte"

	signStrictContains = "strictContains"
	signContains       = "contains"
	signEndsWith       = "endsWith"
	signStartsWith     = "startsWith"
	signIn             = "in"
	signNotIn          = "notIn"
)

// QueryFunc describes a conditional function which perform comparison
type QueryFunc func(x, y interface{}) bool

// eq checks whether x, y are deeply eq
func eq(x, y interface{}) bool {
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		y = fv
	}
	return reflect.DeepEqual(x, y)
}

// neq checks whether x, y are deeply not equal
func neq(x, y interface{}) bool {
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		y = fv
	}
	return !eq(x, y)
}

//  gt checks whether x is greather than y
func gt(x, y interface{}) bool {
	xv, ok := x.(float64)
	if !ok {
		return false
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv > fv
	}
	return false
}

//  lt checks whether x is less than y
func lt(x, y interface{}) bool {
	xv, ok := x.(float64)
	if !ok {
		return false
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv < fv
	}
	return false
}

//  gte checks whether x is greather than or equal y
func gte(x, y interface{}) bool {
	xv, ok := x.(float64)
	if !ok {
		return false
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv >= fv
	}
	return false
}

//  lte checks whether x is less than or equal y
func lte(x, y interface{}) bool {
	xv, ok := x.(float64)
	if !ok {
		return false
	}
	// if the y value is numeric (int/int8-int64/float32/float64) then convert to float64
	if fv, ok := toFloat64(y); ok {
		return xv <= fv
	}
	return false
}

// strStrictContains works like `select * from table where column like %a%`.
func strStrictContains(x, y interface{}) bool {
	xv, okX := x.(string)
	if !okX {
		return false
	}
	yv, okY := y.(string)
	if !okY {
		return false
	}
	return strings.Contains(xv, yv)
}

// strContains works like `select * from table where column like %a%`. this is case insensitive serarch
func strContains(x, y interface{}) bool {
	xv, okX := x.(string)
	if !okX {
		return false
	}
	yv, okY := y.(string)
	if !okY {
		return false
	}
	return strings.Contains(strings.ToLower(xv), strings.ToLower(yv))
}

// strStartsWith works like `select * from table where column like a%`. Basically find value starts with 'a'
func strStartsWith(x, y interface{}) bool {
	xv, okX := x.(string)
	if !okX {
		return false
	}
	yv, okY := y.(string)
	if !okY {
		return false
	}
	return strings.HasPrefix(xv, yv)
}

// strEndsWith works like `select * from table where column like %o`. Basically find value ends with 'a'
func strEndsWith(x, y interface{}) bool {
	xv, okX := x.(string)
	if !okX {
		return false
	}
	yv, okY := y.(string)
	if !okY {
		return false
	}
	return strings.HasSuffix(xv, yv)
}

// in checks whether any value exist. e.g: in("id", []int{1,3,5,8})
func in(x, y interface{}) bool {
	if yv, ok := y.([]string); ok {
		for _, v := range yv {
			if eq(x, v) {
				return true
			}
		}
	}
	if yv, ok := y.([]int); ok {
		for _, v := range yv {
			if eq(x, v) {
				return true
			}
		}
	}
	if yv, ok := y.([]float64); ok {
		for _, v := range yv {
			if eq(x, v) {
				return true
			}
		}
	}
	return false
}

// notIn checks whether any value does not contain provided values. e.g: in("id", []int{1,3,5,8})
func notIn(x, y interface{}) bool {
	if yv, ok := y.([]string); ok {
		for _, v := range yv {
			if eq(x, v) {
				return false
			}
		}
	}
	if yv, ok := y.([]int); ok {
		for _, v := range yv {
			if eq(x, v) {
				return false
			}
		}
	}
	if yv, ok := y.([]float64); ok {
		for _, v := range yv {
			if eq(x, v) {
				return false
			}
		}
	}
	return true
}

func loadDefaultQueryMap() map[string]QueryFunc {
	// queryMap contains the registered conditional functions
	var queryMap = make(map[string]QueryFunc)

	queryMap[signEq] = eq
	queryMap[signEqEng] = eq

	queryMap[signNotEq] = neq
	queryMap[signNotEqEng] = neq
	queryMap[signNotEqAnother] = neq // also an alias of not equal

	queryMap[signGt] = gt
	queryMap[signGtEng] = gt

	queryMap[signLt] = lt
	queryMap[signLtEng] = lt

	queryMap[signGtE] = gte
	queryMap[signGtEEng] = gte

	queryMap[signLtE] = lte
	queryMap[signLtEEng] = lte

	queryMap[signStrictContains] = strStrictContains
	queryMap[signContains] = strContains
	queryMap[signStartsWith] = strStartsWith
	queryMap[signEndsWith] = strEndsWith

	queryMap[signIn] = in
	queryMap[signNotIn] = notIn
	return queryMap
}
