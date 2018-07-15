package truco

//State is the current state of the game
type State struct {
	Cards   Deck
	Players []Player
}
