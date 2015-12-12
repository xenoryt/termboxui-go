package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/xenoryt/termui-go"
)

func main() {
	//Initialize log
	f, err := os.Create("log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Initialize termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Create a new label at (3,3) with dimensions 18x5
	lbl := termui.NewLabel(3, 3, 18, 5)
	fmt.Fprintf(lbl, "Test Message! AB testing fox jumped over the fence!")
	fmt.Fprintln(lbl, "Moar messages! with moar line wrapping!")

mainloop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		lbl.Redraw()
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowDown:
				lbl.NextPage()
			case termbox.KeyArrowUp:
				lbl.PrevPage()
			case termbox.KeyEsc:
				break mainloop
			default:
				if ev.Ch == 'q' {
					break mainloop
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
