package xborders

import "fmt"

func (b *Borders) GetBorder(center bool) string {
	return processBorder(borderStyles[b.style], b.spacer, b.content, center)
}

func (b *Borders) PrintBorder(center bool) {
	fmt.Println(processBorder(borderStyles[b.style], b.spacer, b.content, center))
}
