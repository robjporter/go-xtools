package is

import (
	"net"
	"regexp"
	"strings"
	"time"
	"net/url"
	"reflect"
)

func IsEmpty(tmp string) bool {
	if len(tmp) == 0 {
		return true
	} else {
		return false
	}
}

func IsNumber(s interface{}) bool {
	switch s.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return true
	}

	return false
}

func IsEmail(email string) bool {
	p := `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
	ok, err := regexp.MatchString(p, email)
	if err != nil {
		return false
	}
	return ok
}

func IsEmail2(email string) bool {
	r := regexp.MustCompile(`.+@.+\..+`)
	return r.MatchString(email)
}

func IsNotEmpty(tmp string) bool {
	if len(tmp) == 0 {
		return false
	} else {
		return true
	}
}

func IsAlphaNumeric(char string) bool {
	return IsStringContainAlpha(char) && IsStringContainNumber(char)
}

func IsStringAllAlpha(char string) bool {
	allAlpha := true
	letters := []byte(char)
	for i := 0; i < len(letters); i++ {
		if 'a' <= letters[i] && letters[i] <= 'z' || 'A' <= letters[i] && letters[i] <= 'Z' {

		} else {
			allAlpha = false
		}
	}
	return allAlpha
}

func IsStringContainAlpha(char string) bool {
	allAlpha := false
	letters := []byte(char)
	for i := 0; i < len(letters); i++ {
		if 'a' <= letters[i] && letters[i] <= 'z' || 'A' <= letters[i] && letters[i] <= 'Z' {
			allAlpha = true
		}
	}
	return allAlpha
}

func IsStringContainNumber(char string) bool {
	allAlpha := false
	letters := []byte(char)
	for i := 0; i < len(letters); i++ {
		if '0' <= letters[i] && letters[i] <= '9' {
			allAlpha = true
		}
	}
	return allAlpha
}

func IsStringUppercase(char string) bool {
	allUpper := false
	letters := []byte(char)
	for i := 0; i < len(letters); i++ {
		if 'A' <= letters[i] && letters[i] <= 'Z' {
		} else {
			allUpper = false
		}
	}
	return allUpper
}

func IsCharUppercase(char byte) bool {
	return 'A' <= char && char <= 'Z'
}

func IsBlank(tmp string) bool {
	tmp = strings.TrimSpace(tmp)
	return IsEmpty(tmp)
}

func IsNotBlank(tmp string) bool {
	tmp = strings.TrimSpace(tmp)
	return IsNotEmpty(tmp)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func IsInt(input interface{}) bool {
	switch input.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}
func IsFloat(input interface{}) bool {
	switch input.(type) {
	case float32, float64:
		return true
	}
	return false
}

func IsBool(input interface{}) bool {
	switch input.(type) {
	case bool:
		return true
	}
	return false
}

func IsString(input interface{}) bool {
	switch input.(type) {
	case string:
		return true
	}
	return false
}

func IsTime(input interface{}) bool {
	switch input.(type) {
	case time.Time:
		return true
	}
	return false
}

func IsInStringSlice(needle string, stack []string) bool {
	if len(stack) > 0 {
		for i := 0; i < len(stack); i++ {
			if needle == stack[i] {
				return true
			}
		}
	}
	return false
}

func IsInIntSlice(needle int, stack []int) bool {
	if len(stack) > 0 {
		for i := 0; i < len(stack); i++ {
			if needle == stack[i] {
				return true
			}
		}
	}
	return false
}

func IsIPAddress(ip string) bool {
	test := net.ParseIP(ip)
	if test.To4() == nil {
		return false
	}
	return true
}

func IsInSlice(v interface{}, sl []interface{}) bool {
	switch reflect.TypeOf(sl).Kind() {
		case reflect.Slice:
			for _, vv := range sl {
				if vv == v {
					return true
				}
			}
		default:
			return false
	}
	return false
}

func IsURL(uri string) bool {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return strings.HasPrefix(u.Scheme, "http")
}

func IsUUID(uuid string) bool {
	r := regexp.MustCompile("[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}")
	return r.MatchString(uuid)
}