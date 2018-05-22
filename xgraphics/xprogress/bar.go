package xprogress

type Xprogress struct {
	title   string
	maxLine int
	bars    []*Bar
}

func New(title string) *Xprogress {

	p := &Xprogress{
		title:   title,
		maxLine: gMaxLine,
	}

	gMaxLine++
	printf(gMaxLine, "title: %s", title)
	return p
}

func (p *Xprogress) NewBar(prefix string, total int) *Bar {
	gMaxLine++
	return NewBar(gMaxLine, prefix, total)
}