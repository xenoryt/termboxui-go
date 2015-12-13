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

	split := termboxui.NewSplit(-5, termboxui.SplitHorizontal)
	vsplit := termboxui.NewSplit(-0.25, termboxui.SplitVertical)
	vsplit.Place(split)

	lbl := termboxui.NewLabel()
	b, err := ioutil.ReadFile("Mark.Twain-Tom.Sawyer.txt")
	fmt.Fprintf(lbl, string(b))

	split.Place(lbl)

	lbl2 := termboxui.NewLabel()
	fmt.Fprintf(lbl2, string(b))
	vsplit.Place(lbl2)

mainloop:
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		vsplit.Draw()
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowUp:
				lbl.Scroll(-3)
			case termbox.KeyArrowDown:
				lbl.Scroll(3)
			case termbox.KeyEsc:
				break mainloop
			default:
				switch ev.Ch {
				case 'q':
					break mainloop
				case '+':
					lbl2.NextPage()
				case '-':
					lbl2.PrevPage()
				}
			}
		case termbox.EventResize:
			vsplit.Resize(ev.Width, ev.Height)
		case termbox.EventError:
			log.Print(ev.Err)
			panic(ev.Err)
		}
	}
	log.Print("Exiting")
}
