package xdisplay

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Display struct{}

func New() *Display { return &Display{} }

func (d *Display) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (d *Display) TerminalSize() (int, int, error) {
	height := 0
	width := 0
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err == nil {
		splits := strings.Split(string(out), " ")
		if len(splits) == 2 {
			a, e := strconv.Atoi(splits[0])
			if e == nil {
				b, e := strconv.Atoi(strings.TrimRight(splits[1], "\n"))
				height = a
				if e == nil {
					width = b
				}
			}
		}
	}

	return height, width, err
}
