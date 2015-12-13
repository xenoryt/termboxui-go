package termboxui

import (
	"io"
	"strings"
	"unicode/utf8"

	"github.com/nsf/termbox-go"
)

//NewLabel creates a new label
func NewLabel() *Label {
	lbl := &Label{x: -1, y: -1}
	return lbl
}

//Label creates an area that displays text
type Label struct {
	x, y          int
	width, height int

	borders    bool
	viewHeight int
	viewWidth  int

	Title   string
	content []string

	//The position (index) we are in the content
	startLine, startPos int
	endLine, endPos     int

	//buffer contains each row of text formatted so that it has
	//each string fits in the label
	buffer [][]byte

	fg, bg termbox.Attribute
}

func (lbl Label) Origin() (x, y int)        { return }
func (lbl Label) Size() (width, height int) { return }

func (lbl *Label) Move(x, y int) {
	lbl.x = x
	lbl.y = y
}
func (lbl *Label) Resize(width, height int) {
	lbl.width = width
	lbl.height = height
	lbl.checkViewSize()

	// We want to format the text so that it fits the new dimensions
	lbl.buffer = lbl.formatText(lbl.content)
}

func (lbl *Label) SetBorders(borders bool) {
	lbl.borders = borders
	lbl.checkViewSize()
}

func (lbl *Label) checkViewSize() {
	lbl.viewHeight = lbl.height
	lbl.viewWidth = lbl.width
	if lbl.borders {
		lbl.viewHeight -= 4
		lbl.viewWidth -= 4
	}
}

func (lbl *Label) Clear() {
	lbl.content = nil
	lbl.startPos = 0
	lbl.endPos = 0
}

func (lbl *Label) SetBG(attr termbox.Attribute) {
	lbl.fg = attr
}
func (lbl *Label) SetFG(attr termbox.Attribute) {
	lbl.bg = attr
}

//Draw writes the buffered text onto the screen
func (lbl Label) Draw() {
	if lbl.buffer == nil {
		return
	}

	for y := 0; y < lbl.viewHeight && lbl.startLine+y < len(lbl.buffer); y++ {
		slice := lbl.buffer[lbl.startLine+y][:]
		for x := 0; x < lbl.viewWidth; x++ {
			if len(slice) == 0 {
				termbox.SetCell(x+lbl.x, y+lbl.y, ' ', lbl.fg, lbl.bg)
			} else {
				r, size := utf8.DecodeRune(slice)
				termbox.SetCell(x+lbl.x, y+lbl.y, r, lbl.fg, lbl.bg)
				slice = slice[size:]
			}
		}
	}
}

//Redraw clears any previous text in the label and then perform a Draw
func (lbl Label) Redraw() {
	Fill(lbl.x, lbl.y, lbl.width, lbl.height, termbox.Cell{Ch: ' '})
	lbl.Draw()
}

//Write content to the label
func (lbl *Label) Write(p []byte) (n int, err error) {
	newText := strings.Split(string(p), "\n")
	lbl.content = append(lbl.content, newText...)
	lbl.buffer = append(lbl.buffer, lbl.formatText(newText)...)
	return len(p), nil
}

func (lbl Label) formatText(lines []string) (fmt [][]byte) {
	if lines == nil || lbl.width == 0 {
		return nil
	}
	// Initialize the buffer
	fmt = make([][]byte, 0, len(lines))

	for _, curLine := range lines {
		tmp := WrapText(curLine, lbl.viewWidth)
		for _, wrappedLine := range tmp {
			fmt = append(fmt, []byte(wrappedLine))
		}
	}

	return
}

func (lbl *Label) NextPage() error {
	lbl.startLine += lbl.viewHeight
	if lbl.startLine > len(lbl.buffer) {
		lbl.startLine = len(lbl.buffer) - 1
		return io.EOF
	}
	return nil
}
func (lbl *Label) PrevPage() error {
	lbl.startLine -= lbl.viewHeight
	if lbl.startLine < 0 {
		lbl.startLine = 0
		return io.EOF
	}
	return nil
}

func (lbl *Label) Scroll(amt int) error {
	lbl.startLine += amt
	if lbl.startLine > len(lbl.buffer) {
		lbl.startLine = len(lbl.buffer) - 1
		return io.EOF
	}
	if lbl.startLine < 0 {
		lbl.startLine = 0
		return io.EOF
	}
	return nil
}
