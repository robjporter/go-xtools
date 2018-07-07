package xborders

import (
	"bytes"
	"strings"

	"github.com/mattn/go-runewidth"
	"github.com/robjporter/go-xtools/xgraphics/xdisplay"
)

func processBorder(b borderStyle, spacers int, lines []string, center bool) string {
	max := 0
	for _, l := range lines {
		w := widthANSI(l)
		if w > max {
			max = w
		}
	}

	max += spacers * 2
	leftpadding := 0

	if center {
		_, width, _ := xdisplay.New().TerminalSize()
		if max < width {
			tmp := width - max
			leftpadding = tmp / 2
		}
	}

	padding := strings.Repeat(" ", leftpadding)

	var w bytes.Buffer
	ml := widthANSI(b[borderMiddle])
	maxp := max + ml*2

	w.WriteString(padding + b[borderTopLeft])
	tl := widthANSI(b[borderTop])
	for i := 0; i < maxp; i += tl {
		w.WriteString(b[borderTop])
	}
	w.WriteString(b[borderTopRight])
	w.WriteString("\n")

	var esc string
	for _, line := range lines {
		// continue escapes
		w.WriteString(esc)

		w.WriteString(padding + b[borderLeft])
		w.WriteString(b[borderMiddle])
		w.WriteString(strings.Repeat(" ", spacers) + line)
		l := widthANSI(line) - ml + spacers
		for i := 0; i < max-l; i += ml {
			w.WriteString(b[borderMiddle])
		}
		w.WriteString(b[borderRight])
		if esc != "" {
			w.WriteString("\x1b[0m")
		}
		w.WriteString("\n")

		esc = escCont(line)
	}

	w.WriteString(padding + b[borderBottomLeft])
	bl := widthANSI(b[borderBottom])
	for i := 0; i < maxp; i += bl {
		w.WriteString(b[borderBottom])
	}
	w.WriteString(b[borderBottomRight])

	return w.String()
}

func widthANSI(s string) int {
	a := false
	w := 0
	for _, r := range s {
		if r == '\x1b' {
			a = true
		}
		if a {
			if r == 'm' {
				a = false
			}
			continue
		}
		c := runewidth.RuneWidth(r)
		if c != -1 {
			w += c
		}
	}
	return w
}

func escCont(s string) string {
	i := strings.LastIndex(s, "\x1b")
	if i == -1 {
		return ""
	}
	for j := i; j < len(s); j++ {
		if s[j] == 'm' {
			return s[i : j+1]
		}
	}
	return s[i:]
}
