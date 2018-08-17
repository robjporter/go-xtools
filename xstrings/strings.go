package xstrings

import (
	"bytes"
	"math/rand"
	"fmt"
	"io"
	"strings"
	"unicode"
	"unicode/utf8"
	"time"
	"crypto/sha1"
	"encoding/hex"
	"crypto/sha256"
	rand2 "crypto/rand"
)

const (
	letterBytes   = `123456789afhjkoqrsuvwxyzAFHJKQRSUVWXYZ`
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var (
	noop = func(a rune) rune { return a }
	src  = rand.NewSource(time.Now().UnixNano())
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func IsInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func PosInSlice(str string, list []string) int {
	for k, v := range list {
		if v == str {
			return k
		}
	}
	return -1
}

func SubString(str string, startIndex, endIndex int) string {
	return string([]rune(str)[startIndex:endIndex])
}

func SubStringStart(str string, startIndex int) string {
	endIndex := utf8.RuneCountInString(str)
	return SubString(str, startIndex, endIndex)
}

func SubStringEnd(str string, endIndex int) string {
	return SubString(str, 0, endIndex)
}

func Truncate(str string, length int, withExtenders bool) string {
	extenders := "..."
	if utf8.RuneCountInString(str) <= length {
		return str
	}
	if withExtenders {
		length -= len(extenders)
	}
	return SubStringEnd(str, length) + extenders
}

func StringsBetween(str, start, end string) (between []string) {
	between = make([]string, 0)
	fsplit := strings.Split(str, start)
	for i := 1; i < len(fsplit); i++ {
		ssplit := strings.SplitN(fsplit[i], end, 2)
		if len(ssplit) >= 2 {
			between = append(between, ssplit[0])
		}
	}
	return
}

func StringBetween(str, start, end string) string {
	fsplit := strings.Split(str, start)
	if len(fsplit) == 2 {
		return strings.Split(fsplit[1], end)[0]
	}
	return ""
}

func Center(str, padding string, width uint) string {
	text := []rune(str)
	if len(text) >= int(width) {
		return str
	}
	padrunes := []rune(padding)

	out := make([]rune, int(width))
	pos := 0

	padwidth := int(width)/2 - 2 - len(text)/2
	if len(str)%2 == 0 {
		padwidth++
	}
	for i := 0; i < padwidth; i++ {
		out[pos] = padrunes[i%len(padrunes)]
		pos++
	}
	out[pos] = ' '
	pos++
	for i := 0; i < len(text); i++ {
		out[pos] = text[i]
		pos++
	}
	out[pos] = ' '
	pos++

	if len(str)%2 == 1 {
		padwidth++
	}
	for i := 0; i < padwidth; i++ {
		out[pos] = padrunes[i%len(padrunes)]
		pos++
	}
	return string(out)
}

func Len(str string) int {
	return utf8.RuneCountInString(str)
}

func writePadString(output *bytes.Buffer, pad string, padLen, remains int) {
	var r rune
	var size int

	repeats := remains / padLen

	for i := 0; i < repeats; i++ {
		output.WriteString(pad)
	}

	remains = remains % padLen

	if remains != 0 {
		for i := 0; i < remains; i++ {
			r, size = utf8.DecodeRuneInString(pad)
			output.WriteRune(r)
			pad = pad[size:]
		}
	}
}

func ToTrain(s string) string {
	return snaker(s, '-', unicode.ToUpper, unicode.ToUpper, noop)
}

func ToSpinal(s string) string {
	return snaker(s, '-', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

func ToSnake(s string) string {
	return snaker(s, '_', unicode.ToLower, unicode.ToLower, unicode.ToLower)
}

func ToSnakeUpper(s string) string {
	return snaker(s, '_', unicode.ToUpper, unicode.ToUpper, unicode.ToUpper)
}

func ToCamel(s string) string {
	return snaker(s, rune(0), unicode.ToUpper, unicode.ToUpper, noop)
}

func ToCamelLower(s string) string {
	return snaker(s, rune(0), unicode.ToLower, unicode.ToUpper, noop)
}

func UUID4() (string, bool) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand2.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", false
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), true
}

func snaker(s string, wordSeparator rune, firstRune func(rune) rune, firstRuneOfWord func(rune) rune, otherRunes func(rune) rune) string {
	useWordSeperator := wordSeparator != rune(0)
	newS := []rune{}

	// pops a rune off newS
	lastRuneIsWordSeparator := func() bool {
		if len(newS) > 0 {
			return newS[len(newS)-1] == wordSeparator
		}
		return true
	}

	prev := wordSeparator
	for _, cur := range s {
		isWordBoundary := (unicode.IsLower(prev) && unicode.IsUpper(cur)) || unicode.IsSpace(prev)

		if !unicode.IsLetter(cur) {
			// ignore
		} else if isWordBoundary {
			if useWordSeperator && !lastRuneIsWordSeparator() {
				newS = append(newS, wordSeparator)
			}
			newS = append(newS, firstRuneOfWord(cur))
		} else {
			newS = append(newS, otherRunes(cur))
		}

		prev = cur
	}

	if len(newS) > 0 {
		newS[0] = firstRune(newS[0])
	}

	return string(newS)
}

func RandStringWithLengthLimit(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func Sha1(in string) string {
	hasher := sha1.New()
	hasher.Write([]byte(in))
	out := hasher.Sum(nil)

	return hex.EncodeToString(out)
}

func Sha256(in string) string {
	hasher := sha256.New()
	hasher.Write([]byte(in))
	out := hasher.Sum(nil)

	return hex.EncodeToString(out)
}

func MaskString(orig, mask string, revealLength, length int) string {

	if orig == "" {
		return ""
	}

	if length <= 0 {
		length = len(orig)
	}

	str := ""
	for len(str) < length {
		str = str + mask
	}

	if revealLength == -1 {
		return str
	}

	// fmt.Println(str, length, revealLength, orig, len(orig))
	str = fmt.Sprintf("%s%s", str[:length-revealLength], orig[len(orig)-revealLength:])

	return str
}
