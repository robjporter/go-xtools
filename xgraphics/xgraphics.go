package xgraphics

import (
	"github.com/robjporter/go-xtools/xgraphics/xborders"
	"github.com/robjporter/go-xtools/xgraphics/xcolors"
	"github.com/robjporter/go-xtools/xgraphics/xdisplay"
	"github.com/robjporter/go-xtools/xgraphics/xicons"
	"github.com/robjporter/go-xtools/xgraphics/xlines"
	"github.com/robjporter/go-xtools/xgraphics/xprogress"
	"github.com/robjporter/go-xtools/xgraphics/xtermo"
)

type Graphics struct {
	Borders  *xborders.Borders
	Colors   *xcolors.Colors
	Icons    *xicons.Icons
	Lines    *xlines.Lines
	Display  *xdisplay.Display
	Terminal *xtermo.Framebuffer
}

func New() *Graphics {
	xtermo.Init()
	defer xtermo.Stop()
	w, h, _ := xtermo.Size()
	return &Graphics{Borders: xborders.New(), Colors: xcolors.New(), Icons: xicons.New(), Lines: xlines.New(), Display: xdisplay.New(), Terminal: xtermo.NewFramebuffer(w, h)}
}

func NewProgress(title string) *xprogress.Xprogress {
	return xprogress.New(title)
}

func ProgressLn(args ...interface{}) {
	xprogress.Println(args...)
}
