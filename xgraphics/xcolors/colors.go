package xcolors

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
)

const (
	FOREGROUND    = "\033[38;5;"
	FOREGROUNDEND = "m"
	BACKGROUND    = "\033[48;5;"
	BACKGROUNDEND = "m"
	ALPHA         = uint8(255)
	ESCAPE        = "\x1b"
	RESET         = "\033[0m"
	CLEAR         = "\033[2J"
	ANSIClear     = "\033[0m"
)

const (
	RED         ColorNumber = 0
	BLUE        ColorNumber = 1
	GREEN       ColorNumber = 2
	BLACK       ColorNumber = 3
	WHITE       ColorNumber = 4
	LIGHTRED    ColorNumber = 5
	DARKRED     ColorNumber = 6
	LIGHTBLUE   ColorNumber = 7
	DARKBLUE    ColorNumber = 8
	LIGHTGREEN  ColorNumber = 9
	DARKGREEN   ColorNumber = 10
	LIGHTYELLOW ColorNumber = 11
	YELLOW      ColorNumber = 12
	DARKYELLOW  ColorNumber = 13
	LIGHTORANGE ColorNumber = 14
	ORANGE      ColorNumber = 15
	DARKORANGE  ColorNumber = 16
	LIGHTPINK   ColorNumber = 17
	PINK        ColorNumber = 18
	DARKPINK    ColorNumber = 19
	LIGHTGREY   ColorNumber = 20
	GREY        ColorNumber = 21
	DARKGREY    ColorNumber = 22
	YAHOO       ColorNumber = 23
)

var colorsMap = map[string][]string{
	"bold":          []string{"\x1B[1m", "\x1B[22m"},
	"italic":        []string{"\u001b[3m", "\u001b[23m"},
	"underline":     []string{"\x1B[4m", "\x1B[24m"},
	"inverse":       []string{"\x1B[7m", "\x1B[27m"},
	"strikethrough": []string{"\x1B[9m", "\x1B[29m"},
	"blinking":      []string{"\033[5;41m", "\033[25;41m"},
}

var colors = map[string]color.RGBA{
	"black": color.RGBA{0, 0, 0, ALPHA},
	"white": color.RGBA{255, 255, 255, ALPHA},

	"lred": color.RGBA{0xff, 0x69, 0x61, ALPHA},
	"red":  color.RGBA{255, 0, 0, ALPHA},
	"dred": color.RGBA{0x32, 0x14, 0x14, ALPHA},

	"lblue": color.RGBA{0x0, 0xee, 0xee, ALPHA},
	"blue":  color.RGBA{0x1e, 0x90, 0xff, ALPHA},
	"dblue": color.RGBA{0, 0, 255, ALPHA},

	"lgreen": color.RGBA{0x90, 0xee, 0x90, ALPHA},
	"green":  color.RGBA{0, 255, 0, ALPHA},
	"dgreen": color.RGBA{0x0, 0x64, 0x0, ALPHA},

	"lyellow": color.RGBA{0xfc, 0xe8, 0x83, ALPHA},
	"yellow":  color.RGBA{0xff, 0xcc, 0x0, ALPHA},
	"dyellow": color.RGBA{0x99, 0x65, 0x15, ALPHA},

	"lorange": color.RGBA{0xed, 0x91, 0x21, ALPHA},
	"orange":  color.RGBA{0xe9, 0x69, 0x2c, ALPHA},
	"dorange": color.RGBA{0xfe, 0x47, 0x28, ALPHA},

	"lpink": color.RGBA{0xe7, 0xac, 0xcf, ALPHA},
	"pink":  color.RGBA{0xff, 0x77, 0xff, ALPHA},
	"dpink": color.RGBA{0x80, 0x0, 0x20, ALPHA},

	"lgrey": color.RGBA{0xc0, 0xc0, 0xc0, ALPHA},
	"grey":  color.RGBA{127, 140, 141, ALPHA},
	"dgrey": color.RGBA{0x1c, 0x35, 0x2d, ALPHA},

	"yahoo": color.RGBA{0x41, 0x0, 0x93, ALPHA},
}

type ColorNumber int

type Colors struct {
	rawString string
	Color     map[string]ColorNumber
}

// INITIALISE ////////////////////////////////////////////////////////////////////////////////////

func New() *Colors {
	c := getColors()
	return &Colors{Color: c}
}

func getColors() map[string]ColorNumber {
	tmp := make(map[string]ColorNumber)
	tmp["RED"] = ColorNumber(0)
	tmp["BLUE"] = ColorNumber(1)
	tmp["GREEN"] = ColorNumber(2)
	tmp["BLACK"] = ColorNumber(3)
	tmp["WHITE"] = ColorNumber(4)
	tmp["LIGHTRED"] = ColorNumber(5)
	tmp["DARKRED"] = ColorNumber(6)
	tmp["LIGHTBLUE"] = ColorNumber(7)
	tmp["DARKBLUE"] = ColorNumber(8)
	tmp["LIGHTGREEN"] = ColorNumber(9)
	tmp["DARKGREEN"] = ColorNumber(10)
	tmp["LIGHTYELLOW"] = ColorNumber(11)
	tmp["YELLOW"] = ColorNumber(12)
	tmp["DARKYELLOW"] = ColorNumber(13)
	tmp["LIGHTORANGE"] = ColorNumber(14)
	tmp["ORANGE"] = ColorNumber(15)
	tmp["DARKORANGE"] = ColorNumber(16)
	tmp["LIGHTPINK"] = ColorNumber(17)
	tmp["PINK"] = ColorNumber(18)
	tmp["DARKPINK"] = ColorNumber(19)
	tmp["LIGHTGREY"] = ColorNumber(20)
	tmp["GREY"] = ColorNumber(21)
	tmp["DARKGREY"] = ColorNumber(22)
	tmp["YAHOO"] = ColorNumber(23)
	return tmp
}

func (c *Colors) NewString() *Colors {
	c.rawString = ""
	return c
}

// COLORS ////////////////////////////////////////////////////////////////////////////////////

func (c *ColorNumber) Color() color.RGBA {
	switch *c {
	case RED:
		return colors["red"]
	case BLUE:
		return colors["blue"]
	case GREEN:
		return colors["green"]
	case BLACK:
		return colors["black"]
	case WHITE:
		return colors["white"]
	case LIGHTRED:
		return colors["lred"]
	case DARKRED:
		return colors["dred"]
	case LIGHTBLUE:
		return colors["lblue"]
	case DARKBLUE:
		return colors["dblue"]
	case LIGHTGREEN:
		return colors["lgreen"]
	case DARKGREEN:
		return colors["dgreen"]
	case LIGHTYELLOW:
		return colors["lyellow"]
	case YELLOW:
		return colors["yellow"]
	case DARKYELLOW:
		return colors["dyellow"]
	case LIGHTORANGE:
		return colors["lorange"]
	case ORANGE:
		return colors["orange"]
	case DARKORANGE:
		return colors["dorange"]
	case LIGHTPINK:
		return colors["lpink"]
	case PINK:
		return colors["pink"]
	case DARKPINK:
		return colors["dpink"]
	case LIGHTGREY:
		return colors["lgrey"]
	case GREY:
		return colors["grey"]
	case DARKGREY:
		return colors["dgrey"]
	case YAHOO:
		return colors["yahoo"]
	}
	return colors["black"]
}

// FORE / BACK ////////////////////////////////////////////////////////////////////////////////////

func (c *Colors) Fore(a ColorNumber) *Colors {
	return c.FGColor(a.Color())
}

func (c *Colors) Foreground(a ColorNumber) *Colors {
	return c.FGColor(a.Color())
}

func (c *Colors) Back(a ColorNumber) *Colors {
	return c.BGColor(a.Color())
}

func (c *Colors) Background(a ColorNumber) *Colors {
	return c.BGColor(a.Color())
}

// CLEAR ////////////////////////////////////////////////////////////////////////////////////

func (c *Colors) ClearScreen() {
	fmt.Println(CLEAR)
}

// EFFECTS ////////////////////////////////////////////////////////////////////////////////////

func (c *Colors) BlinkingText(text string) *Colors {
	c.rawString += colorsMap["blinking"][0] + text + colorsMap["blinking"][1]
	return c
}

func (c *Colors) Blinking() *Colors {
	c.rawString += colorsMap["blinking"][0]
	return c
}

func (c *Colors) BoldText(text string) *Colors {
	c.rawString += colorsMap["bold"][0] + text + colorsMap["bold"][1]
	return c
}

func (c *Colors) Bold() *Colors {
	c.rawString += colorsMap["bold"][0]
	return c
}

func (c *Colors) ItalicText(text string) *Colors {
	c.rawString += colorsMap["italic"][0] + text + colorsMap["italic"][1]
	return c
}
func (c *Colors) Italic() *Colors {
	c.rawString += colorsMap["italic"][0]
	return c
}

func (c *Colors) UnderlineText(text string) *Colors {
	c.rawString += colorsMap["underline"][0] + text + colorsMap["underline"][1]
	return c
}
func (c *Colors) Underline() *Colors {
	c.rawString += colorsMap["underline"][0]
	return c
}

func (c *Colors) InverseText(text string) *Colors {
	c.rawString += colorsMap["inverse"][0] + text + colorsMap["inverse"][1]
	return c
}
func (c *Colors) Inverse() *Colors {
	c.rawString += colorsMap["inverse"][0]
	return c
}

func (c *Colors) StrikeText(text string) *Colors {
	c.rawString += colorsMap["strikethrough"][0] + text + colorsMap["strikethrough"][1]
	return c
}
func (c *Colors) Strike() *Colors {
	c.rawString += colorsMap["strikethrough"][0]
	return c
}

func (c *Colors) PlainText(text string) *Colors {
	c.rawString += RESET + text + RESET
	return c
}

func (c *Colors) Plain() *Colors {
	c.rawString += RESET
	return c
}

// OUTPUT ////////////////////////////////////////////////////////////////////////////////////

func (c *Colors) String() string {
	return c.rawString
}

// RAW ////////////////////////////////////////////////////////////////////////////////////

func (c *Colors) BGColor(col color.Color) *Colors {
	c.rawString += buildColor(BACKGROUND, getColor(col), BACKGROUNDEND)
	return c
}

func (c *Colors) FGColor(col color.Color) *Colors {
	c.rawString += buildColor(FOREGROUND, getColor(col), FOREGROUNDEND)
	return c
}

func (c *Colors) BGColorText(col color.Color, text string) *Colors {
	c.rawString += buildColor(BACKGROUND, getColor(col), BACKGROUNDEND) + text + RESET
	return c
}

func (c *Colors) FGColorText(col color.Color, text string) *Colors {
	c.rawString += buildColor(FOREGROUND, getColor(col), FOREGROUNDEND) + text + RESET
	return c
}

// SUPPORT ////////////////////////////////////////////////////////////////////////////////////

func (c *Colors) Reset() *Colors {
	c.rawString += RESET
	return c
}

func (c *Colors) Text(text string) *Colors {
	c.rawString += text
	return c
}

// INTERNAL ////////////////////////////////////////////////////////////////////////////////////

func buildColor(start, col, end string) string {
	return start + col + end
}

func getColor(col color.Color) string {
	const begin = 16
	const ratio = 5.0 / (1<<16 - 1)
	rf, gf, bf, _ := col.RGBA()
	r := int(round(ratio * float64(rf)))
	g := int(round(ratio * float64(gf)))
	b := int(round(ratio * float64(bf)))
	val := r*6*6 + g*6 + b + begin
	return strconv.Itoa(val)
}

func round(x float64) float64 {
	return math.Floor(x + 0.5)
}
