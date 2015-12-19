package termboxui

import "github.com/nsf/termbox-go"

//Creates a new View with the specified buffer size
//REQ: bufw and bufh must be greater than 0
func NewView() *View {
	w, h := termbox.Size()
	return &View{width: w, height: h}
}

//View imitates another termbox session. Can "pre-render"
//cells here and display them later on
type View struct {
	x, y          int
	width, height int
}

func (v *View) Origin() (x, y int)        { return v.x, v.y }
func (v *View) Size() (width, height int) { return v.width, v.height }

//Move moves the location of the view by the specified offset.
//It does not move the content that has already been rendered
//but all future content will be rendered at the new
//location.
func (v *View) Move(xOffset, yOffset int) {
	v.x += xOffset
	v.y += yOffset
}

//MoveTo moves the view to the specified location.
//It does not move the content that has already been rendered
//but all future content will be rendered at the new
//location.
func (v *View) MoveTo(x, y int) {
	v.x = x
	v.y = y
}

func (v *View) Resize(w, h int) {
	v.width = w
	v.height = h
}

func (v *View) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if x < 0 || x >= v.width {
		return
	}
	if y < 0 || y >= v.height {
		return
	}
	termbox.SetCell(v.x+x, v.y+y, ch, fg, bg)
}

func (v *View) Clear(fg, bg termbox.Attribute) {
	Fill(v.x, v.y, v.width, v.height, termbox.Cell{Fg: fg, Bg: bg, Ch: ' '})
}
func (v *View) ClearDefault() {
	Fill(v.x, v.y, v.width, v.height, termbox.Cell{Fg: termbox.ColorDefault, Bg: termbox.ColorDefault, Ch: ' '})
}
