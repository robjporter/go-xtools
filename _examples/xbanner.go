package main

import (
	"fmt"

	"../xbanner"
	"github.com/robjporter/go-utils/terminal/colours"
)

func main() {
	displaybanner()
	picturebanner()
}

func displaybanner() {
	xbanner.PrintNewFigure("LIBRARY", "varsity", true)
}

func picturebanner() {
	fmt.Println("")
	fmt.Println("PICTURE BANNER **********************************************")
	str := colours.Blocks("{{@brightblue}}{{@!BLANK(80)}}\n")
	str += colours.Blocks("{{@red}}{{@!BLANK(18)}}{{@brightblue}}{{@!BLANK(16)}}{{@black}}{{@!BLANK(30)}}{{@brightblue}}{{@!BLANK(16)}}\n")
	str += colours.Blocks("{{@red}}{{@!BLANK(32)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(30)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(14)}}\n")
	str += colours.Blocks("{{@brightred}}{{@!BLANK(4)}}{{@red}}{{@!BLANK(26)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@magenta}}{{@!BLANK(22)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(12)}}\n")
	str += colours.Blocks("{{@brightred}}{{@!BLANK(30)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@magenta}}{{@!BLANK(16)}}{{@black}}{{@!BLANK(4)}}{{@magenta}}{{@!BLANK(6)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(4)}}{{@brightblue}}{{@!BLANK(6)}}\n")
	str += colours.Blocks("{{@brightred}}{{@!BLANK(30)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(16)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(6)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(4)}}\n")
	str += colours.Blocks("{{@brightyellow}}{{@!BLANK(18)}}{{@brightred}}{{@!BLANK(12)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(16)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@magenta}}{{@!BLANK(6)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(4)}}\n")
	str += colours.Blocks("{{@brightyellow}}{{@!BLANK(22)}}{{@black}}{{@!BLANK(2)}}{{@brightyellow}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(16)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(8)}}{{@white}}{{@!BLANK(8)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(4)}}\n")
	str += colours.Blocks("{{@brightyellow}}{{@!BLANK(20)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(2)}}{{@brightyellow}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(16)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(22)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(4)}}\n")
	str += colours.Blocks("{{@brightgreen}}{{@!BLANK(18)}}{{@brightyellow}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(8)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(14)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(26)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(2)}}\n")
	str += colours.Blocks("{{@brightgreen}}{{@!BLANK(22)}}{{@white}}{{@!BLANK(8)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(14)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@brightyellow}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(10)}}{{@brightyellow}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(2)}}\n")
	str += colours.Blocks("{{@brightgreen}}{{@!BLANK(22)}}{{@black}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(14)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(2)}}\n")
	str += colours.Blocks("{{@blue}}{{@!BLANK(18)}}{{@brightgreen}}{{@!BLANK(8)}}{{@black}}{{@!BLANK(6)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(14)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@magenta}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(16)}}{{@magenta}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(2)}}\n")
	str += colours.Blocks("{{@blue}}{{@!BLANK(30)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@magenta}}{{@!BLANK(14)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(12)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(4)}}\n")
	str += colours.Blocks("{{@brightblue}}{{@!BLANK(18)}}{{@blue}}{{@!BLANK(10)}}{{@black}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(6)}}{{@magenta}}{{@!BLANK(14)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(18)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(6)}}\n")
	str += colours.Blocks("{{@brightblue}}{{@!BLANK(26)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(20)}}{{@black}}{{@!BLANK(18)}}{{@brightblue}}{{@!BLANK(8)}}\n")
	str += colours.Blocks("{{@brightblue}}{{@!BLANK(24)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(6)}}{{@black}}{{@!BLANK(32)}}{{@white}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(12)}}\n")
	str += colours.Blocks("{{@brightblue}}{{@!BLANK(24)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(2)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@brightblue}}{{@!BLANK(12)}}{{@black}}{{@!BLANK(2)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(4)}}{{@white}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(2)}}{{@brightblue}}{{@!BLANK(12)}}\n")
	str += colours.Blocks("{{@brightblue}}{{@!BLANK(24)}}{{@black}}{{@!BLANK(6)}}{{@brightblue}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(6)}}{{@brightblue}}{{@!BLANK(12)}}{{@black}}{{@!BLANK(6)}}{{@brightblue}}{{@!BLANK(4)}}{{@black}}{{@!BLANK(6)}}{{@brightblue}}{{@!BLANK(12)}}\n")
	str += colours.Blocks("{{@brightblue}}{{@!BLANK(80)}}\n")
	fmt.Println(str)
}
