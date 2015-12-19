package termboxui

import "github.com/nsf/termbox-go"

//Fill fills a rectangular area with the given cell
func Fill(x, y, w, h int, cell termbox.Cell) {
	for i := y; i < y+h; i++ {
		for j := x; j < x+w; j++ {
			termbox.SetCell(j, i, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

//FillView fills a rectangular area in a View with the given cell
func FillView(view *View, x, y, w, h int, cell termbox.Cell) {
	for i := y; i < y+h; i++ {
		for j := x; j < x+w; j++ {
			view.SetCell(j, i, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

//Draw a vertical line starting at point (x, y) with length h
func DrawVertLine(x, y int, h int) {
	for i := y; i < y+h; i++ {
		termbox.SetCell(x, i, '│', termbox.ColorDefault, termbox.ColorDefault)
	}
}

//Draw a horizontal line starting at point (x, y) with length w
func DrawHorzLine(x, y int, w int) {
	for i := x; i < x+w; i++ {
		termbox.SetCell(i, y, '─', termbox.ColorDefault, termbox.ColorDefault)
	}
}

//Draws a box along the perimeter of the rectangular area
func DrawBox(x, y, w, h int) {
	//Draw the top and bottom
	for i := x + 1; i < x+w; i++ {
		termbox.SetCell(i, y, '─', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(i, y+h, '─', termbox.ColorDefault, termbox.ColorDefault)
	}

	//Draw the sides
	for i := y + 1; i < y+h; i++ {
		termbox.SetCell(x, i, '│', termbox.ColorDefault, termbox.ColorDefault)
		termbox.SetCell(x+w, i, '│', termbox.ColorDefault, termbox.ColorDefault)
	}

	//Draw the cornors
	termbox.SetCell(x, y, '┌', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(x, y+h, '└', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(x+w, y, '┐', termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(x+w, y+h, '┘', termbox.ColorDefault, termbox.ColorDefault)
}
