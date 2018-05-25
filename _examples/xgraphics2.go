package main

import (
	"../xgraphics"
	"time"
)


func main() {
	g := xgraphics.New()
	g.Terminal.ASCIIRect(0, 1, g.Terminal.GetWidth(), 3, false, false) // Draw a 20x20 ASCII rectangle
	g.Terminal.ASCIIRect(0, 3, g.Terminal.GetWidth()/2, g.Terminal.GetHeight()-3, false, false)
	text := "I'm now using termo!"
	a := len(text)
	pos := g.Terminal.GetWidth()/2 - (a/2)
	g.Terminal.SetText(pos, 2, text) // Draw text
	g.Terminal.CenterText(g.Terminal.GetWidth()/2, g.Terminal.GetHeight()/2-5, "About termo_example.\n")
	g.Terminal.AttribRect(0, 4, g.Terminal.GetWidth(), 1, g.Terminal.GetBoldWhiteOnBlack()) // Set character colors/attributes
	b := g.Terminal.NewCellState()
	b.Attrib = g.Terminal.GetAttrNone()
	b.BGColor = g.Terminal.GetColorBlue()
	b.FGColor = g.Terminal.GetColorGray()
	g.Terminal.AttribRect(0, 1, g.Terminal.GetWidth(), g.Terminal.GetHeight(), b)
	g.Terminal.ASCIIRect(0, 1, g.Terminal.GetWidth(), g.Terminal.GetHeight()-1, false, false)
	g.Terminal.Flush()
	time.Sleep(2*time.Second)

	g.Terminal.CenterText(g.Terminal.GetWidth()/2, g.Terminal.GetHeight()/2-5, "About termo_example 2.\n")
	g.Terminal.Flush()

}