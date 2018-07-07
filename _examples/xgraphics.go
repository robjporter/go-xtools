package main

import (
	"fmt"
	"image/color"

	"../xgraphics"
	"sync"
	"time"
)

func main() {
	var lines []string
	lines = append(lines, "")
	lines = append(lines, "This is the title")
	lines = append(lines, "")

	g := xgraphics.New()
	g.Display.ClearScreen()
	h, w, e := g.Display.TerminalSize()
	fmt.Println("Error:", e)
	fmt.Printf("%dhx%dw\n", h, w)

	g.Borders.SetBorderStyle("circles")
	g.Borders.SetContent(lines)
	g.Borders.SetSpacer(12)

	g.Borders.PrintBorder(true)

	g.Icons.PrintIconStyles()
	g.Icons.PrintIcon("cross")
	g.Icons.PrintIcon("alien")

	g.Lines.SetLineStyle("line0")
	g.Lines.PrintLine(60)

	g.Lines.PrintLineTitleLeft(60, "Results")
	g.Lines.PrintLineTitleCenter(60, "* Results *")
	g.Lines.PrintLineTitleRight(60, "Results")

	g.Colors.ClearScreen()
	fmt.Println(g.Colors.BoldText("BOLD").PlainText(" - ").ItalicText("ITALIC").PlainText(" - ").UnderlineText("UNDERLINE").PlainText(" - ").InverseText("INVERSE").PlainText(" - ").StrikeText("STRIKE").PlainText(" - ").String())

	fmt.Println(g.Colors.BGColorText(color.RGBA{R: 191, G: 10, B: 25}, "TEST").Reset())
	fmt.Println(g.Colors.PlainText(" - ").FGColorText(color.RGBA{R: 191, G: 10, B: 25}, "TEST").Reset())
	fmt.Println(g.Colors.PlainText(" - ").BGColor(color.RGBA{R: 191, G: 10, B: 25}).FGColor(color.RGBA{R: 0, G: 0, B: 0}).Text("TEST").Reset().Reset())
	fmt.Println(g.Colors.PlainText(" - ").Foreground(g.Colors.Color["RED"]).Text("TEST").Reset())
	fmt.Println(g.Colors.PlainText(" - ").Background(g.Colors.Color["RED"]).Text("TEST").Reset())
	fmt.Println(g.Colors.PlainText(" - ").Background(g.Colors.Color["WHITE"]).Foreground(g.Colors.Color["BLUE"]).Text("TEST").Reset())

	a := g.Colors.NewString().Fore(g.Colors.Color["WHITE"]).Back(g.Colors.Color["BLUE"]).Text("WHITE").Reset()
	fmt.Println(a.String())

	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["GREEN"]).Text("GREEN").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["GREEN"]).Bold().Text("GREEN").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["GREEN"]).Bold().Underline().Text("GREEN TEXT").PlainText(" - ").Foreground(g.Colors.Color["BLUE"]).Text("BLUE TEXT").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["BLACK"]).Text("BLACK").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["WHITE"]).Text("WHITE").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTRED"]).Text("LIGHT RED").Reset())
	fmt.Println(g.Colors.NewString().Fore(g.Colors.Color["RED"]).Text("RED").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKRED"]).Text("DARK RED").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTBLUE"]).Text("LIGHT BLUE").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["BLUE"]).Text("BLUE").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKBLUE"]).Text("DARK BLUE").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTGREEN"]).Text("LIGHT GREEN").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["GREEN"]).Text("GREEN").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKGREEN"]).Text("DARK GREEN").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTYELLOW"]).Text("LIGHT YELLOW").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["YELLOW"]).Text("YELLOW").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKYELLOW"]).Text("DARK YELLOW").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTORANGE"]).Text("LIGHT ORANGE").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["ORANGE"]).Text("ORANGE").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKORANGE"]).Text("DARK ORANGE").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTPINK"]).Text("LIGHT PINK").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["PINK"]).Text("PINK").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKPINK"]).Text("DARK PINK").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["LIGHTGREY"]).Text("LIGHT GREY").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["GREY"]).Text("GREY").Reset())
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["DARKGREY"]).Text("DARK GREY").Reset())

	fmt.Println("")
	fmt.Println(g.Colors.NewString().Foreground(g.Colors.Color["YAHOO"]).Text("YAHOO").Reset())
	fmt.Println("")

	fmt.Println(g.Colors.NewString().BlinkingText("   Blinking Red   "))

	pgb := xgraphics.NewProgress("多线程进度条")
	xgraphics.ProgressLn("进度条1")
	b := pgb.NewBar("1st", 20000)
	xgraphics.ProgressLn("进度条2")
	b2 := pgb.NewBar("2st", 10000)
	xgraphics.ProgressLn("进度条3")
	b3 := pgb.NewBar("3st", 30000)

	b.SetSpeedSection(900, 100)
	b2.SetSpeedSection(900, 100)
	b3.SetSpeedSection(900, 100)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 20000; i++ {
			b.Add()
			time.Sleep(time.Second / 2000)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			b2.Add()
			time.Sleep(time.Second / 1000)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 30000; i++ {
			b3.Add()
			time.Sleep(time.Second / 3000)
		}
	}()
	wg.Wait()
}
