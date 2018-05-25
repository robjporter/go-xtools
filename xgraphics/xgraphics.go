package xgraphics

import (
	"./xborders"
	"./xcolors"
	"./xdisplay"
	"./xicons"
	"./xlines"
	"./xprogress"
	"./xtermo"
)


type Graphics struct {
	Borders *xborders.Borders
	Colors  *xcolors.Colors
	Icons   *xicons.Icons
	Lines   *xlines.Lines
	Display *xdisplay.Display
	Terminal *xtermo.Framebuffer
}

func New() *Graphics {
	xtermo.Init()
	defer xtermo.Stop()
	w, h, _ := xtermo.Size()
	return &Graphics{Borders: xborders.New(), Colors: xcolors.New(), Icons: xicons.New(), Lines: xlines.New(), Display: xdisplay.New(),Terminal:xtermo.NewFramebuffer(w,h)}
}

func NewProgress(title string) *xprogress.Xprogress {
	return xprogress.New(title)
}

func ProgressLn(args ...interface{}) {
	xprogress.Println(args...)
}