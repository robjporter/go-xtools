package xborders

func New() *Borders {
	return &Borders{}
}

func (b *Borders) SetBorderStyle(style string) *Borders {
	b.style = style
	return b
}

func (b *Borders) SetContent(lines []string) *Borders {
	b.content = lines
	return b
}

func (b *Borders) SetSpacer(space int) *Borders {
	b.spacer = space
	return b
}
