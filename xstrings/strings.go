package xstrings

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

var noop = func(a rune) rune { return a }

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

func Substring(in string, start int, length uint) string {
	if length > 0 {
		size := len(in)
		if start < 0 {
			frontIndex := size + start
			if frontIndex < 0 {
				return ""
			}
			start = frontIndex
		}
		if start >= size {
			return ""
		}
		rearIndex := start + int(length)
		if rearIndex >= size {
			return in[start:]
		}
		return in[start:rearIndex]
	}
	return ""
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
	n, err := io.ReadFull(rand.Reader, uuid)
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
