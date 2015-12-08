package termui

import "github.com/nsf/termbox-go"

//Box creates an area that displays text
type Box struct {
	x, y          int
	width, height int

	borders bool

	Title   string
	content string

	fg, bg termbox.Attribute
}

func (b *Box) Origin() (x, y int)        { return }
func (b *Box) Size() (width, height int) { return }

// SetBorders
func (b *Box) SetBorders(borders bool) {
	b.borders = borders
}

func (b *Box) SetBG(attr termbox.Attribute) {
	b.fg = attr
}
func (b *Box) SetFG(attr termbox.Attribute) {
	b.bg = attr
}

func (b *Box) Write(p []byte) (n int, err error) {
	b.content = string(p[:])
	return len(b.content), nil
}
