package main

import (
	"fmt"

	"github.com/n0mori/truco/lib"
)

func main() {
	/*
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
	*/

	player := truco.NewPlayer("localhost")

	var str string
	for {
		fmt.Scanln(&str)

		if str == "quit" {
			break
		}

		player.Move(str)
	}

	player.Close()
}
