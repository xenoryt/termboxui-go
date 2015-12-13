package termboxui

import "github.com/nsf/termbox-go"

type View struct {
	x, y          int
	width, height int

	// Position of the "cursor"
	cx, cy int

	borders bool

	Title string
}

func (v *View) Origin() (x, y int)        { return }
func (v *View) Size() (width, height int) { return }

/*WriteRow writes the string on the row the cursor is on and then moves
the cursor down one row. If the string exceeds width, then remaining
characters are ignored.
Returns EOF if cursor is outside the view.
*/
func (v *View) WriteRow(str []rune, fg, bg termbox.Attribute) error {
	if v.cy > v.height {
		return ErrEOF
	}
	y := v.y + v.cy
	for i := 0; i < v.width; i++ {
		termbox.SetCell(v.x+v.cx, y, str[i], fg, bg)
	}
	v.cy++
	return nil
}

/*WriteRowCell writes the row of termbox.Cell on the row the cursor is on
and then moves the cursor down one row. If the string exceeds width, then
remaining characters are ignored.
Returns EOF if cursor is outside the view.
*/
func (v *View) WriteRowCell(row []termbox.Cell) error {
	if v.cy > v.height {
		return ErrEOF
	}
	y := v.y + v.cy
	for i := 0; i < v.width; i++ {
		termbox.SetCell(v.x+v.cx, y, row[i].Ch, row[i].Fg, row[i].Bg)
	}
	v.cy++
	return nil
}

//Cursor returns the position of the cursor in the View
func (v *View) Cursor() (cx, cy int) { return }

func (v *View) MoveCursor(x, y int) {
	v.cx += x
	v.cy += y
}
func (v *View) MoveCursorTo(x, y int) {
	v.cx = x
	v.cy = y
}
