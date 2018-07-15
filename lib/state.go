package truco

//PlayerState groups the informations about the player
type PlayerState struct {
	Tentos int
	Maos   int
	Cards  Deck
	Active bool
}

//GameState is the current state of the game
type GameState struct {
	PlayerStates []PlayerState
	Round        int
}

//NewPlayerState returns a pointer to a newly created PlayerState
func NewPlayerState() PlayerState {
	return PlayerState{
		Tentos: 0,
		Maos:   0,
		Cards:  make(Deck, 0),
		Active: false}
}

//Deal the cards to the you and enemy players
func Deal(you, enemy *PlayerState) {
	deck := NewDeck()
	you.Cards = deck[:4]
	enemy.Cards = deck[5:9]
	print(len(you.Cards), len(enemy.Cards))
}

//StartGame sets the initial state of the game.
func StartGame() GameState {
	you, enemy := NewPlayerState(), NewPlayerState()
	Deal(&you, &enemy)
	you.Active = true
	enemy.Active = false

	return GameState{
		PlayerStates: []PlayerState{you, enemy},
		Round:        0}

}
