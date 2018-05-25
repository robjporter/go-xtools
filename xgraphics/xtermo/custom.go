package xtermo


func (f *Framebuffer) NewCellState() CellState {
	return CellState{}
}

func (f *Framebuffer) GetWidth() int {
	return width
}

func (f *Framebuffer) GetHeight() int {
	return height
}

func (f *Framebuffer) GetBoldWhiteOnBlack() CellState {
	return BoldWhiteOnBlack
}

func (f *Framebuffer) GetAttrNone() Attribute {
	return AttrNone
}

func (f *Framebuffer) GetColorGray() Color {
	return ColorGray
}

func (f *Framebuffer) GetColorBlue() Color {
	return ColorBlue
}
