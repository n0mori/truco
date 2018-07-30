package truco

import (
	"github.com/nsf/termbox-go"
)

//Sets constants related to card rendering
const (
	CardWidth  = 4
	CardHeight = 3
)

//DrawTable draws the playing table, i.e. the terminal area
func DrawTable() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorGreen)
}

//Draw draws a card in the position x, y. If visible is false, draws the cardback
/*
func (c Card) Draw(x, y int, visible bool) {
	str := " A234567JQK"
	faces := []rune(str)
	if visible {
		for i := x; i < x+CardWidth; i++ {
			for j := y; j < y+CardHeight; j++ {
				termbox.SetCell(i, j, ' ', termbox.ColorDefault, termbox.ColorWhite)
			}
		}

		var color termbox.Attribute
		if c.Suit == SuitSpade || c.Suit == SuitClub {
			color = termbox.ColorBlack
		} else {
			color = termbox.ColorRed
		}
		termbox.SetCell(x, y, faces[c.Face], color, termbox.ColorWhite)
		termbox.SetCell(x, y+1, c.Suit, color, termbox.ColorWhite)
	} else {
		for i := x; i < x+CardWidth; i++ {
			for j := y; j < y+CardHeight; j++ {
				termbox.SetCell(i, j, 'X', termbox.ColorCyan, termbox.ColorRed)
			}
		}
	}
}
*/
