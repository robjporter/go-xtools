package xlines

import (
	"fmt"
	"strings"
)

type Lines struct {
	line string
}

func New() *Lines {
	return &Lines{}
}

func (l *Lines) SetLineStyle(style string) *Lines {
	l.line = style
	return l
}

func (l *Lines) GetLine(length int) string {
	return strings.Repeat(lineStyles[l.line][0], length)
}

func (l *Lines) PrintLine(length int) {
	fmt.Println(strings.Repeat(lineStyles[l.line][0], length))
}

func (l *Lines) GetLineTitleLeft(length int, title string) string {
	left := strings.Repeat(lineStyles[l.line][0], 2)
	center := " " + title + " "
	right := strings.Repeat(lineStyles[l.line][0], length-len(title)+2)

	return left + center + right
}

func (l *Lines) PrintLineTitleLeft(length int, title string) {
	left := strings.Repeat(lineStyles[l.line][0], 2)
	center := " " + title + " "
	right := strings.Repeat(lineStyles[l.line][0], length-len(title)-4)

	fmt.Println(left + center + right)
}

func (l *Lines) GetLineTitleCenter(length int, title string) string {
	size := length / 2
	size -= len(title) / 2
	size -= 1
	left := strings.Repeat(lineStyles[l.line][0], size)
	center := " " + title + " "
	right := strings.Repeat(lineStyles[l.line][0], size)

	return left + center + right
}

func (l *Lines) PrintLineTitleCenter(length int, title string) {
	size := length / 2
	size -= len(title) / 2
	size -= 1
	left := strings.Repeat(lineStyles[l.line][0], size)
	center := " " + title + " "
	right := strings.Repeat(lineStyles[l.line][0], size)

	fmt.Println(left + center + right)
}

func (l *Lines) GetLineTitleRight(length int, title string) string {
	left := strings.Repeat(lineStyles[l.line][0], length-len(title)-4)
	center := " " + title + " "
	right := strings.Repeat(lineStyles[l.line][0], 2)

	return left + center + right
}

func (l *Lines) PrintLineTitleRight(length int, title string) {
	left := strings.Repeat(lineStyles[l.line][0], length-len(title)-4)
	center := " " + title + " "
	right := strings.Repeat(lineStyles[l.line][0], 2)

	fmt.Println(left + center + right)
}
