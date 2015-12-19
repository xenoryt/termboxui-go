package termboxui

import (
	"errors"
	"log"

	"github.com/nsf/termbox-go"
)

type SplitType int

const (
	SplitHorizontal SplitType = iota
	SplitVertical
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

type Split interface {
	Draw()
	Place(Window) error
	Remove(Window)
	RemoveFirst()
	RemoveLast()
	Move(x, y int)
	Resize(w, h int)
}

//NewSplit creates a new horizontal split.
//If the location is positive then it is based starting from the left/top.
//If the location is negative then it starts from the right/bottom.
//If location is on the interval (0, 1), then
//it is treated as a percentage and the split will appear that percent down the screen.
//If location is >= 1 then it is truncated and will appear on the location'th row.
//The split will take up the entire screen by default.
//A VSplit.Move and VSplit.Resize is necessary to place it in the correct position if
//that is not the behaviour you want.
func NewSplit(location float32, sType SplitType) Split {
	w, h := termbox.Size()
	if location > -1 && location < 0 {
		location = 1 + location
	}
	if sType == SplitVertical {
		return &VSplit{x: 0, y: 0, width: w, height: h, location: location, children: make([]Window, 2)}
	} else {
		return &HSplit{x: 0, y: 0, width: w, height: h, location: location, children: make([]Window, 2)}
	}
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
		s.children[1].Move(s.GetSplitLoc()+1, s.y)
	}
}

func (s *VSplit) Resize(w, h int) {
	s.width = w
	s.height = h

	splitx := s.GetSplitLoc()

	//Resize children
	if s.children[0] != nil {
		s.children[0].Resize(splitx-s.x, s.height)
	}

	if s.children[1] != nil {
		//need to move the second child since the location of split changed
		s.children[1].Move(s.GetSplitLoc()+1, s.y)
		s.children[1].Resize(s.x+s.width-splitx-1, s.height)
	}
}

//Place places the window either to the left or right of the split.
//If both spaces are empty, it will be placed to the left.
//If both spaces have been taken, this will return an error.
func (s *VSplit) Place(win Window) error {
	splitx := s.GetSplitLoc()
	if s.children[0] == nil {
		s.children[0] = win
		win.Move(s.x, s.y)
		win.Resize(splitx-s.x-1, s.height)
	} else if s.children[1] == nil {
		s.children[1] = win
		win.Move(s.GetSplitLoc()+1, s.y)
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
func (s *VSplit) GetSplitLoc() int {
	if s.location > 0 && s.location < 1 {
		return s.x + int(float32(s.width)*s.location)
	}
	return s.x + (s.width+int(s.location))%s.width
}

//Draw draws the split and its children
func (s *VSplit) Draw() {
	DrawVertLine(s.GetSplitLoc(), s.y, s.height)
	for _, f := range s.children {
		if f != nil {
			f.Draw()
		}
	}
}

//HSplit creates a vertical divider and tiles windows
//next to the split.
type HSplit struct {
	x, y          int
	width, height int

	children []Window
	location float32
}

func (s *HSplit) Move(x, y int) {
	s.x = x
	s.y = y

	s.moveChild(3)
}

func (s *HSplit) Resize(w, h int) {
	s.width = w
	s.height = h

	//Resize children
	s.resizeChild(3)
}

//moves child into the correct location.
//1 is first child
//2 is second child
//3 is both
func (s *HSplit) moveChild(c int) {
	if c&1 != 0 && s.children[0] != nil {
		s.children[0].Move(s.x, s.y)
	}
	if c&2 != 0 && s.children[1] != nil {
		s.children[1].Move(s.x, s.GetSplitLoc()+1)
	}
}
func (s *HSplit) resizeChild(c int) {
	splity := s.GetSplitLoc()
	if c&1 != 0 && s.children[0] != nil {
		s.children[0].Resize(s.width, splity-s.y)
	}
	if c&2 != 0 && s.children[1] != nil {
		s.children[1].Move(s.x, s.GetSplitLoc()+1)
		s.children[1].Resize(s.width, s.y+s.height-splity-1)
	}
}

//Place places the window either to the left or right of the split.
//If both spaces are empty, it will be placed to the left.
//If both spaces have been taken, this will return an error.
func (s *HSplit) Place(win Window) error {
	if s.children[0] == nil {
		s.children[0] = win
		s.moveChild(1)
		s.resizeChild(1)
	} else if s.children[1] == nil {
		s.children[1] = win
		s.resizeChild(2) //resizing second child moves and resizes
	} else {
		log.Print("container full")
		return errors.New("HSplit container is full")
	}
	return nil
}

//Remove removes the window and makes its occupied space available again.
func (s *HSplit) Remove(win Window) {
	for i, f := range s.children {
		if f == win {
			s.children[i] = nil
		}
	}
}

//RemoveFirst removes the window to the left
func (s *HSplit) RemoveFirst() {
	s.children[0] = nil
}

//RemoveLast removes the window to the right
func (s *HSplit) RemoveLast() {
	s.children[1] = nil
}

//Gets the location of the split relative to the entire screen.
func (s *HSplit) GetSplitLoc() int {
	if s.location > 0 && s.location < 1 {
		return s.y + int(float32(s.height)*s.location)
	}
	return s.y + (s.height+int(s.location))%s.height
}

//Draw draws the split and its children
func (s *HSplit) Draw() {
	DrawHorzLine(s.x, s.GetSplitLoc(), s.width)
	for _, f := range s.children {
		if f != nil {
			f.Draw()
		}
	}
}
