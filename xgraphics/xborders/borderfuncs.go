package xborders

import "fmt"

func (b *Borders) GetBorder() string {
	return processBorder(borderStyles[b.style], b.spacer, b.content)
}

func (b *Borders) PrintBorder() {
	fmt.Println(processBorder(borderStyles[b.style], b.spacer, b.content))
}
