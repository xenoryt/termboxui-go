package termui

import "github.com/nsf/termbox-go"

//Label creates an area that displays text
type Label struct {
	x, y          int
	width, height int

	borders bool

	Title   string
	content string

	buffer []string

	fg, bg termbox.Attribute
}

func (lbl *Label) Origin() (x, y int)        { return }
func (lbl *Label) Size() (width, height int) { return }

func (lbl *Label) SetBorders(borders bool) {
	lbl.borders = borders
}

func (lbl *Label) Clear() {
	lbl.buffer = nil
	lbl.content = ""
}

func (lbl *Label) SetBG(attr termbox.Attribute) {
	lbl.fg = attr
}
func (lbl *Label) SetFG(attr termbox.Attribute) {
	lbl.bg = attr
}

//Write content to the label
func (lbl *Label) Write(p []byte) (n int, err error) {
	lbl.content = string(p[:])
	return len(lbl.content), nil
}
