package truco

import (
	"fmt"
	"math/rand"
	"time"
)

//Card is the basic piece for the card game
type Card struct {
	Face int
	Suit string
}

//Deck is a slice of cards. It contains the 40 necessary cards for playing truco
type Deck []Card

//Suits used on the card deck
const (
	SuitSpade   = "Spade"
	SuitHeart   = "Heart"
	SuitDiamond = "Diamond"
	SuitClub    = "Club"
)

// NewDeck creates and returns a new deck.
func NewDeck() Deck {
	rand.Seed(time.Now().UnixNano())
	d := make(Deck, 40, 40)
	suits := []string{SuitSpade, SuitHeart, SuitDiamond, SuitClub}

	for i, s := range suits {
		for j := 0; j < 10; j++ {
			d[10*i+j] = Card{j + 1, s}
		}
	}

	return d
}

//Shuffle changes randomly the order of the cards
func Shuffle(d Deck) {
	rand.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

//Print the card
func (c Card) Print() {
	fmt.Printf("%d %s\n", c.Face, c.Suit)
}

//Value returns the card value
func (c Card) Value() int {
	if c.Face == 1 && c.Suit == SuitSpade {
		return 12
	}

	if c.Face == 4 && c.Suit == SuitClub {
		return 14
	}

	if c.Face == 7 {
		if c.Suit == SuitDiamond {
			return 11
		} else if c.Suit == SuitHeart {
			return 13
		}
	}

	return (c.Face+6)%10 + 1
}
