package termboxui

import (
	"errors"
	"log"

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
	log.Printf("termbox size: %d x %d\n", w, h)
	if location > -1 && location < 0 {
		location = 1 - location
	}
	return &VSplit{x: 0, y: 0, width: w, height: h, location: location, children: make([]Window, 2)}
}

//VSplit creates a vertical divider and tiles windows
//next to the split.
type VSplit struct {
	x, y          int
	width, height int

	children []Window
	location float32
}

func (s *VSplit) Move(x, y int) {
	s.x = x
	s.y = y

	if s.children[0] != nil {
		s.children[0].Move(x, y)
	}
	if s.children[1] != nil {
		s.children[1].Move(s.getSplitLoc()+1, s.y)
	}
}

func (s *VSplit) Resize(w, h int) {
	s.width = w
	s.height = h

	splitx := s.getSplitLoc()

	//Resize children
	if s.children[0] != nil {
		s.children[0].Resize(splitx-s.x, s.height)
	}

	if s.children[1] != nil {
		//need to move the second child since the location of split changed
		s.children[1].Move(s.getSplitLoc()+1, s.y)
		s.children[1].Resize(s.x+s.width-splitx-1, s.height)
	}
}

//Place places the window either to the left or right of the split.
//If both spaces are empty, it will be placed to the left.
//If both spaces have been taken, this will return an error.
func (s *VSplit) Place(win Window) error {
	splitx := s.getSplitLoc()
	if s.children[0] == nil {
		s.children[0] = win
		win.Move(s.x, s.y)
		win.Resize(splitx-s.x-1, s.height)
	} else if s.children[1] == nil {
		s.children[1] = win
		win.Move(s.getSplitLoc()+1, s.y)
		win.Resize(s.x+s.width-splitx-1, s.height)
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

//Gets the location of the split relative to the entire screen.
func (s *VSplit) getSplitLoc() int {
	if s.location > 0 && s.location < 1 {
		return s.x + int(float32(s.width)*s.location)
	}
	return s.x + (s.width+int(s.location))%s.width
}

//Draw draws the split and its children
func (s *VSplit) Draw() {
	DrawVertLine(s.getSplitLoc(), s.y, s.height)
	for _, f := range s.children {
		if f != nil {
			f.Draw()
		}
	}
}
