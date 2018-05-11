package xgraphics

import (
	"./xborders"
	"./xcolors"
	"./xdisplay"
	"./xicons"
	"./xlines"
)

type Graphics struct {
	Borders *xborders.Borders
	Colors  *xcolors.Colors
	Icons   *xicons.Icons
	Lines   *xlines.Lines
	Display *xdisplay.Display
}

func New() *Graphics {
	return &Graphics{Borders: xborders.New(), Colors: xcolors.New(), Icons: xicons.New(), Lines: xlines.New(), Display: xdisplay.New()}
}
