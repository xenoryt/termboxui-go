package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/nsf/termbox-go"
	"github.com/xenoryt/termboxui-go"
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
	err = termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	split := termboxui.NewSplit(0.5)

	lbl := termboxui.NewLabel()
	lbl.Move(3, 6)
	b, err := ioutil.ReadFile("Mark.Twain-Tom.Sawyer.txt")
	fmt.Fprintf(lbl, string(b))

	split.Place(lbl)

mainloop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		split.Draw()
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				lbl.Scroll(-5)
			case termbox.KeyArrowDown:
				lbl.Scroll(5)
			case termbox.KeyEsc:
				break mainloop
			default:
				if ev.Ch == 'q' {
					break mainloop
				}
			}
		case termbox.EventResize:
			split.Resize(ev.Width, ev.Height)
		case termbox.EventError:
			log.Print(ev.Err)
			panic(ev.Err)
		}
	}
	log.Print("Exiting")
}
