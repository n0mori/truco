package main

import (
	"time"

	"github.com/n0mori/truco/lib"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {

	}
	defer termbox.Close()

	deck := truco.NewDeck()
	truco.DrawTable()
	termbox.Flush()

	for i, c := range deck {
		c.Draw((i%10)*5, (i/10)*4, true)
	}
	deck[0].Draw(0, 0, false)
	termbox.Flush()
	time.Sleep(time.Second * 10)
}
