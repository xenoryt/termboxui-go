package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	tbui "github.com/xenoryt/termboxui-go"
)

func DrawBufferTo(buffer [][]termbox.Cell, v *tbui.View) {
	for y, row := range buffer {
		for x, cell := range row {
			v.SetCell(x, y, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func HandleInput(ev termbox.Event, v *tbui.View) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEsc:
			return true
		case termbox.KeyArrowRight:
			v.ClearDefault()
			v.Move(1, 0)
		case termbox.KeyArrowLeft:
			v.ClearDefault()
			v.Move(-1, 0)
		default:
			switch ev.Ch {
			case 'q':
				return true
			}
		}
	case termbox.EventError:
		panic(ev.Err)
	}
	return false
}

func RandColor() termbox.Attribute {
	return termbox.Attribute(rand.Intn(int(termbox.ColorWhite)))
}

func RandomizeColor(cells []*termbox.Cell) {
	for _, cell := range cells {
		cell.Fg = RandColor()
	}
}

func main() {
	//Initialize log
	f, err := os.Create("log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Initialize termbox
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	v := tbui.NewView()
	v.Move(3, 3)
	v.Resize(10, 10)

	const BufferSize = 10

	buffer := make([][]termbox.Cell, BufferSize)
	border := make([]*termbox.Cell, 0, BufferSize*4-4)

	for y := range buffer {
		buffer[y] = make([]termbox.Cell, BufferSize)
		for x, _ := range buffer[y] {
			cell := &buffer[y][x]
			if x == 0 || x == len(buffer[y])-1 || y == 0 || y == len(buffer)-1 {
				cell.Ch = '#'
				border = append(border, cell)
			} else {
				cell.Ch = ' '
			}
			cell.Bg = termbox.ColorDefault
			cell.Fg = termbox.ColorDefault
		}
	}

	// Check for user input
	inp := make(chan termbox.Event, 3)
	go func() {
		for {
			ev := termbox.PollEvent()
			inp <- ev
		}
	}()

	ticker := time.NewTicker(time.Second)

mainloop:
	for {
		v.Clear(termbox.ColorDefault, termbox.ColorDefault)
		RandomizeColor(border)
		DrawBufferTo(buffer, v)
		termbox.Flush()
		select {
		case ev := <-inp:
			log.Print("handling input")
			if HandleInput(ev, v) {
				break mainloop
			}
		case <-ticker.C:
		}
	}
}
