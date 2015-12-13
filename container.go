package termboxui

import (
	"errors"

	"github.com/nsf/termbox-go"
)

//Window represents any UI element
type Window interface {
	Draw()
	Move(x, y int)
	Resize(width, height int)
}

//Container represents anything that can
//store and display Windows.
//Generally containers are used to tile the Windows
//as well as handling resizing the Windows
type Container interface {
	Draw()
	Place(Window) error
	Remove(Window)
	Move(x, y int)
	Resize(width, height int)
}

//NewSplit creates a new horizontal split. If location is on the interval (0, 1), then
//it is treated as a percentage and the split will appear that percent down the screen.
//If location is >0 then it is truncated and will appear on the location'th row.
//The split will take up the entire screen by default.
//A VSplit.Move and VSplit.Resize is necessary to place it in the correct position if
//that is not the behaviour you want.
func NewSplit(location float32) *VSplit {
	w, h := termbox.Size()
	if location > -1 && location < 0 {
		location = 1 - location
	}
	return &VSplit{x: 0, y: 0, width: w, height: h, location: location}
}

//VSplit creates a vertical divider and tiles windows
//next to the split.
type VSplit struct {
	x, y          int
	width, height int

	children []Window
	location float32
}

//Place places the window either to the left or right of the split.
//If both spaces are empty, it will be placed to the left.
//If both spaces have been taken, this will return an error.
func (s *VSplit) Place(win Window, size int) error {
	if s.children[0] == nil {
		s.children[0] = win
	} else if s.children[1] == nil {
		s.children[1] = win
	} else {
		return errors.New("VSplit container is full")
	}
	return nil
}

//Remove removes the window and makes its occupied space available again.
func (s *VSplit) Remove(win Window) {
	for i, f := range s.children {
		if f == win {
			s.children[i] = nil
		}
	}
}

//RemoveFirst removes the window to the left
func (s *VSplit) RemoveFirst() {
	s.children[0] = nil
}

//RemoveLast removes the window to the right
func (s *VSplit) RemoveLast() {
	s.children[1] = nil
}

//Gets the location of the split
func (s *VSplit) getSplitLoc() int {
	if s.location > 0 && s.location < 1 {
		return s.x + s.width*int(s.location)
	}
	return s.x + (s.width+int(s.location))%s.width
}

//Draw draws the split and its children
func (s *VSplit) Draw() {
	DrawVertLine(s.getSplitLoc(), s.y, s.height)
	for _, f := range s.children {
		f.Draw()
	}
}
