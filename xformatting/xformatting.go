// Package fmtutil implements formatting for numbers as common meaningful
// values.
package xformatting

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// SI represents an integer which can format itself with SI prefixes.
type SI uint64

const (
	K = 1024 << (10 * iota)
	M
	G
	T
	P
	E
)

func ToFormattedBytes(x uint64) string {
	n := 0.0
	s := ""
	switch {
	case x < K:
		return strconv.FormatUint(uint64(x), 10)
	case x < M:
		s = "k"
		n = float64(x) / K
	case x < G:
		s = "M"
		n = float64(x) / M
	case x < T:
		s = "G"
		n = float64(x) / G
	case x < P:
		s = "T"
		n = float64(x) / T
	case x < E:
		s = "P"
		n = float64(x) / P
	default:
		s = "E"
		n = float64(x) / E
	}

	return strconv.FormatFloat(ToPrec(n, 1), 'f', -1, 64) + s
}

// ToPrec rounds a number to the given number of total digits.
func ToPrec(n float64, prec int) float64 {
	n *= float64(prec) * 10
	x := float64(int64(n + 0.5))
	return x / (float64(prec) * 10)
}

const (
	Sec   = time.Second
	Min   = Sec * 60
	Hr    = Min * 60
	Day   = Hr * 24
	Week  = Day * 7
	Month = Day * 30
	Year  = Day * 365
)

// LongDuration formats a duration that is most likely much longer than what
// package time will handle. It uses the units seconds, minutes, hours, days,
// weeks, months (30 days), and years (365 days).
func LongDuration(n time.Duration) string {
	p := func(n time.Duration, s string) string {
		return fmt.Sprintf("%d%s", n, s)
	}

	switch {
	case n < Sec:
		return n.String()
	case n < Min:
		return p(n/Sec, "s")
	case n < Hr:
		return p(n/Min, "m")
	case n < Day:
		return p(n/Hr, "h")
	case n < 2*Week:
		return p(n/Day, "d")
	case n < Month:
		return p(n/Week, "w")
	case n < Year:
		return p(n/Month, "mo")
	default:
		return p(n/Year, "y")
	}
}

// HMS formats a duration as a colon-separated timestamp.
//
// If n is less than 60 seconds, the format will be 0:ss.
// If n is less than 60 minutes, the format will be m:ss.
// Otherwise, the format will be h:mm:ss.
func ShortDuration(n time.Duration) string {
	if n < time.Second {
		return "0:00"
	}
	t := n / time.Second
	sec := t % 60
	if t < 60 {
		return fmt.Sprintf("0:%02d", sec)
	}
	t /= 60
	min := t % 60
	if t < 60 {
		return fmt.Sprintf("%d:%02d", min, sec)
	}
	return fmt.Sprintf("%d:%02d:%02d", t/60, min, sec)
}

var romanNumerals = [][]string{
	{"I", "V", "X"},
	{"X", "L", "C"},
	{"C", "D", "M"},
}

// FormatRoman formats an integer as Roman numerals. It panics if n < 1, much
// like a Roman would when confronted with the concept of zero or negative
// numbers.
func FormatRoman(n int) string {
	if n < 1 {
		panic("Roman calligrapher unable to comprehend concept of zero or negative numbers")
	}

	var (
		i int
		s string
	)

	for _, c := range romanNumerals {
		i = n
		n /= 10
		i -= n * 10
		s = formatRomanDigit(i, c[0], c[1], c[2]) + s

		if n == 0 {
			return s
		}
	}

	return strings.Repeat("M", n) + s
}

// formatRomanDigit formats a single digit in the order of magnitude identified
// by the given one, five, and ten numerals.
func formatRomanDigit(n int, one, five, ten string) string {
	switch {
	case n < 4:
		return strings.Repeat(one, n)
	case n == 4:
		return one + five
	case n < 9:
		return five + strings.Repeat(one, n-5)
	default:
		return one + ten
	}
}

/*



	fmt.Println(fmtutil.FormatRoman(50))
	fmt.Println(fmtutil.ShortDuration(time.Duration(44 * time.Second)))
	fmt.Println(fmtutil.ShortDuration(time.Duration(44 * time.Minute)))
	fmt.Println(fmtutil.ShortDuration(time.Duration(44 * time.Hour)))
	fmt.Println(fmtutil.LongDuration(time.Duration(44 * (24 * time.Hour))))
	fmt.Println(fmtutil.LongDuration(time.Duration(44 * (7 * 24 * time.Hour))))
	fmt.Println(fmtutil.LongDuration(time.Duration(44 * (4 * 7 * 24 * time.Hour))))

	fmt.Println(fmtutil.ToPrec(3.14567890, 2))
	fmt.Println(fmtutil.ToPrec(3.14567890, 3))
	fmt.Println(fmtutil.ToPrec(3.14567890, 4))
	fmt.Println(fmtutil.ToPrec(3.14567890, 5))

	fmt.Println(fmtutil.ToFormattedBytes(1234))
	fmt.Println(fmtutil.ToFormattedBytes(12345))
	fmt.Println(fmtutil.ToFormattedBytes(123456))
	fmt.Println(fmtutil.ToFormattedBytes(1234567))
	fmt.Println(fmtutil.ToFormattedBytes(12345678))
	fmt.Println(fmtutil.ToFormattedBytes(123456789))
	fmt.Println(fmtutil.ToFormattedBytes(1234567890))
	fmt.Println(fmtutil.ToFormattedBytes(12345678901))
	fmt.Println(fmtutil.ToFormattedBytes(123456789012))
	fmt.Println(fmtutil.ToFormattedBytes(1234567890123))
	fmt.Println(fmtutil.ToFormattedBytes(12345678901234))
	fmt.Println(fmtutil.ToFormattedBytes(123456789012345))
	fmt.Println(fmtutil.ToFormattedBytes(1234567890123456))


*/
