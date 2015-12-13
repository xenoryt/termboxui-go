package termboxui

func NewFrame() *Frame {
	return &Frame{x: 0, y: 0}
}

type Frame struct {
	x, y          int
	width, height int

	child Window
}

func (f *Frame) Move(x, y int) {
	f.x = x
	f.y = y
	if f.child != nil {
		f.child.Move(x, y)
	}
}
func (f *Frame) Resize(w, h int) {
	f.width = w
	f.height = h
	if f.child != nil {
		f.child.Resize(w, h)
	}
}

func (f *Frame) Place(win Window) {
	f.child = win
}

func (f *Frame) Draw() {
	if f.child != nil {
		f.child.Draw()
	}
}
